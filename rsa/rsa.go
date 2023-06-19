package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	_ "golang.org/x/crypto/ssh"
	"hash"
)

type RSA struct {
	// PKCS1:  公钥使用  PKCS1 私钥使用 PKCS1
	// PKCS8:  公钥使用  PKCS1 私钥使用 PKCS8
	Type   string // 类型 PKCS1 / PKCS8
	Format string // 格式 pem

}

func GetInstance(t string) *RSA {
	switch t {
	case "PKCS1":
		return &RSA{Type: "PKCS1", Format: "pem"}
	case "PKCS8":
		return &RSA{Type: "PKCS8", Format: "pem"}
	case "OAEP":
		return &RSA{Type: "PKCS8", Format: "pem"}
	}
	return &RSA{Type: "PKCS1", Format: "pem"}
}

// GenerateKey 生成公钥和私钥  要生成的密钥位数 1024 / 2048
func (r *RSA) GenerateKey(bits int) ([]byte, []byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	switch r.Type {
	case "PKCS1":
		// 私钥生成
		// 通过 x509 标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
		X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey) // PKCS1 和 8 是不一致的
		// 使用pem格式对x509输出的内容进行编码
		privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
		priKey := pem.EncodeToMemory(&privateBlock)
		if err != nil {
			return nil, nil, err
		}

		// 公钥生成
		X509PublicKey := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)
		publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
		pubKey := pem.EncodeToMemory(&publicBlock)

		return priKey, pubKey, nil
	case "PKCS8":
		// 私钥生成
		// 通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
		X509PrivateKey, err := x509.MarshalPKCS8PrivateKey(privateKey)
		if err != nil {
			return nil, nil, err
		}
		// 使用 pem 格式对 x509 输出的内容进行编码
		privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
		priKey := pem.EncodeToMemory(&privateBlock)
		if err != nil {
			return nil, nil, err
		}
		// 公钥生成
		X509PublicKey := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)
		publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
		pubKey := pem.EncodeToMemory(&publicBlock)
		return priKey, pubKey, nil
	default:
		return nil, nil, errors.New("no such type")
	}
}

// Encrypt 加密
func (r *RSA) Encrypt(data string, key []byte) ([]byte, error) {
	// pem解码
	block, _ := pem.Decode(key)
	// x509解码
	publicKeyInterface, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	result, err := rsa.EncryptPKCS1v15(rand.Reader, publicKeyInterface, []byte(data))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OAEPEncrypt OAEP 加密
func (r *RSA) OAEPEncrypt(data string, key []byte, hash hash.Hash) ([]byte, error) {
	// pem解码
	block, _ := pem.Decode(key)
	// x509解码
	publicKeyInterface, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	result, err := rsa.EncryptOAEP(hash, rand.Reader, publicKeyInterface, []byte(data), nil)
	return result, nil
}

// Decrypt 解密
func (r *RSA) Decrypt(data []byte, key []byte) ([]byte, error) {
	// pem解码
	block, _ := pem.Decode(key)
	switch r.Type {
	case "PKCS1":
		// x509解码
		privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		cipherText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, data)
		if err != nil {
			return nil, err
		}
		return cipherText, nil
	case "PKCS8":
		privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rsaKey, _ := privateKey.(*rsa.PrivateKey)
		cipherText, err := rsa.DecryptPKCS1v15(rand.Reader, rsaKey, data)
		if err != nil {
			return nil, err
		}
		return cipherText, nil
	}
	return nil, errors.New("so such type")
}

// OAEPDecrypt OAEP 解密
func (r *RSA) OAEPDecrypt(data []byte, key []byte, hash hash.Hash) ([]byte, error) {
	// pem解码
	block, _ := pem.Decode(key)
	switch r.Type {
	case "PKCS1":
		// x509解码
		privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		result, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, data, nil)
		if err != nil {
			return nil, err
		}
		return result, nil
	case "PKCS8":
		privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rsaKey, _ := privateKey.(*rsa.PrivateKey)
		result, err := rsa.DecryptOAEP(hash, rand.Reader, rsaKey, data, nil)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	return nil, errors.New("so such type")
}

//// Encrypt 加密
//func Encrypt(data string, key []byte) ([]byte, error) {
//	// pem解码
//	block, _ := pem.Decode(key)
//	// x509解码
//	//publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
//	publicKeyInterface, err := x509.ParsePKCS1PublicKey(block.Bytes)
//	if err != nil {
//		return nil, err
//	}
//	// 对明文进行加密
//	//cipherText, err := rsa.EncryptOAEP(
//	//	sha256.New(),
//	//	rand.Reader, publicKeyInterface, []byte(data), nil)
//	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKeyInterface, []byte(data))
//	if err != nil {
//		return nil, err
//	}
//	return cipherText, nil
//}
//
//// Decrypt 解密
//func Decrypt(data []byte, key []byte) ([]byte, error) {
//	// pem解码
//	block, _ := pem.Decode(key)
//	// x509解码
//	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
//	//privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
//	if err != nil {
//		return nil, err
//	}
//	cipherText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, data)
//	if err != nil {
//		return nil, err
//	}
//	return cipherText, nil
//}
//
//// GenerateKey 生成公钥和私钥  要生成的密钥位数 1024 / 2048
//func GenerateKey(bits int) ([]byte, []byte, error) {
//	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
//	if err != nil {
//		return nil, nil, err
//	}
//	// 私钥生成
//	// 通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
//	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey) // PKCS1 和 8 是不一致的
//	//X509PrivateKey, err := x509.MarshalPKCS8PrivateKey(privateKey)
//	if err != nil {
//		return nil, nil, err
//	}
//	//X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
//	// 使用pem格式对x509输出的内容进行编码
//	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
//	priKey := pem.EncodeToMemory(&privateBlock)
//
//	// x509.MarshalECPrivateKey()
//	// x509.MarshalPKCS8PrivateKey()
//	// x509.MarshalPKCS1PrivateKey()
//	// x509.MarshalPKCS1PublicKey()
//	// x509.MarshalPKIXPublicKey()
//	if err != nil {
//		return nil, nil, err
//	}
//
//	// 公钥生成
//	X509PublicKey := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)
//	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
//	pubKey := pem.EncodeToMemory(&publicBlock)
//
//	return priKey, pubKey, nil
//}
