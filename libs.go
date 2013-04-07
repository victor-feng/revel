package revel

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"io"
)

// Sign a given string with the app-configured secret key.
// If no secret key is set, returns the empty string.
// Return the signature in base64 (URLEncoding).
func Sign(message string) string {
	if len(secretKey) == 0 {
		return ""
	}
	mac := hmac.New(sha1.New, secretKey)
	io.WriteString(mac, message)
	return hex.EncodeToString(mac.Sum(nil))
}
