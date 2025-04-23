package sm4

import (
	"errors"
	"github.com/tjfoc/gmsm/sm4"
)

// SM4分组密码算法 - SM4 block cipher algorithm
// 遵循的SM4标准号为: GM/T 0002-2012

type SM4 struct {
	Mode string // 加密模式  ECB / CBC / CFB / OFB	default: ECB
}

// GetInstance 实例化一个 des 对象  如: GetInstance("ECB/PKCS7Padding")
func GetInstance(str string) *SM4 {
	if str != "" {
		// todo 需要验证处理
		return &SM4{Mode: str}
	}
	return &SM4{Mode: "ECB"}
}

// Encrypt 加密处理
func (d *SM4) Encrypt(data []byte, key []byte, iv []byte) ([]byte, error) {
	if len(iv) != 0 {
		err := sm4.SetIV(iv) // 设置SM4算法实现的IV值,不设置则使用默认值
		if err != nil {      // 设置 iv 失败
			return nil, err
		}
	}
	switch d.Mode {
	case "CBC":
		return sm4.Sm4Cbc(key, data, true)
	case "ECB":
		return sm4.Sm4Ecb(key, data, true)
	case "CFB":
		return sm4.Sm4CFB(key, data, true)
	case "OFB":
		return sm4.Sm4OFB(key, data, true)
	}
	return nil, errors.New("no such mode")
}

// Decrypt 解密处理
func (d *SM4) Decrypt(data []byte, key []byte, iv []byte) ([]byte, error) {
	if len(iv) != 0 {
		err := sm4.SetIV(iv) // 设置SM4算法实现的IV值,不设置则使用默认值
		if err != nil {      // 设置 iv 失败
			return nil, err
		}
	}
	switch d.Mode {
	case "CBC":
		return sm4.Sm4Cbc(key, data, false)
	case "ECB":
		return sm4.Sm4Ecb(key, data, false)
	case "CFB":
		return sm4.Sm4CFB(key, data, false)
	case "OFB":
		return sm4.Sm4OFB(key, data, false)
	}
	return nil, errors.New("no such mode")
}
