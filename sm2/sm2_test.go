package sm2

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	instance := GetInstance()
	pri, pub, err := instance.GenerateKey(nil)
	if err != nil {
		t.Errorf("%s", err)
	}
	fmt.Printf("private key len: %v\n", len(pri))
	fmt.Printf("private key hex: %v\n", hex.EncodeToString(pri))
	fmt.Printf("private key base64: %v\n", base64.StdEncoding.EncodeToString(pri))
	fmt.Printf("private key: %v\n", string(pri))

	fmt.Printf("public key len: %v\n", len(pub))
	fmt.Printf("public key hex: %v\n", hex.EncodeToString(pub))
	fmt.Printf("public key base64: %v\n", base64.StdEncoding.EncodeToString(pub))
	fmt.Printf("public key: %v\n", string(pub))
}

// 测试 SM2 加解密
func TestEncryptDecrypt1(t *testing.T) {
	instance := GetInstance()
	pri, pub, err := instance.GenerateKey(nil)
	if err != nil {
		t.Errorf("%s", err)
	}
	fmt.Printf("public key hex: %v\n", hex.EncodeToString(pub))
	fmt.Printf("public key base64: %v\n", base64.StdEncoding.EncodeToString(pub))
	fmt.Printf("private key hex: %v\n", hex.EncodeToString(pri))
	fmt.Printf("private key base64: %v\n", base64.StdEncoding.EncodeToString(pri))

	// 使用公钥加密
	data := "bluefrog"
	pubEncrypt, err := instance.Encrypt(data, pub)
	if err != nil {
		t.Errorf("%s", err)
	}
	fmt.Printf("public key encrypt hex: %v\n", hex.EncodeToString(pubEncrypt))
	fmt.Printf("public key encrypt base64: %v\n", base64.StdEncoding.EncodeToString(pubEncrypt))

	// 使用私钥解密
	priDecrypt, err := instance.Decrypt(pubEncrypt, pri, nil)
	if err != nil {
		t.Errorf("%s", err)
	}
	fmt.Printf("private key decrypt hex: %v\n", hex.EncodeToString(priDecrypt))
	fmt.Printf("private key decrypt base64: %v\n", base64.StdEncoding.EncodeToString(priDecrypt))
	fmt.Printf("private key decrypt: %v\n", string(priDecrypt))
}

//func TestSm2(t *testing.T) {
//	priv, err := sm2.GenerateKey(rand.Reader) // 生成密钥对
//	fmt.Println(priv)
//	if err != nil {
//		t.Fatal(err)
//	}
//	fmt.Printf("%v\n", priv.Curve.IsOnCurve(priv.X, priv.Y)) // 验证是否为sm2的曲线
//	pub := &priv.PublicKey
//	msg := []byte("123456")
//	d0, err := pub.EncryptAsn1(msg, rand.Reader)
//	if err != nil {
//		fmt.Printf("Error: failed to encrypt %s: %v\n", msg, err)
//		return
//	}
//	// fmt.Printf("Cipher text = %v\n", d0)
//	d1, err := priv.DecryptAsn1(d0)
//	if err != nil {
//		fmt.Printf("Error: failed to decrypt: %v\n", err)
//	}
//	fmt.Printf("clear text = %s\n", d1)
//	d2, err := sm2.Encrypt(pub, msg, rand.Reader, sm2.C1C2C3)
//	if err != nil {
//		fmt.Printf("Error: failed to encrypt %s: %v\n", msg, err)
//		return
//	}
//	// fmt.Printf("Cipher text = %v\n", d0)
//	d3, err := sm2.Decrypt(priv, d2, sm2.C1C2C3)
//	if err != nil {
//		fmt.Printf("Error: failed to decrypt: %v\n", err)
//	}
//	fmt.Printf("clear text = %s\n", d3)
//	msg, _ = ioutil.ReadFile("ifile")             // 从文件读取数据
//	sign, err := priv.Sign(rand.Reader, msg, nil) // 签名
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	err = ioutil.WriteFile("TestResult", sign, os.FileMode(0644))
//	if err != nil {
//		t.Fatal(err)
//	}
//	signdata, _ := ioutil.ReadFile("TestResult")
//	ok := priv.Verify(msg, signdata) // 密钥验证
//	if ok != true {
//		fmt.Printf("Verify error\n")
//	} else {
//		fmt.Printf("Verify ok\n")
//	}
//	pubKey := priv.PublicKey
//	ok = pubKey.Verify(msg, signdata) // 公钥验证
//	if ok != true {
//		fmt.Printf("Verify error\n")
//	} else {
//		fmt.Printf("Verify ok\n")
//	}
//}
