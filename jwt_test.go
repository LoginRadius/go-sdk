package loginradius_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"

	loginradius "github.com/LoginRadius/go-sdk/v12"
)

const jwtSecret = "a-shared-secret-at-least-32-bytes-long!!"

func signHS(t *testing.T, method jwt.SigningMethod, claims jwt.MapClaims, secret string) string {
	t.Helper()
	s, err := jwt.NewWithClaims(method, claims).SignedString([]byte(secret))
	if err != nil {
		t.Fatalf("sign: %v", err)
	}
	return s
}

func liveClaims() jwt.MapClaims {
	return jwt.MapClaims{
		"sub": "uid-123",
		"iss": "LoginRadius",
		"aud": "my-app",
		"exp": time.Now().Add(time.Hour).Unix(),
		"nbf": time.Now().Add(-time.Minute).Unix(),
	}
}

func pemPublic(t *testing.T, pub any) []byte {
	t.Helper()
	der, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	return pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: der})
}

func TestValidateJWTHappyPath(t *testing.T) {
	token := signHS(t, jwt.SigningMethodHS256, liveClaims(), jwtSecret)
	claims, err := loginradius.ValidateJWT(token, loginradius.JWTValidationParams{
		Algorithm: loginradius.JWTHS256,
		Key:       []byte(jwtSecret),
		Issuer:    "LoginRadius",
		Audience:  "my-app",
	})
	if err != nil {
		t.Fatalf("ValidateJWT: %v", err)
	}
	if claims["sub"] != "uid-123" {
		t.Errorf("sub = %v, want uid-123", claims["sub"])
	}
}

// THE attack this utility exists to stop. Against an RS256-configured app, an
// attacker signs a token with HS256 using the PUBLIC key as the HMAC secret.
// A validator that reads the algorithm from the token header accepts it.
func TestValidateJWTRejectsAlgorithmConfusion(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("keygen: %v", err)
	}
	pub := pemPublic(t, &key.PublicKey)

	// The attacker's forgery: HMAC-signed with the public key as the secret.
	forged := signHS(t, jwt.SigningMethodHS256, liveClaims(), string(pub))

	_, err = loginradius.ValidateJWT(forged, loginradius.JWTValidationParams{
		Algorithm: loginradius.JWTRS256, // what the app is actually configured for
		Key:       pub,
	})
	if err == nil {
		t.Fatal("algorithm-confusion forgery was ACCEPTED — the algorithm is not pinned")
	}
}

// `alg: none` — the other classic. jwt/v5 refuses to sign one, so the token is
// assembled by hand, which is exactly how an attacker would produce it.
func TestValidateJWTRejectsAlgNone(t *testing.T) {
	// {"alg":"none","typ":"JWT"} . {"sub":"attacker","exp":<far future>} . <empty>
	unsigned := "eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0." +
		"eyJzdWIiOiJhdHRhY2tlciIsImV4cCI6NDEwMjQ0NDgwMH0."
	for _, alg := range []loginradius.JWTAlgorithm{loginradius.JWTHS256, loginradius.JWTRS256} {
		if _, err := loginradius.ValidateJWT(unsigned, loginradius.JWTValidationParams{
			Algorithm: alg,
			Key:       []byte(jwtSecret),
		}); err == nil {
			t.Errorf("alg=none token was ACCEPTED when expecting %s", alg)
		}
	}
}

func TestValidateJWTRejectsTamperedSignature(t *testing.T) {
	token := signHS(t, jwt.SigningMethodHS256, liveClaims(), "a-different-secret-entirely-32-bytes")
	if _, err := loginradius.ValidateJWT(token, loginradius.JWTValidationParams{
		Algorithm: loginradius.JWTHS256,
		Key:       []byte(jwtSecret),
	}); err == nil {
		t.Fatal("token signed with the wrong secret was accepted")
	}
}

func TestValidateJWTRejectsExpired(t *testing.T) {
	claims := liveClaims()
	// Beyond the 60s leeway.
	claims["exp"] = time.Now().Add(-10 * time.Minute).Unix()
	token := signHS(t, jwt.SigningMethodHS256, claims, jwtSecret)
	if _, err := loginradius.ValidateJWT(token, loginradius.JWTValidationParams{
		Algorithm: loginradius.JWTHS256, Key: []byte(jwtSecret),
	}); err == nil {
		t.Fatal("expired token was accepted")
	}
}

func TestValidateJWTRejectsNotYetValid(t *testing.T) {
	claims := liveClaims()
	claims["nbf"] = time.Now().Add(10 * time.Minute).Unix()
	token := signHS(t, jwt.SigningMethodHS256, claims, jwtSecret)
	if _, err := loginradius.ValidateJWT(token, loginradius.JWTValidationParams{
		Algorithm: loginradius.JWTHS256, Key: []byte(jwtSecret),
	}); err == nil {
		t.Fatal("not-yet-valid token was accepted")
	}
}

