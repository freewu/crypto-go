# crypto-go
 
    封装 golang crypto 提高易用性

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
fmt.Printf("CBC/PKCS7Padding Encrypt base64: %v\n", base64.StdEncoding.EncodeToString(encrypted))
fmt.Printf("CBC/PKCS7Padding Encrypt hex: %v\n", hex.EncodeToString(encrypted))

decrypted, err := instance.Decrypt(encrypted, key, iv)
fmt.Printf("CBC/PKCS7Padding Decrypt base64: %v\n", base64.StdEncoding.EncodeToString(decrypted))
fmt.Printf("CBC/PKCS7Padding Decrypt hex: %v\n", hex.EncodeToString(decrypted))
fmt.Printf("CBC/PKCS7Padding Decrypt string: %v\n", string(decrypted))
```

## AES
```go
import "github.com/freewu/crypto-go/aes"

instance := aes.GetInstance("CBC/PKCS7Padding")
encrypted, err := instance.Encrypt(data, key, iv)
fmt.Printf("CBC/PKCS7Padding Encrypt base64: %v\n", base64.StdEncoding.EncodeToString(encrypted))
fmt.Printf("CBC/PKCS7Padding Encrypt hex: %v\n", hex.EncodeToString(encrypted))

decrypted, err := instance.Decrypt(encrypted, key, iv)
fmt.Printf("CBC/PKCS7Padding Decrypt base64: %v\n", base64.StdEncoding.EncodeToString(decrypted))
fmt.Printf("CBC/PKCS7Padding Decrypt hex: %v\n", hex.EncodeToString(decrypted))
fmt.Printf("CBC/PKCS7Padding Decrypt string: %v\n", string(decrypted))
```

