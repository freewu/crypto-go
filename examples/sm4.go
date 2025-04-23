package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/freewu/crypto-go/sm4"
	"strings"
)

func main() {
	// 定义明文
	data := []byte("hello world")
	// 密钥
	key := []byte("1234567812345678") // key 必须 16
	// 偏移量
	iv := []byte("abcdefgh12345678") // 需要 16
	fmt.Printf("data: %v\n", string(data))
	fmt.Printf("key: %v\n", string(key))
	fmt.Printf("iv: %v\n", string(iv))

	// === CBC模式 Cipher Block Chaining 模式，译为密文分组链接模式
	fmt.Printf("\n===== CBC模式 %v\n", strings.Repeat("=", 20))
	instance := sm4.GetInstance("CBC")
	encrypted, err := instance.Encrypt(data, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("CBC Encrypt base64: %v\n", base64.StdEncoding.EncodeToString(encrypted))
	fmt.Printf("CBC Encrypt hex: %v\n", hex.EncodeToString(encrypted))

	decrypted, err := instance.Decrypt(encrypted, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("CBC Decrypt base64: %v\n", base64.StdEncoding.EncodeToString(decrypted))
	fmt.Printf("CBC Decrypt hex: %v\n", hex.EncodeToString(decrypted))
	fmt.Printf("CBC Decrypt string: %v\n", string(decrypted))

	// === CFB模式 Cipher FeedBack 模式，译为密文反馈模式
	fmt.Printf("\n===== CFB模式 %v\n", strings.Repeat("=", 20))
	instance = sm4.GetInstance("CFB")
	encrypted, err = instance.Encrypt(data, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("CFB Encrypt base64: %v\n", base64.StdEncoding.EncodeToString(encrypted))
	fmt.Printf("CFB Encrypt hex: %v\n", hex.EncodeToString(encrypted))

	decrypted, err = instance.Decrypt(encrypted, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("CFB Decrypt base64: %v\n", base64.StdEncoding.EncodeToString(decrypted))
	fmt.Printf("CFB Decrypt hex: %v\n", hex.EncodeToString(decrypted))
	fmt.Printf("CFB Decrypt string: %v\n", string(decrypted))

	// === OFB模式 Output Feedback 模式，译为输出反馈模式。
	fmt.Printf("\n===== OFB模式 %v\n", strings.Repeat("=", 20))
	instance = sm4.GetInstance("OFB")
	encrypted, err = instance.Encrypt(data, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("OFB Encrypt base64: %v\n", base64.StdEncoding.EncodeToString(encrypted))
	fmt.Printf("OFB Encrypt hex: %v\n", hex.EncodeToString(encrypted))

	decrypted, err = instance.Decrypt(encrypted, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("OFB Decrypt base64: %v\n", base64.StdEncoding.EncodeToString(decrypted))
	fmt.Printf("OFB Decrypt hex: %v\n", hex.EncodeToString(decrypted))
	fmt.Printf("OFB Decrypt string: %v\n", string(decrypted))

	// === ECB模式 Electronic Codebook 模式，译为电子密码本模式
	fmt.Printf("\n===== ECB模式 %v\n", strings.Repeat("=", 20))
	instance = sm4.GetInstance("ECB")
	encrypted, err = instance.Encrypt(data, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ECB Encrypt base64: %v\n", base64.StdEncoding.EncodeToString(encrypted))
	fmt.Printf("ECB Encrypt hex: %v\n", hex.EncodeToString(encrypted))

	decrypted, err = instance.Decrypt(encrypted, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ECB Decrypt base64: %v\n", base64.StdEncoding.EncodeToString(decrypted))
	fmt.Printf("ECB Decrypt hex: %v\n", hex.EncodeToString(decrypted))
	fmt.Printf("ECB Decrypt string: %v\n", string(decrypted))
}
