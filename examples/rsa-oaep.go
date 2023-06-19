package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/freewu/crypto-go/rsa"
)

func main() {
	instance := rsa.GetInstance("PKCS1")
	// 生成公钥 & 私钥
	pri, pub, err := instance.GenerateKey(1024)
	if err != nil {
		panic(err)
	}
	fmt.Printf("private key: %v\n", string(pri))
	fmt.Printf("public key: %v\n", string(pub))

	// 使用公钥加密
	data := "bluefrog"
	pubEncrypt, err := instance.OAEPEncrypt(data, pub, sha256.New())
	if err != nil {
		panic(err)
	}
	fmt.Printf("OAEP PKCS1 public key encrypt hex: %v\n", hex.EncodeToString(pubEncrypt))
	fmt.Printf("OAEP PKCS1 public key encrypt base64: %v\n", base64.StdEncoding.EncodeToString(pubEncrypt))

	// 使用私钥解密
	priDecrypt, err := instance.OAEPDecrypt(pubEncrypt, pri, sha256.New())
	if err != nil {
		panic(err)
	}
	fmt.Printf("OAEP PKCS1 private key decrypt hex: %v\n", hex.EncodeToString(priDecrypt))
	fmt.Printf("OAEP PKCS1 private key decrypt base64: %v\n", base64.StdEncoding.EncodeToString(priDecrypt))
	fmt.Printf("OAEP PKCS1 private key decrypt: %v\n", string(priDecrypt))

	instance = rsa.GetInstance("PKCS8")
	// 生成公钥 & 私钥
	pri, pub, err = instance.GenerateKey(2048)
	if err != nil {
		panic(err)
	}
	fmt.Printf("private key: %v\n", string(pri))
	fmt.Printf("public key: %v\n", string(pub))

	// 使用公钥加密
	pubEncrypt, err = instance.OAEPEncrypt(data, pub, md5.New())
	if err != nil {
		panic(err)
	}
	fmt.Printf("OAEP PKCS8 public key encrypt hex: %v\n", hex.EncodeToString(pubEncrypt))
	fmt.Printf("OAEP PKCS8 public key encrypt base64: %v\n", base64.StdEncoding.EncodeToString(pubEncrypt))

	// 使用私钥解密
	priDecrypt, err = instance.OAEPDecrypt(pubEncrypt, pri, md5.New())
	if err != nil {
		panic(err)
	}
	fmt.Printf("OAEP PKCS8 private key decrypt hex: %v\n", hex.EncodeToString(priDecrypt))
	fmt.Printf("OAEP PKCS8 private key decrypt base64: %v\n", base64.StdEncoding.EncodeToString(priDecrypt))
	fmt.Printf("OAEP PKCS8 private key decrypt: %v\n", string(priDecrypt))
}
