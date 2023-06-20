package tea

import (
	"bytes"
	"errors"
	"github.com/freewu/crypto-go/padding"
	"golang.org/x/crypto/tea"
)

type Tea struct {
	Round   int    // 加轮数,必须为偶数
	Padding string // 填充算法 需要明文 8 字节一组
}

func GetInstance(round int, padding string) *Tea {
	return &Tea{
		//Round:   64,
		//Padding: "PKCS5Padding",
		Round:   round,
		Padding: padding,
	}
}

// Encrypt 加密
func (t *Tea) Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	cipher, err := tea.NewCipherWithRounds(key, t.Round)
	if err != nil {
		return nil, err
	}
	//l := len(plaintext)
	//// 如果能被整除
	//if l%8 == 0 {
	//
	//}
	// 需要明文 8 字节一组 这里使用填充算法做填充处理,保证 8 字节组不抛出异常
	pack, err := t.padding(plaintext, 8)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("plaintext = %x\n", plaintext)
	//fmt.Printf("pack = %x\n", pack)
	var result = make([]byte, len(pack))
	cipher.Encrypt(result, pack)
	//fmt.Printf("result = %x\n", result)
	return result, nil
}

// Decrypt 解密
func (t *Tea) Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	cipher, err := tea.NewCipherWithRounds(key, t.Round)
	if err != nil {
		return nil, err
	}
	var result = make([]byte, len(ciphertext))
	cipher.Decrypt(result, ciphertext)
	unpadding, err := t.unpadding(result)
	if err != nil {
		return nil, err
	}
	// 处理 满 8 填充组 会导致加密后会是 0的问题
	return bytes.TrimRight(unpadding, string(byte(0))), nil
}

// 补码处理
func (t *Tea) padding(data []byte, blockSize int) ([]byte, error) {
	switch t.Padding {
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
func (t *Tea) unpadding(data []byte) ([]byte, error) {
	switch t.Padding {
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
