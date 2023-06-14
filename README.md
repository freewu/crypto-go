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

crypto.MD5("admin") // 21232f297a57a5a743894a0e4a801fc3
crypto.SHA1("admin") // d033e22ae348aeb5660fc2140aec35850c4da997
crypto.SHA256("admin") // 8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918

crypto.HmacMD5("admin", "123456") // 20238ad293024e2ea2f505db927cd52e
crypto.HmacSHA1("admin", "123456") // 3c39afa93e0b12c28f1f08b18488ebd4ad2e5858
crypto.HmacSHA256("admin", "123456") // 69c6fda19329b530a43354615c573bf640de9d59a814d8cf3a9760c2e5c614d8
```

## DES
```go
import "github.com/freewu/crypto-go/des"

instance := des.GetInstance("CBC/PKCS7Padding")
encrypted, err := instance.Encrypt(data, key, iv)
fmt.Printf("DES/CBC/PKCS7Padding Encrypt base64: %v\n", base64.StdEncoding.EncodeToString(encrypted))
fmt.Printf("DES/CBC/PKCS7Padding Encrypt hex: %v\n", hex.EncodeToString(encrypted))

decrypted, err := instance.Decrypt(encrypted, key, iv)
fmt.Printf("DES/CBC/PKCS7Padding Decrypt base64: %v\n", base64.StdEncoding.EncodeToString(decrypted))
fmt.Printf("DES/CBC/PKCS7Padding Decrypt hex: %v\n", hex.EncodeToString(decrypted))
fmt.Printf("DES/CBC/PKCS7Padding Decrypt string: %v\n", string(decrypted))
```

## 3DES
```go
import "github.com/freewu/crypto-go/desede"

instance := desede.GetInstance("CBC/PKCS7Padding")
encrypted, err := instance.Encrypt(data, key, iv)
fmt.Printf("3DES/CBC/PKCS7Padding Encrypt base64: %v\n", base64.StdEncoding.EncodeToString(encrypted))
fmt.Printf("3DES/CBC/PKCS7Padding Encrypt hex: %v\n", hex.EncodeToString(encrypted))

decrypted, err := instance.Decrypt(encrypted, key, iv)
fmt.Printf("3DES/CBC/PKCS7Padding Decrypt base64: %v\n", base64.StdEncoding.EncodeToString(decrypted))
fmt.Printf("3DES/CBC/PKCS7Padding Decrypt hex: %v\n", hex.EncodeToString(decrypted))
fmt.Printf("3DES/CBC/PKCS7Padding Decrypt string: %v\n", string(decrypted))
```

## AES
```go
import "github.com/freewu/crypto-go/aes"

instance := aes.GetInstance("CBC/PKCS7Padding")
encrypted, err := instance.Encrypt(data, key, iv)
fmt.Printf("AES/CBC/PKCS7Padding Encrypt base64: %v\n", base64.StdEncoding.EncodeToString(encrypted))
fmt.Printf("AES/CBC/PKCS7Padding Encrypt hex: %v\n", hex.EncodeToString(encrypted))

decrypted, err := instance.Decrypt(encrypted, key, iv)
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
encrypted, err := rc4.Encrypt(data, key)
fmt.Printf("Encrypt base64: %v\n", base64.StdEncoding.EncodeToString(encrypted))
fmt.Printf("Encrypt hex: %v\n", hex.EncodeToString(encrypted))

// 解密
decrypted, err := rc4.Decrypt(encrypted, key)
fmt.Printf("Decrypt base64: %v\n", base64.StdEncoding.EncodeToString(decrypted))
fmt.Printf("Decrypt hex: %v\n", hex.EncodeToString(decrypted))
fmt.Printf("Decrypt string: %v\n", string(decrypted))
```

