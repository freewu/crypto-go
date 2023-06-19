package rsa

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
)

//func TestOrigin(t *testing.T) {
//
//}
//
func TestGenerateKey(t *testing.T) {
	instance := GetInstance("")
	pri, pub, err := instance.GenerateKey(1024)
	if err != nil {
		t.Errorf("%s", err)
	}
	fmt.Printf("1024 private key len: %v\n", len(pri))
	fmt.Printf("1024 private key hex: %v\n", hex.EncodeToString(pri))
	fmt.Printf("1024 private key base64: %v\n", base64.StdEncoding.EncodeToString(pri))
	fmt.Printf("1024 private key: %v\n", string(pri))
	fmt.Printf("1024 public key len: %v\n", len(pub))
	fmt.Printf("1024 public key hex: %v\n", hex.EncodeToString(pub))
	fmt.Printf("1024 public key base64: %v\n", base64.StdEncoding.EncodeToString(pub))
	fmt.Printf("1024 public key: %v\n", string(pub))

	pri, pub, err = instance.GenerateKey(2048)
	if err != nil {
		t.Errorf("%s", err)
	}
	fmt.Printf("2048 private key len: %v\n", len(pri))
	fmt.Printf("2048 private key hex: %v\n", hex.EncodeToString(pri))
	fmt.Printf("2048 private key base64: %v\n", base64.StdEncoding.EncodeToString(pri))
	fmt.Printf("2048 private key: %v\n", string(pri))
	fmt.Printf("2048 public key len: %v\n", len(pub))
	fmt.Printf("2048 public key hex: %v\n", hex.EncodeToString(pub))
	fmt.Printf("2048 public key base64: %v\n", base64.StdEncoding.EncodeToString(pub))
	fmt.Printf("2048 public key: %v\n", string(pub))
}

// 测试加解密
func TestEncryptDecrypt1(t *testing.T) {
	instance := GetInstance("PKCS1")
	pri, pub, err := instance.GenerateKey(1024)
	if err != nil {
		t.Errorf("%s", err)
	}
	// 使用公钥加密
	data := "bluefrog"
	pubEncrypt, err := instance.Encrypt(data, pub)
	if err != nil {
		t.Errorf("%s", err)
	}
	fmt.Printf("PKCS1 public key encrypt hex: %v\n", hex.EncodeToString(pubEncrypt))
	fmt.Printf("PKCS1 public key encrypt base64: %v\n", base64.StdEncoding.EncodeToString(pubEncrypt))

	// 使用私钥解密
	priDecrypt, err := instance.Decrypt(pubEncrypt, pri)
	if err != nil {
		t.Errorf("%s", err)
	}
	fmt.Printf("PKCS1 private key decrypt hex: %v\n", hex.EncodeToString(priDecrypt))
	fmt.Printf("PKCS1 private key decrypt base64: %v\n", base64.StdEncoding.EncodeToString(priDecrypt))
	fmt.Printf("PKCS1 private key decrypt: %v\n", string(priDecrypt))

	instance = GetInstance("PKCS8")
	pri, pub, err = instance.GenerateKey(1024)
	if err != nil {
		t.Errorf("%s", err)
	}
	// 使用公钥加密
	pubEncrypt, err = instance.Encrypt(data, pub)
	if err != nil {
		t.Errorf("%s", err)
	}
	fmt.Printf("PKCS8 public key encrypt hex: %v\n", hex.EncodeToString(pubEncrypt))
	fmt.Printf("PKCS8 public key encrypt base64: %v\n", base64.StdEncoding.EncodeToString(pubEncrypt))

	// 使用私钥解密
	priDecrypt, err = instance.Decrypt(pubEncrypt, pri)
	if err != nil {
		t.Errorf("%s", err)
	}
	fmt.Printf("PKCS8 private key decrypt hex: %v\n", hex.EncodeToString(priDecrypt))
	fmt.Printf("PKCS8 private key decrypt base64: %v\n", base64.StdEncoding.EncodeToString(priDecrypt))
	fmt.Printf("PKCS8 private key decrypt: %v\n", string(priDecrypt))
}

//// 测试加解密
//func TestEncryptDecrypt(t *testing.T) {
//	pri, pub, err := GenerateKey(1024)
//	if err != nil {
//		t.Errorf("%s", err)
//	}
//	// 使用公钥加密
//	data := "bluefrog"
//	pubEncrypt, err := Encrypt(data, pub)
//	if err != nil {
//		t.Errorf("%s", err)
//	}
//	fmt.Printf("1024 public key encrypt hex: %v\n", hex.EncodeToString(pubEncrypt))
//	fmt.Printf("1024 public key encrypt base64: %v\n", base64.StdEncoding.EncodeToString(pubEncrypt))
//
//	// 使用私钥解密
//	priDecrypt, err := Decrypt(pubEncrypt, pri)
//	if err != nil {
//		t.Errorf("%s", err)
//	}
//	fmt.Printf("1024 private key decrypt hex: %v\n", hex.EncodeToString(priDecrypt))
//	fmt.Printf("1024 private key decrypt base64: %v\n", base64.StdEncoding.EncodeToString(priDecrypt))
//	fmt.Printf("1024 private key decrypt: %v\n", string(priDecrypt))
//
//	pri, pub, err = GenerateKey(2048)
//	if err != nil {
//		t.Errorf("%s", err)
//	}
//	// 使用公钥加密
//	data = "bluefrog"
//	pubEncrypt, err = Encrypt(data, pub)
//	if err != nil {
//		t.Errorf("%s", err)
//	}
//	fmt.Printf("2048 public key encrypt hex: %v\n", hex.EncodeToString(pubEncrypt))
//	fmt.Printf("2048 public key encrypt base64: %v\n", base64.StdEncoding.EncodeToString(pubEncrypt))
//
//	// 使用私钥解密
//	priDecrypt, err = Decrypt(pubEncrypt, pri)
//	if err != nil {
//		t.Errorf("%s", err)
//	}
//	fmt.Printf("2048 private key decrypt hex: %v\n", hex.EncodeToString(priDecrypt))
//	fmt.Printf("2048 private key decrypt base64: %v\n", base64.StdEncoding.EncodeToString(priDecrypt))
//	fmt.Printf("2048 private key decrypt: %v\n", string(priDecrypt))
//}
