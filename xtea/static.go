package xtea

import (
	"golang.org/x/crypto/xtea"
)

// Encrypt 加密 静态方法 不做填充处理
func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	cipher, err := xtea.NewCipher(key)
	if err != nil {
		return nil, err
	}
	var result = make([]byte, len(plaintext))
	cipher.Encrypt(result, plaintext)
	return result, nil
}

// Decrypt 解密 静态方法
func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	cipher, err := xtea.NewCipher(key)
	if err != nil {
		return nil, err
	}
	var result = make([]byte, len(ciphertext))
	cipher.Decrypt(result, ciphertext)
	return result, nil
}
