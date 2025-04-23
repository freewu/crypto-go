package crypto

// SM3密码杂凑算法 - SM3 cryptographic hash algorithm
// 遵循的SM3标准号为： GM/T 0004-2012

import (
	"fmt"
	"github.com/tjfoc/gmsm/sm3"
)

func SM3(str string) string {
	h := sm3.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}
