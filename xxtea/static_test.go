package xxtea

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestStaticEncryptDecrypt(t *testing.T) {
	// 加密
	chipertext, err := Encrypt([]byte("bluefrog"), []byte("1234567812345678"))
	if err != nil {
		t.Errorf("throw exception %v\n", err)
	}
	fmt.Printf("xxtea encrypt hex: %v\n", hex.EncodeToString(chipertext))
	fmt.Printf("xxtea encrypt base64: %v\n", base64.StdEncoding.EncodeToString(chipertext))
	// 解密
	result, err := Decrypt(chipertext, []byte("1234567812345678"))
	if err != nil {
		t.Errorf("throw exception %v\n", err)
	}
	fmt.Printf("xxtea decrypt hex: %v\n", hex.EncodeToString(result))
	fmt.Printf("xxtea decrypt base64: %v\n", base64.StdEncoding.EncodeToString(result))
	fmt.Printf("xxtea decrypt: %v\n", string(result))
}
