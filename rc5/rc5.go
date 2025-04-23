package rc5

import "github.com/lf777/go-rc5"

type RC5 struct {
	Round    uint // 加密轮数
	WordSize uint // 16 / 32 / 64 加密分块大小
}

// GetInstance
// rounds uint // 加密轮数
// wordSize uint // 16  / 32 / 64 加密分块大小
func GetInstance(rounds uint, wordSize uint) *RC5 {
	return &RC5{Round: rounds, WordSize: wordSize}
}

func (r *RC5) Encrypt(data []byte, key []byte) ([]byte, error) {
	cipher, err := rc5.NewCipher(key, r.Round, r.WordSize)
	if err != nil {
		return nil, err
	}
	result := make([]byte, len(data))
	cipher.Encrypt(result, data)
	return result, nil
}

// Decrypt 解密处理
func (r *RC5) Decrypt(data []byte, key []byte) ([]byte, error) {
	cipher, err := rc5.NewCipher(key, r.Round, r.WordSize)
	if err != nil {
		return nil, err
	}
	result := make([]byte, len(data))
	cipher.Decrypt(result, data)
	return result, nil
}

//// Encrypt 加密处理
//func Encrypt(data []byte, key []byte) ([]byte, error) {
//	result := make([]byte, len(data))
//	cipher, err := rc5.NewCipher(key)
//	if err != nil {
//		return nil, err
//	}
//	cipher.XORKeyStream(result, data)
//	return result, nil
//}
//
//// Decrypt 解密处理
//func Decrypt(data []byte, key []byte) ([]byte, error) {
//	result := make([]byte, len(data))
//	cipher, err := rc5.NewCipher(key)
//	if err != nil {
//		return nil, err
//	}
//	cipher.XORKeyStream(result, data)
//	return result, nil
//}
