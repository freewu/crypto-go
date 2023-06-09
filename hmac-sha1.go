package crypto

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

func HmacSHA1(str string, key string) string {
	hash := hmac.New(sha1.New, []byte(key))
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum([]byte("")))
}
