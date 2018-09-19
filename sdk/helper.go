package loginradius

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/pbkdf2"
)

// TimeAlt is a struct to hold Time values from all sources
type TimeAlt struct {
	time.Time
}

// HTTPError struct to return HTTP information
type HTTPError struct {
	Request    string
	Response   string
	StatusCode int
}

func (e *HTTPError) Error() string {
	return e.Response
}

// UnmarshalJSON is an override of time's parsing method
func (j *TimeAlt) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return nil
	}
	t, err := time.Parse(time.RFC3339, s)
	if err == nil {
		*j = TimeAlt{t}
		return nil
	}

	t2, err2 := time.Parse("2006-01-02T15:04:05-0700", s)
	if err2 == nil {
		*j = TimeAlt{t2}
		return nil
	}

	t3, err3 := time.Parse("01/02/2006 03:04:05 PM +07:00 ", s)
	if err3 == nil {
		*j = TimeAlt{t3}
		return nil
	}
	return nil
}

// CreateRequest takes in a method (EX: "GET","PUT","POST"), the server URI
// and a request body (For empty bodies, pass in an empty string) and generates
// an *http.Request object which can be passed into RunRequest to process the request
func CreateRequest(method, uri string, body interface{}) (*http.Request, error) {
	buffer := new(bytes.Buffer)
	encodeErr := json.NewEncoder(buffer).Encode(body)
	if encodeErr != nil {
		return nil, encodeErr
	}

	req, newReqErr := http.NewRequest(method, uri, buffer)
	if newReqErr != nil {
		return nil, newReqErr
	}
	return req, nil
}

// RunRequest takes in an *http.Request object and a destination struct to hold the
// Response body. The destination struct should contain fields for all JSON attributes
// or an error will be returned. This method processes the http request and places data
// into respective fields within the destination struct. For struct fields that do not
// have a corresponding JSON attribute, the default empty object will be provided for the
// struct field. Errors will also be returned if the response status code is 2xx.
func RunRequest(request *http.Request, dst interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	request = request.WithContext(ctx)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		decoder := json.NewDecoder(resp.Body)
		decoder.DisallowUnknownFields()
		return decoder.Decode(dst)
	}
	reqBody, err2 := request.GetBody()
	if err2 != nil {
		return err2
	}
	requestData, err3 := ioutil.ReadAll(reqBody)
	if err3 != nil {
		return err3
	}
	responseData, err4 := ioutil.ReadAll(resp.Body)
	if err4 != nil {
		return err4
	}

	retError := &HTTPError{string(requestData), string(responseData), resp.StatusCode}
	return retError
}

// GenerateSOTT is a function that generates a SOTT through the methods described here:
// https://docs.loginradius.com/api/v2/user-registration/sott
func GenerateSOTT() string {
	plainText := generatePlainText()
	tempToken := encrypt(plainText)
	token := strings.Replace(tempToken, "-", "+", -1)
	readyToken := strings.Replace(token, "_", "/", -1)
	hash := getMD5Hash(readyToken)
	retToken := readyToken + "*" + hash
	return retToken
}

func generatePlainText() string {
	key := os.Getenv("APIKEY")
	initTime := time.Now().UTC().Add(time.Duration(-10) * time.Minute)
	endTime := time.Now().UTC().Add(time.Duration(10) * time.Minute)
	initTimestamp := fmt.Sprintf("%s %d%s", initTime.Format("2006/1/2"), initTime.Hour(), initTime.Format(":4:5"))
	endTimestamp := fmt.Sprintf("%s %d%s", endTime.Format("2006/1/2"), endTime.Hour(), endTime.Format(":4:5"))
	retTime := initTimestamp + "#" + key + "#" + endTimestamp
	return retTime
}

func pKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func encrypt(plaintext string) string {
	secret := os.Getenv("APISECRET")
	initVector := "tu89geji340t89u2"
	salt := make([]byte, 8)
	password := pbkdf2.Key([]byte(secret), salt, 10000, 32, sha1.New)

	data := []byte(plaintext)

	block, err := aes.NewCipher(password)
	if err != nil {
		panic(err)
	}

	iv := []byte(initVector)

	blockSize := block.BlockSize()
	origData := pKCS5Padding(data, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(origData))
	blockMode.CryptBlocks(encrypted, origData)

	return base64.URLEncoding.EncodeToString(encrypted)
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	_, err := hasher.Write([]byte(text))
	if err != nil {
		log.Print("Error creating hash for SOTT.")
	}
	return hex.EncodeToString(hasher.Sum(nil))
}
