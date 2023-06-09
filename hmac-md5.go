package crypto

import (
	"crypto/hmac"
	"crypto/md5"
	"fmt"
)

func HmacMD5(str string, key string) string {
	hash := hmac.New(md5.New, []byte(key))
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum([]byte("")))
}
