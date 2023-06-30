package xxtea

import (
	"github.com/xxtea/xxtea-go/xxtea"
)

// Encrypt 加密 静态方法 不做填充处理
func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	return xxtea.Encrypt(plaintext, key), nil
}

// Decrypt 解密 静态方法
func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	return xxtea.Decrypt(ciphertext, key), nil
}
