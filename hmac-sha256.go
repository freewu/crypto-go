package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func HmacSHA256(str string, key string) string {
	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum([]byte("")))
}
