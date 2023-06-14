package rc4

import (
	"crypto/rc4"
)

// Encrypt 加密处理
func Encrypt(data []byte, key []byte) ([]byte, error) {
	result := make([]byte, len(data))
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	cipher.XORKeyStream(result, data)
	return result, nil
}

// Decrypt 解密处理
func Decrypt(data []byte, key []byte) ([]byte, error) {
	result := make([]byte, len(data))
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	cipher.XORKeyStream(result, data)
	return result, nil
}
