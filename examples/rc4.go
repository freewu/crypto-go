package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/freewu/crypto-go/rc4"
)

func main() {
	// 定义明文
	data := []byte("hello world")
	// 密钥
	key := []byte("12345678")

	fmt.Printf("data: %v\n", string(data))
	fmt.Printf("key: %v\n", string(key))

	encrypted, err := rc4.Encrypt(data, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encrypt base64: %v\n", base64.StdEncoding.EncodeToString(encrypted))
	fmt.Printf("Encrypt hex: %v\n", hex.EncodeToString(encrypted))

	decrypted, err := rc4.Decrypt(encrypted, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypt base64: %v\n", base64.StdEncoding.EncodeToString(decrypted))
	fmt.Printf("Decrypt hex: %v\n", hex.EncodeToString(decrypted))
	fmt.Printf("Decrypt string: %v\n", string(decrypted))
}
