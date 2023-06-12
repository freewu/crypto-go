package des

import (
	"crypto/cipher"
	"crypto/des"
	"errors"
	"github.com/freewu/crypto-go/padding"
	"strings"
)

type DES struct {
	Mode    string // 加密模式  ECB / CBC / CFB / OFB / CTR  default: ECB
	Padding string // 填充模式 NoPadding / ZeroPadding / PKCS7 / PKCS5
}

// GetInstance 实例化一个 des 对象  如: GetInstance("ECB/PKCS7Padding")
func GetInstance(str string) *DES {
	if str != "" {
		arr := strings.Split(str, "/")
		// todo 需要验证处理
		if len(arr) == 1 {
			return &DES{Mode: arr[0], Padding: "PKCS7Padding"}
		}
		return &DES{Mode: arr[0], Padding: arr[1]}
	}
	return &DES{Mode: "ECB", Padding: "PKCS7Padding"}
}

// 补码处理
func (d *DES) padding(data []byte, blockSize int) ([]byte, error) {
	switch d.Padding {
	case "NoPadding":
		return padding.NoPaddingPack(data, blockSize)
	case "ZeroPadding":
		return padding.ZeroPaddingPack(data, blockSize)
	case "PKCS7Padding":
		return padding.PKCS7PaddingPack(data, blockSize)
	case "PKCS5Padding":
		return padding.PKCS5PaddingPack(data, blockSize)
	case "X923Padding":
		return padding.X923PaddingPack(data, blockSize)
	case "ISO78164Padding":
		return padding.ISO78164PaddingPack(data, blockSize)
	case "ISO10126Padding":
		return padding.ISO10126PaddingPack(data, blockSize)
	case "TBCPadding":
		return padding.TBCPaddingPack(data, blockSize)
	}
	return nil, errors.New("no such padding algorithm")
}

// 去掉补码处理
func (d *DES) unpadding(data []byte) ([]byte, error) {
	switch d.Padding {
	case "NoPadding":
		return padding.NoPaddingUnPack(data), nil
	case "ZeroPadding":
		return padding.ZeroPaddingUnPack(data), nil
	case "PKCS7Padding":
		return padding.PKCS7PaddingUnPack(data), nil
	case "PKCS5Padding":
		return padding.PKCS5PaddingUnPack(data), nil
	case "X923Padding":
		return padding.X923PaddingUnPack(data), nil
	case "ISO78164Padding":
		return padding.ISO78164PaddingUnPack(data), nil
	case "ISO10126Padding":
		return padding.ISO10126PaddingUnPack(data), nil
	case "TBCPadding":
		return padding.TBCPaddingUnPack(data), nil
	}
	return nil, errors.New("no such unpadding algorithm")
}

// Encrypt 加密处理
func (d *DES) Encrypt(data []byte, key []byte, iv []byte) ([]byte, error) {
	// 将字节秘钥转换成block快
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 补码处理
	data, err = d.padding(data, block.BlockSize())
	if err != nil {
		return nil, err
	}
	// 创建明文长度的字节数组
	encrypted := make([]byte, len(data))

	switch d.Mode {
	case "CBC":
		blockMode := cipher.NewCBCEncrypter(block, iv)
		// 加密明文,加密后的数据放到数组中
		blockMode.CryptBlocks(encrypted, data)
		return encrypted, nil
	case "ECB": // go 的 DES 默认隐藏了 ECB模式, 因为go认为ECB不安全, 所以不建议使用,就隐藏了
		bs := block.BlockSize()
		dst := encrypted
		for len(data) > 0 {
			//对明文按照blocksize进行分块加密
			//必要时可以使用go关键字进行并行加密
			block.Encrypt(dst, data[:bs])
			data = data[bs:]
			dst = dst[bs:]
		}
		return encrypted, nil
	case "CFB":
		blockMode := cipher.NewCFBEncrypter(block, iv)
		// 加密明文,加密后的数据放到数组中
		blockMode.XORKeyStream(encrypted, data)
		return encrypted, nil
	case "OFB":
		blockMode := cipher.NewOFB(block, iv)
		blockMode.XORKeyStream(encrypted, data)
		return encrypted, nil
	case "CTR":
		blockMode := cipher.NewCTR(block, iv)
		blockMode.XORKeyStream(encrypted, data)
		return encrypted, nil
	}

	return nil, errors.New("no such mode")
}

// Decrypt 解密处理
func (d *DES) Decrypt(data []byte, key []byte, iv []byte) ([]byte, error) {
	// 将字节秘钥转换成block快
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 传入的解密数据不是满块
	if len(data)%block.BlockSize() != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}

	// 创建明文长度的字节数组
	decrypted := make([]byte, len(data))

	switch d.Mode {
	case "CBC":
		blockMode := cipher.NewCBCDecrypter(block, iv)
		// 加密明文,加密后的数据放到数组中
		blockMode.CryptBlocks(decrypted, data)
		return d.unpadding(decrypted)
	case "ECB":
		bs := block.BlockSize()
		dst := decrypted
		for len(data) > 0 {
			block.Decrypt(dst, data[:bs])
			data = data[bs:]
			dst = dst[bs:]
		}
		return d.unpadding(decrypted)
	case "CFB":
		blockMode := cipher.NewCFBDecrypter(block, iv)
		// 加密明文,加密后的数据放到数组中
		blockMode.XORKeyStream(decrypted, data)
		return d.unpadding(decrypted)
	case "OFB":
		blockMode := cipher.NewOFB(block, iv)
		blockMode.XORKeyStream(decrypted, data)
		return d.unpadding(decrypted)
	case "CTR":
		blockMode := cipher.NewCTR(block, iv)
		blockMode.XORKeyStream(decrypted, data)
		return d.unpadding(decrypted)
	}

	return nil, errors.New("no such mode")
}