// A token with no exp never stops being valid, so it is refused outright.
func TestValidateJWTRequiresExpiry(t *testing.T) {
	token := signHS(t, jwt.SigningMethodHS256, jwt.MapClaims{"sub": "uid-123"}, jwtSecret)
	_, err := loginradius.ValidateJWT(token, loginradius.JWTValidationParams{
		Algorithm: loginradius.JWTHS256, Key: []byte(jwtSecret),
	})
	if !errors.Is(err, loginradius.ErrJWTMissingExpiry) {
		t.Fatalf("err = %v, want ErrJWTMissingExpiry", err)
	}
}

func TestValidateJWTChecksIssuerAndAudience(t *testing.T) {
	token := signHS(t, jwt.SigningMethodHS256, liveClaims(), jwtSecret)
	for _, tc := range []struct{ name, iss, aud string }{
		{"wrong issuer", "SomeoneElse", "my-app"},
		{"wrong audience", "LoginRadius", "another-app"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := loginradius.ValidateJWT(token, loginradius.JWTValidationParams{
				Algorithm: loginradius.JWTHS256, Key: []byte(jwtSecret),
				Issuer: tc.iss, Audience: tc.aud,
			}); err == nil {
				t.Fatal("mismatch was accepted")
			}
		})
	}
}

// Issuer and audience are optional; omitting them must not silently disable
// the checks that are NOT optional.
func TestValidateJWTIssuerAndAudienceAreOptional(t *testing.T) {
	token := signHS(t, jwt.SigningMethodHS256, liveClaims(), jwtSecret)
	if _, err := loginradius.ValidateJWT(token, loginradius.JWTValidationParams{
		Algorithm: loginradius.JWTHS256, Key: []byte(jwtSecret),
	}); err != nil {
		t.Fatalf("ValidateJWT: %v", err)
	}
}

func TestValidateJWTRoundTripsRSAAndECDSA(t *testing.T) {
	rsaKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("rsa keygen: %v", err)
	}
	rsaToken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, liveClaims()).SignedString(rsaKey)
	if err != nil {
		t.Fatalf("rsa sign: %v", err)
	}
	if _, err := loginradius.ValidateJWT(rsaToken, loginradius.JWTValidationParams{
		Algorithm: loginradius.JWTRS256, Key: pemPublic(t, &rsaKey.PublicKey),
	}); err != nil {
		t.Errorf("RS256: %v", err)
	}

	ecKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("ec keygen: %v", err)
	}
	ecToken, err := jwt.NewWithClaims(jwt.SigningMethodES256, liveClaims()).SignedString(ecKey)
	if err != nil {
		t.Fatalf("ec sign: %v", err)
	}
	if _, err := loginradius.ValidateJWT(ecToken, loginradius.JWTValidationParams{
		Algorithm: loginradius.JWTES256, Key: pemPublic(t, &ecKey.PublicKey),
	}); err != nil {
		t.Errorf("ES256: %v", err)
	}
}

// Handing over a PRIVATE key is a serious mistake; it is named rather than
// surfaced as a vague parse failure.
func TestValidateJWTRejectsPrivateKeyAsVerificationKey(t *testing.T) {
	key, _ := rsa.GenerateKey(rand.Reader, 2048)
	priv := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key),
	})
	token, _ := jwt.NewWithClaims(jwt.SigningMethodRS256, liveClaims()).SignedString(key)
	_, err := loginradius.ValidateJWT(token, loginradius.JWTValidationParams{
		Algorithm: loginradius.JWTRS256, Key: priv,
	})
	if !errors.Is(err, loginradius.ErrJWTInvalidKey) {
		t.Fatalf("err = %v, want ErrJWTInvalidKey", err)
	}
}

func TestValidateJWTRejectsUnknownAlgorithmAndEmptyKey(t *testing.T) {
	token := signHS(t, jwt.SigningMethodHS256, liveClaims(), jwtSecret)
	if _, err := loginradius.ValidateJWT(token, loginradius.JWTValidationParams{
		Algorithm: "HS1024", Key: []byte(jwtSecret),
	}); !errors.Is(err, loginradius.ErrJWTUnsupportedAlgorithm) {
		t.Errorf("err = %v, want ErrJWTUnsupportedAlgorithm", err)
	}
	if _, err := loginradius.ValidateJWT(token, loginradius.JWTValidationParams{
		Algorithm: loginradius.JWTHS256,
	}); !errors.Is(err, loginradius.ErrJWTInvalidKey) {
		t.Errorf("err = %v, want ErrJWTInvalidKey", err)
	}
}
