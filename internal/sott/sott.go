// The SOTT package is used to generate a LoginRadius Secured One Time Token
package sott

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/LoginRadius/go-sdk/lrerror"
	"golang.org/x/crypto/pbkdf2"
)

// Generates a SOTT through the methods described here:
// https://docs.loginradius.com/api/v2/user-registration/sott
func Generate(key string, secret string) string {
	plainText := generatePlainText(key)
	tempToken := encrypt(plainText, secret)
	token := strings.Replace(tempToken, "-", "+", -1)
	readyToken := strings.Replace(token, "_", "/", -1)
	hash := getMD5Hash(readyToken)
	return readyToken + "*" + hash
}

func generatePlainText(k string) string {
	key := k
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

func encrypt(plaintext string, secret string) string {
	initVector := "tu89geji340t89u2"
	salt := make([]byte, 8)
	password := pbkdf2.Key([]byte(secret), salt, 10000, 32, sha1.New)

	data := []byte(plaintext)

	block, err := aes.NewCipher(password)
	if err != nil {
		err = lrerror.New("EncryptionError", "Error occurred during sott encryption", err)
		log.Println(err.Error())
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
		err = lrerror.New("MD5HashError", "Error creating hash for SOTT", err)
		log.Println(err.Error())
	}
	return hex.EncodeToString(hasher.Sum(nil))
}
