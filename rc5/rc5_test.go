package rc5

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
)

// 测试加密轮数 &
//func TestRound(t *testing.T) {
//	instance := GetInstance(0, 16)
//	encrypt, err := instance.Encrypt([]byte("bluefrog"), []byte(""))
//}

func TestEncryptDecrypt(t *testing.T) {
	instance := GetInstance(16, 32)
	encrypted, err := instance.Encrypt([]byte("bluefrog123"), []byte("123456"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Encrypt base64: %v\n", base64.StdEncoding.EncodeToString(encrypted))
	fmt.Printf("Encrypt hex: %v\n", hex.EncodeToString(encrypted))

	decrypted, err := instance.Decrypt(encrypted, []byte("123456"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypt base64: %v\n", base64.StdEncoding.EncodeToString(decrypted))
	fmt.Printf("Decrypt hex: %v\n", hex.EncodeToString(decrypted))
	fmt.Printf("Decrypt string: %v\n", string(decrypted))
}
