# crypto-go
 
    封装 golang crypto 提高易用性

# AES / DES / 3DES 加密模式
```
ECB 电码本模式 Electronic Codebook Book 
CBC 密码分组链接模式 Cipher Block Chaining
CTR 计算器模式 Counter 
CFB 密码反馈模式 Cipher FeedBack
OFB 输出反馈模式 Output FeedBack
```

# AES / DES / 3DES 填充方式
```
NoPadding
ZeroPadding
PKCS7Padding
PKCS5Padding
X923Padding
ISO78164Padding
ISO10126Padding
TBCPadding
```

# 使用
## Hash
```go
import "github.com/freewu/crypto-go"

fmt.Printf("%v\n",crypto.MD5("admin")) // 21232f297a57a5a743894a0e4a801fc3
fmt.Printf("%v\n",crypto.SHA1("admin")) // d033e22ae348aeb5660fc2140aec35850c4da997
fmt.Printf("%v\n",crypto.SHA256("admin")) // 8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918

fmt.Printf("%v\n",crypto.HmacMD5("admin", "123456")) // 20238ad293024e2ea2f505db927cd52e
fmt.Printf("%v\n",crypto.HmacSHA1("admin", "123456")) // 3c39afa93e0b12c28f1f08b18488ebd4ad2e5858
fmt.Printf("%v\n",crypto.HmacSHA256("admin", "123456")) // 69c6fda19329b530a43354615c573bf640de9d59a814d8cf3a9760c2e5c614d8
```

## PBKDF2
```go
import "github.com/freewu/crypto-go"

password := "admin"
salt := "12345678"
keyLength := 32
iter := 1000

fmt.Printf("%v\n", crypto.PBKDF2(password, salt, iter, keyLength, sha256.New)) // e882380cbc438cf60f74a127d3291fe4
fmt.Printf("%v\n", crypto.PBKDF2MD5(password, salt, iter, keyLength)) // 9d47fcd48a33ddedaa9d2912e67c177f
fmt.Printf("%v\n", crypto.PBKDF2SHA1(password, salt, iter, keyLength)) // 40229bb690c4670a1e5110eb740b5020
fmt.Printf("%v\n", crypto.PBKDF2SHA256(password, salt, iter, keyLength)) // e882380cbc438cf60f74a127d3291fe4
fmt.Printf("%v\n", crypto.PBKDF2SHA512(password, salt, iter, keyLength)) // 47d32ed52f2d15fa4128456b5747dd7b
```

## DES
```go
import "github.com/freewu/crypto-go/des"

instance := des.GetInstance("CBC/PKCS7Padding")
encrypted, _ := instance.Encrypt(data, key, iv)
fmt.Printf("DES/CBC/PKCS7Padding Encrypt base64: %v\n", base64.StdEncoding.EncodeToString(encrypted))
fmt.Printf("DES/CBC/PKCS7Padding Encrypt hex: %v\n", hex.EncodeToString(encrypted))

decrypted, _ := instance.Decrypt(encrypted, key, iv)
fmt.Printf("DES/CBC/PKCS7Padding Decrypt base64: %v\n", base64.StdEncoding.EncodeToString(decrypted))
fmt.Printf("DES/CBC/PKCS7Padding Decrypt hex: %v\n", hex.EncodeToString(decrypted))
fmt.Printf("DES/CBC/PKCS7Padding Decrypt string: %v\n", string(decrypted))
```

## 3DES
```go
import "github.com/freewu/crypto-go/desede"

instance := desede.GetInstance("CBC/PKCS7Padding")
encrypted, _ := instance.Encrypt(data, key, iv)
fmt.Printf("3DES/CBC/PKCS7Padding Encrypt base64: %v\n", base64.StdEncoding.EncodeToString(encrypted))
fmt.Printf("3DES/CBC/PKCS7Padding Encrypt hex: %v\n", hex.EncodeToString(encrypted))

decrypted, _ := instance.Decrypt(encrypted, key, iv)
fmt.Printf("3DES/CBC/PKCS7Padding Decrypt base64: %v\n", base64.StdEncoding.EncodeToString(decrypted))
fmt.Printf("3DES/CBC/PKCS7Padding Decrypt hex: %v\n", hex.EncodeToString(decrypted))
fmt.Printf("3DES/CBC/PKCS7Padding Decrypt string: %v\n", string(decrypted))
```

## AES
```go
import "github.com/freewu/crypto-go/aes"

instance := aes.GetInstance("CBC/PKCS7Padding")
encrypted, _ := instance.Encrypt(data, key, iv)
fmt.Printf("AES/CBC/PKCS7Padding Encrypt base64: %v\n", base64.StdEncoding.EncodeToString(encrypted))
fmt.Printf("AES/CBC/PKCS7Padding Encrypt hex: %v\n", hex.EncodeToString(encrypted))

decrypted, _ := instance.Decrypt(encrypted, key, iv)
fmt.Printf("AES/CBC/PKCS7Padding Decrypt base64: %v\n", base64.StdEncoding.EncodeToString(decrypted))
fmt.Printf("AES/CBC/PKCS7Padding Decrypt hex: %v\n", hex.EncodeToString(decrypted))
fmt.Printf("AES/CBC/PKCS7Padding Decrypt string: %v\n", string(decrypted))
```

## RC4
```go
import "github.com/freewu/crypto-go/rc4"

data := []byte("hello world") // 定义明文
key := []byte("12345678") // 密钥

// 加密
encrypted, _ := rc4.Encrypt(data, key)
fmt.Printf("Encrypt base64: %v\n", base64.StdEncoding.EncodeToString(encrypted))
fmt.Printf("Encrypt hex: %v\n", hex.EncodeToString(encrypted))

// 解密
decrypted, _ := rc4.Decrypt(encrypted, key)
fmt.Printf("Decrypt base64: %v\n", base64.StdEncoding.EncodeToString(decrypted))
fmt.Printf("Decrypt hex: %v\n", hex.EncodeToString(decrypted))
fmt.Printf("Decrypt string: %v\n", string(decrypted))
```

