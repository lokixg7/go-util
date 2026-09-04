package crypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

// MD5 returns the lowercase hex MD5 digest of text.
func MD5(text string) string {
	sum := md5.Sum([]byte(text))
	return hex.EncodeToString(sum[:])
}

// SHA1 returns the lowercase hex SHA-1 digest of text.
func SHA1(text string) string {
	sum := sha1.Sum([]byte(text))
	return hex.EncodeToString(sum[:])
}

// SHA256 returns the lowercase hex SHA-256 digest of text.
func SHA256(text string) string {
	sum := sha256.Sum256([]byte(text))
	return hex.EncodeToString(sum[:])
}

// HMACSHA256 returns the lowercase hex HMAC-SHA256 digest of text using key.
func HMACSHA256(key, text string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(text))
	return hex.EncodeToString(mac.Sum(nil))
}

// Base64Encode returns the standard base64 encoding of text.
func Base64Encode(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

// Base64Decode decodes standard base64 data and returns the decoded text.
func Base64Decode(text string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
