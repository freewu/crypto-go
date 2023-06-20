package tea

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestKeySize(t *testing.T) {
	instance := GetInstance(64, "NoPadding")
	// 密钥低于 16位 4位
	_, err := instance.Encrypt([]byte("bluefrog"), []byte("1234"))
	if err == nil || string(err.Error()) != "tea: incorrect key size" {
		t.Errorf("need throw exception tea: incorrect key size")
	}
	// 密钥超过16位 17
	_, err = instance.Encrypt([]byte("bluefrog"), []byte("12345678123456789"))
	if err == nil || string(err.Error()) != "tea: incorrect key size" {
		t.Errorf("need throw exception tea: incorrect key size")
	}
	// 密钥为必须 16位
	_, err = instance.Encrypt([]byte("bluefrog"), []byte("1234567812345678"))
	if err != nil {
		t.Errorf("throw exception %v\n", err)
	}
}

func TestEncryptDecrypt(t *testing.T) {
	instance := GetInstance(64, "PKCS5Padding")
	// 加密
	chipertext, err := instance.Encrypt([]byte("bluefrog"), []byte("1234567812345678"))
	if err != nil {
		t.Errorf("throw exception %v\n", err)
	}
	fmt.Printf("tea encrypt hex: %v\n", hex.EncodeToString(chipertext))
	fmt.Printf("tea encrypt base64: %v\n", base64.StdEncoding.EncodeToString(chipertext))
	// 解密
	result, err := instance.Decrypt(chipertext, []byte("1234567812345678"))
	if err != nil {
		t.Errorf("throw exception %v\n", err)
	}
	fmt.Printf("tea decrypt hex: %v\n", hex.EncodeToString(result))
	fmt.Printf("tea decrypt base64: %v\n", base64.StdEncoding.EncodeToString(result))
	fmt.Printf("tea decrypt: %v\n", string(result))

	// 不足 8 字节一组 使用 padding 算法做填充处理
	// 加密
	chipertext, err = instance.Encrypt([]byte("admin"), []byte("1234567812345678"))
	if err != nil {
		t.Errorf("throw exception %v\n", err)
	}
	fmt.Printf("tea encrypt hex: %v\n", hex.EncodeToString(chipertext))
	fmt.Printf("tea encrypt base64: %v\n", base64.StdEncoding.EncodeToString(chipertext))
	// 解密
	result, err = instance.Decrypt(chipertext, []byte("1234567812345678"))
	if err != nil {
		t.Errorf("throw exception %v\n", err)
	}
	fmt.Printf("tea decrypt hex: %v\n", hex.EncodeToString(result))
	fmt.Printf("tea decrypt base64: %v\n", base64.StdEncoding.EncodeToString(result))
	fmt.Printf("tea decrypt: %v\n", string(result))
}

func TestEncryptDecryptNoPadding(t *testing.T) {
	instance := GetInstance(64, "NoPadding")
	// 加密
	chipertext, err := instance.Encrypt([]byte("bluefrog"), []byte("1234567812345678"))
	if err != nil {
		t.Errorf("throw exception %v\n", err)
	}
	fmt.Printf("tea encrypt hex: %v\n", hex.EncodeToString(chipertext))
	fmt.Printf("tea encrypt base64: %v\n", base64.StdEncoding.EncodeToString(chipertext))
	// 解密
	result, err := instance.Decrypt(chipertext, []byte("1234567812345678"))
	if err != nil {
		t.Errorf("throw exception %v\n", err)
	}
	fmt.Printf("tea decrypt hex: %v\n", hex.EncodeToString(result))
	fmt.Printf("tea decrypt base64: %v\n", base64.StdEncoding.EncodeToString(result))
	fmt.Printf("tea decrypt: %v\n", string(result))

	// 不足 8 字节一组 使用 NoPadding 抛出 panic
	chipertext, err = instance.Encrypt([]byte("admin"), []byte("1234567812345678"))
	if err != nil {
		t.Errorf("throw exception %v\n", err)
	}
	fmt.Printf("tea encrypt hex: %v\n", hex.EncodeToString(chipertext))
	fmt.Printf("tea encrypt base64: %v\n", base64.StdEncoding.EncodeToString(chipertext))
}
