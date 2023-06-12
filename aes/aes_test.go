package aes

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
)

func TestAES(t *testing.T) {
	// 定义明文
	data := []byte("hello world")
	// 密钥
	key := []byte("1234567812345678") // key 必须 16 / 24 / 32 位 128 / 192 /256
	// 偏移量
	iv := []byte("abcdefgh12345678") // 需要 16 / 24 / 32 位 128 / 192 /256
	fmt.Printf("data: %v\n", string(data))
	fmt.Printf("key: %v\n", string(key))
	fmt.Printf("iv: %v\n", string(iv))

	// === CBC模式 Cipher Block Chaining 模式，译为密文分组链接模式
	fmt.Printf("\n===== CBC模式 %v\n", strings.Repeat("=", 20))
	instance := GetInstance("CBC/PKCS7Padding")
	encrypted, err := instance.Encrypt(data, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("CBC/PKCS7Padding Encrypt base64: %v\n", base64.StdEncoding.EncodeToString(encrypted))
	fmt.Printf("CBC/PKCS7Padding Encrypt hex: %v\n", hex.EncodeToString(encrypted))

	decrypted, err := instance.Decrypt(encrypted, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("CBC/PKCS7Padding Decrypt base64: %v\n", base64.StdEncoding.EncodeToString(decrypted))
	fmt.Printf("CBC/PKCS7Padding Decrypt hex: %v\n", hex.EncodeToString(decrypted))
	fmt.Printf("CBC/PKCS7Padding Decrypt string: %v\n", string(decrypted))

	// === CFB模式 Cipher FeedBack 模式，译为密文反馈模式
	fmt.Printf("\n===== CFB模式 %v\n", strings.Repeat("=", 20))
	instance = GetInstance("CFB/PKCS7Padding")
	encrypted, err = instance.Encrypt(data, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("CFB/PKCS7Padding Encrypt base64: %v\n", base64.StdEncoding.EncodeToString(encrypted))
	fmt.Printf("CFB/PKCS7Padding Encrypt hex: %v\n", hex.EncodeToString(encrypted))

	decrypted, err = instance.Decrypt(encrypted, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("CFB/PKCS7Padding Decrypt base64: %v\n", base64.StdEncoding.EncodeToString(decrypted))
	fmt.Printf("CFB/PKCS7Padding Decrypt hex: %v\n", hex.EncodeToString(decrypted))
	fmt.Printf("CFB/PKCS7Padding Decrypt string: %v\n", string(decrypted))

	// === OFB模式 Output Feedback 模式，译为输出反馈模式。
	fmt.Printf("\n===== OFB模式 %v\n", strings.Repeat("=", 20))
	instance = GetInstance("OFB/PKCS7Padding")
	encrypted, err = instance.Encrypt(data, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("OFB/PKCS7Padding Encrypt base64: %v\n", base64.StdEncoding.EncodeToString(encrypted))
	fmt.Printf("OFB/PKCS7Padding Encrypt hex: %v\n", hex.EncodeToString(encrypted))

	decrypted, err = instance.Decrypt(encrypted, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("OFB/PKCS7Padding Decrypt base64: %v\n", base64.StdEncoding.EncodeToString(decrypted))
	fmt.Printf("OFB/PKCS7Padding Decrypt hex: %v\n", hex.EncodeToString(decrypted))
	fmt.Printf("OFB/PKCS7Padding Decrypt string: %v\n", string(decrypted))

	// === CTR模式 Counter 模式，译为计数器模式
	fmt.Printf("\n===== CTR模式 %v\n", strings.Repeat("=", 20))
	instance = GetInstance("CTR/PKCS7Padding")
	encrypted, err = instance.Encrypt(data, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("CTR/PKCS7Padding Encrypt base64: %v\n", base64.StdEncoding.EncodeToString(encrypted))
	fmt.Printf("CTR/PKCS7Padding Encrypt hex: %v\n", hex.EncodeToString(encrypted))

	decrypted, err = instance.Decrypt(encrypted, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("CTR/PKCS7Padding Decrypt base64: %v\n", base64.StdEncoding.EncodeToString(decrypted))
	fmt.Printf("CTR/PKCS7Padding Decrypt hex: %v\n", hex.EncodeToString(decrypted))
	fmt.Printf("CTR/PKCS7Padding Decrypt string: %v\n", string(decrypted))

	// === ECB模式 Electronic Codebook 模式，译为电子密码本模式
	fmt.Printf("\n===== ECB模式 %v\n", strings.Repeat("=", 20))
	instance = GetInstance("ECB/PKCS7Padding")
	encrypted, err = instance.Encrypt(data, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ECB/PKCS7Padding Encrypt base64: %v\n", base64.StdEncoding.EncodeToString(encrypted))
	fmt.Printf("ECB/PKCS7Padding Encrypt hex: %v\n", hex.EncodeToString(encrypted))

	decrypted, err = instance.Decrypt(encrypted, key, iv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ECB/PKCS7Padding Decrypt base64: %v\n", base64.StdEncoding.EncodeToString(decrypted))
	fmt.Printf("ECB/PKCS7Padding Decrypt hex: %v\n", hex.EncodeToString(decrypted))
	fmt.Printf("ECB/PKCS7Padding Decrypt string: %v\n", string(decrypted))
}
