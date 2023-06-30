package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/freewu/crypto-go/xtea"
)

func main() {
	// 加密
	chipertext, err := xtea.Encrypt([]byte("bluefrog"), []byte("1234567812345678"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("tea encrypt hex: %v\n", hex.EncodeToString(chipertext))
	fmt.Printf("tea encrypt base64: %v\n", base64.StdEncoding.EncodeToString(chipertext))
	// 解密
	result, err := xtea.Decrypt(chipertext, []byte("1234567812345678"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("tea decrypt hex: %v\n", hex.EncodeToString(result))
	fmt.Printf("tea decrypt base64: %v\n", base64.StdEncoding.EncodeToString(result))
	fmt.Printf("tea decrypt: %v\n", string(result))
}
