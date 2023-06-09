package crypto

import (
	"crypto/sha256"
	"fmt"
)

func SHA256(str string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(str))) // 将[]byte转成16进制
}
