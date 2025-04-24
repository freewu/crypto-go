package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/freewu/crypto-go/sm2"
)

func main() {
	instance := sm2.GetInstance()
	pri, pub, err := instance.GenerateKey(nil)
	if err != nil {
		fmt.Errorf("生成秘钥失败: %s", err)
	}
	fmt.Printf("public key hex: %v\n", hex.EncodeToString(pub))
	fmt.Printf("public key base64: %v\n", base64.StdEncoding.EncodeToString(pub))
	fmt.Printf("private key hex: %v\n", hex.EncodeToString(pri))
	fmt.Printf("private key base64: %v\n", base64.StdEncoding.EncodeToString(pri))

	// 使用公钥加密
	data := "bluefrog"
	pubEncrypt, err := instance.Encrypt(data, pub)
	if err != nil {
		fmt.Errorf("公钥加密失败: %s", err)
	}
	fmt.Printf("public key encrypt hex: %v\n", hex.EncodeToString(pubEncrypt))
	fmt.Printf("public key encrypt base64: %v\n", base64.StdEncoding.EncodeToString(pubEncrypt))

	// 使用私钥解密
	priDecrypt, err := instance.Decrypt(pubEncrypt, pri, nil)
	if err != nil {
		fmt.Errorf("私钥解密失败: %s", err)
	}
	fmt.Printf("private key decrypt hex: %v\n", hex.EncodeToString(priDecrypt))
	fmt.Printf("private key decrypt base64: %v\n", base64.StdEncoding.EncodeToString(priDecrypt))
	fmt.Printf("private key decrypt: %v\n", string(priDecrypt))
}
