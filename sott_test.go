package loginradius

import (
	"strings"
	"testing"
	"time"
)

// SOTT is shared wire format, not an implementation detail: every LoginRadius
// SDK must produce the same token for the same inputs, or a server-minted SOTT
// from one language will be rejected when another verifies it. The golden value
// below is what pins that — it is the reason these tests exist as more than a
// smoke check.

// TestSOTTGoldenValue pins the exact token the SOTT algorithm must produce.
//
// CROSS-LANGUAGE PARITY: the Node SDK asserts this same golden value in
// __tests__/sott.test.ts. Every LoginRadius SDK must emit a byte-identical
// token — the API validates the AES payload exactly, so a drifted IV,
// iteration count, salt, or timestamp format yields a token that is silently
// rejected. If this fails, fix the implementation, not the expectation. The
// shared parameters live in the SDK generator's shared configuration under `sott:`.
func TestSOTTGoldenValue(t *testing.T) {
	const (
		apiKey    = "test-api-key"
		apiSecret = "test-api-secret"
		golden    = "yvLBFPR3aRNl1YlisgPpEdphb73sUfne2Jem7hTKWU6RLlkcfjYOhe5B7kSHorQS*1059092e1510bfbc5388d7438b943106"
	)
	start := time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
	end := time.Date(2026, 1, 2, 3, 14, 5, 0, time.UTC)

	got, err := GenerateSOTTWithWindow(apiKey, apiSecret, start, end)
	if err != nil {
		t.Fatalf("GenerateSOTTWithWindow: %v", err)
	}
	if got != golden {
		t.Errorf("SOTT drifted from the cross-language golden value:\ngot:  %s\nwant: %s", got, golden)
	}
}

func TestSOTTRejectsMissingCredentials(t *testing.T) {
	start := time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
	end := start.Add(10 * time.Minute)
	if _, err := GenerateSOTTWithWindow("", "secret", start, end); err == nil {
		t.Error("GenerateSOTTWithWindow with no apiKey = nil error, want an error")
	}
	if _, err := GenerateSOTTWithWindow("key", "", start, end); err == nil {
		t.Error("GenerateSOTTWithWindow with no apiSecret = nil error, want an error")
	}
}

func TestSOTTDefaultWindowProducesToken(t *testing.T) {
	got, err := GenerateSOTT("test-api-key", "test-api-secret")
	if err != nil {
		t.Fatalf("GenerateSOTT: %v", err)
	}
	if parts := strings.Split(got, "*"); len(parts) != 2 {
		t.Errorf("SOTT = %q, want base64*md5hex", got)
	}
}
