package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/freewu/crypto-go/xxtea"
)

func main() {
	// 加密
	chipertext, err := xxtea.Encrypt([]byte("bluefrog"), []byte("1234567812345678"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("xxtea encrypt hex: %v\n", hex.EncodeToString(chipertext))
	fmt.Printf("xxtea encrypt base64: %v\n", base64.StdEncoding.EncodeToString(chipertext))
	// 解密
	result, err := xxtea.Decrypt(chipertext, []byte("1234567812345678"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("xxtea decrypt hex: %v\n", hex.EncodeToString(result))
	fmt.Printf("xxtea decrypt base64: %v\n", base64.StdEncoding.EncodeToString(result))
	fmt.Printf("xxtea decrypt: %v\n", string(result))
}
