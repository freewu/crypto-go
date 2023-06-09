package padding

import (
	"bytes"
	"errors"
)

/**
填充至符合块大小的整数倍，填充值为填充数量数
PKCS5Padding 的块大小应为8个字节(PKCS7Padding 的子集)
PKCS7Padding 的块大小可以在1~255的范围内

AES块的大小恰好为8个字节,对于 AES 的 Padding 来说，这两种方式是一样的

PKCS5Padding，PKCS7Padding的子集，块大小固定为8字节。在AES加密当中其实是没有pkcs5的，因为AES的分块是16B而pkcs5只能用于8B，所以我们在AES加密中所说的pkcs5指的就是pkcs7。

如果是数据刚好满足数据块长度也要在元数据后在按PKCS7规则填充一个数据块数据，这样做的目的是为了区分有效数据和补齐数据

# 填充之前:

    0x01 0x02 0x03 0x04

# 填充之后:

    0x01 0x02 0x03 0x04 0x0C 0x0C 0x0C 0x0C 0x0C 0x0C 0x0C 0x0C 0x0C 0x0C 0x0C 0x0C

    0x0C = 12 需要填充的数量
*/

// PKCS7PaddingPack pkcs7 padding pack 处理
func PKCS7PaddingPack(data []byte, blockSize int) ([]byte, error) {
	// KCS7Padding 的块大小可以在 1 ~ 255 的范围内
	if blockSize < 1 && blockSize > 255 {
		return nil, errors.New("blockSize out range")
	}
	// 计算需要填补的个数
	padding := blockSize - len(data)%blockSize
	// 如果刚开始已经完整分组了也需要补一整个分组
	if padding == 0 {
		padding = blockSize
	}
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, paddingText...), nil
}

// PKCS7PaddingUnPack pkcs7 padding unpack 处理
func PKCS7PaddingUnPack(data []byte) []byte {
	l := len(data)
	removeNum := int(data[l-1])
	return data[:(l - removeNum)]
}
