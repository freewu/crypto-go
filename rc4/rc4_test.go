package rc4

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestRC4(t *testing.T) {
	// 定义明文
	data := []byte("hello world")
	// 密钥
	key := []byte("12345678")

	fmt.Printf("data: %v\n", string(data))
	fmt.Printf("key: %v\n", string(key))

	encrypted, err := Encrypt(data, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encrypt base64: %v\n", base64.StdEncoding.EncodeToString(encrypted))
	fmt.Printf("Encrypt hex: %v\n", hex.EncodeToString(encrypted))

	decrypted, err := Decrypt(encrypted, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypt base64: %v\n", base64.StdEncoding.EncodeToString(decrypted))
	fmt.Printf("Decrypt hex: %v\n", hex.EncodeToString(decrypted))
	fmt.Printf("Decrypt string: %v\n", string(decrypted))
}

func TestRC4EmptyKey(t *testing.T) {
	// 定义明文
	data := []byte("hello world")
	// 密钥
	key := []byte("")

	fmt.Printf("data: %v\n", string(data))
	fmt.Printf("key: %v\n", string(key))

	// key 必须为  1 - 256 byte ( 8 - 2048 bit)
	_, err := Encrypt(data, key)
	if err == nil || err.Error() != "crypto/rc4: invalid key size 0" {
		t.Errorf("need throw invalid key size 0 error")
	}
	fmt.Printf("error: %v", err)
}

func TestRC4LargeKey(t *testing.T) {
	// 定义明文
	data := []byte("hello world")
	// 密钥
	key := []byte("12345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678") // 33 字符 32 字符 = 256 位

	fmt.Printf("data: %v\n", string(data))
	fmt.Printf("key: %v\n", string(key))

	// key 必须为  1 - 256 byte ( 8 - 2048 bit)
	_, err := Encrypt(data, key)
	if err == nil || err.Error() != "crypto/rc4: invalid key size 320" {
		t.Errorf("need throw crypto/rc4: invalid key size 320")
	}
	fmt.Printf("error: %v", err)
}
