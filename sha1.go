package crypto

import (
	"crypto/sha1"
	"fmt"
)

func SHA1(str string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(str)))
}
