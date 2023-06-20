package tea

import (
	"golang.org/x/crypto/tea"
)

// Encrypt 加密 静态方法 不做填充处理，必须自己保证 plaintext 可以 每组 8 字节等分
func Encrypt(plaintext []byte, key []byte, round int) ([]byte, error) {
	cipher, err := tea.NewCipherWithRounds(key, round)
	if err != nil {
		return nil, err
	}
	var result = make([]byte, len(plaintext))
	cipher.Encrypt(result, plaintext)
	return result, nil
}

// Decrypt 解密 静态方法
func Decrypt(ciphertext []byte, key []byte, round int) ([]byte, error) {
	cipher, err := tea.NewCipherWithRounds(key, round)
	if err != nil {
		return nil, err
	}
	var result = make([]byte, len(ciphertext))
	cipher.Decrypt(result, ciphertext)
	return result, nil
}
