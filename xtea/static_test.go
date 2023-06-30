package xtea

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/freewu/crypto-go/padding"
	"testing"
)

func TestStaticKeySize(t *testing.T) {
	// 密钥低于 16位 4位
	_, err := Encrypt([]byte("bluefrog"), []byte("1234"))
	if err == nil || string(err.Error()) != "crypto/xtea: invalid key size 4" {
		t.Errorf("need throw exception tea: incorrect key size, error: %v", err.Error())
	}
	// 密钥超过16位 17
	_, err = Encrypt([]byte("bluefrog"), []byte("12345678123456789"))
	if err == nil || string(err.Error()) != "crypto/xtea: invalid key size 17" {
		t.Errorf("need throw exception tea: incorrect key size, error: %v", err.Error())
	}
	// 密钥为必须 16位
	_, err = Encrypt([]byte("bluefrog"), []byte("1234567812345678"))
	if err != nil {
		t.Errorf("throw exception %v\n", err)
	}
}

func TestStaticEncryptDecrypt(t *testing.T) {
	// 加密
	chipertext, err := Encrypt([]byte("bluefrog"), []byte("1234567812345678"))
	if err != nil {
		t.Errorf("throw exception %v\n", err)
	}
	fmt.Printf("xtea encrypt hex: %v\n", hex.EncodeToString(chipertext))
	fmt.Printf("xtea encrypt base64: %v\n", base64.StdEncoding.EncodeToString(chipertext))
	// 解密
	result, err := Decrypt(chipertext, []byte("1234567812345678"))
	if err != nil {
		t.Errorf("throw exception %v\n", err)
	}
	fmt.Printf("xtea decrypt hex: %v\n", hex.EncodeToString(result))
	fmt.Printf("xtea decrypt base64: %v\n", base64.StdEncoding.EncodeToString(result))
	fmt.Printf("xtea decrypt: %v\n", string(result))
}

func TestStaticEncryptDecryptPadding(t *testing.T) {
	// 需要使用自己做填充处理
	// 不足 8 字节分组 抛出 panic : runtime error: index out of range [3] with length 1 [recovered]
	// 加密
	//chipertext, err := Encrypt([]byte("admin"), []byte("1234567812345678"), 64)
	//if err != nil {
	//	t.Errorf("throw exception %v\n", err)
	//}
	//fmt.Printf("xtea encrypt hex: %v\n", hex.EncodeToString(chipertext))
	//fmt.Printf("xtea encrypt base64: %v\n", base64.StdEncoding.EncodeToString(chipertext))
	//// 解密
	//result, err := Decrypt(chipertext, []byte("1234567812345678"), 64)
	//if err != nil {
	//	t.Errorf("throw exception %v\n", err)
	//}
	//fmt.Printf("xtea decrypt hex: %v\n", hex.EncodeToString(result))
	//fmt.Printf("xtea decrypt base64: %v\n", base64.StdEncoding.EncodeToString(result))
	//fmt.Printf("xtea decrypt: %v\n", string(result))

	// 自己做 填充处理
	pack, _ := padding.PKCS5PaddingPack([]byte("admin"), 8)
	// 加密
	chipertext, err := Encrypt(pack, []byte("1234567812345678"))
	if err != nil {
		t.Errorf("throw exception %v\n", err)
	}
	fmt.Printf("xtea encrypt hex: %v\n", hex.EncodeToString(chipertext))
	fmt.Printf("xtea encrypt base64: %v\n", base64.StdEncoding.EncodeToString(chipertext))
	// 解密
	result, err := Decrypt(chipertext, []byte("1234567812345678"))
	if err != nil {
		t.Errorf("throw exception %v\n", err)
	}
	unpack := padding.PKCS5PaddingUnPack(result)
	fmt.Printf("xtea decrypt hex: %v\n", hex.EncodeToString(unpack))
	fmt.Printf("xtea decrypt base64: %v\n", base64.StdEncoding.EncodeToString(unpack))
	fmt.Printf("xtea decrypt: %v\n", string(unpack))
}
