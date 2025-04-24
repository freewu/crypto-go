package sm2

import (
	"crypto/rand"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

type SM2 struct {
}

func GetInstance() *SM2 {
	return &SM2{}
}

// GenerateKey 生成公钥和私钥
func (r *SM2) GenerateKey(pwd []byte) ([]byte, []byte, error) {
	privateKey, err := sm2.GenerateKey(rand.Reader) // 生成密钥对
	if err != nil {
		return nil, nil, err
	}
	// 序列化私钥（PKCS#8）
	privateKeyBytes, err := x509.MarshalSm2PrivateKey(privateKey, pwd)
	if err != nil {
		return nil, nil, err
	}
	// 序列化公钥
	publicKeyBytes, err := x509.MarshalSm2PublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	return privateKeyBytes, publicKeyBytes, nil
}

// Encrypt 加密
func (r *SM2) Encrypt(data string, key []byte) ([]byte, error) {
	publicKeyInterface, err := x509.ParseSm2PublicKey(key)
	if err != nil {
		return nil, err
	}
	result, err := sm2.EncryptAsn1(publicKeyInterface, []byte(data), rand.Reader)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Decrypt 解密
func (r *SM2) Decrypt(data []byte, key []byte, pwd []byte) ([]byte, error) {
	privateKey, err := x509.ParsePKCS8PrivateKey(key, pwd)
	if err != nil {
		return nil, err
	}
	cipherText, err := sm2.DecryptAsn1(privateKey, data)
	if err != nil {
		return nil, err
	}
	return cipherText, nil
}
