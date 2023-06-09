package padding

import "bytes"

/**
ZeroPadding 意思就是在数据块末尾补 0x00
如果刚开始已经完整分组了也需要补一整个分组的 0x00，否则无法解密

# 填充之前:

    example 1:
    0x01 0x02 0x03 0x04 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00 0x00

    example 2:
    0x01 0x02 0x03 0x04 0x05 0x06 0x07 0x08 0x09 0x0A 0x0B 0x0C 0x0D 0x0E 0x0F 0x01

# 填充之后:

    example 1:
    0x01 0x02 0x03 0x04

    example 2:
    0x01 0x02 0x03 0x04 0x05 0x06 0x07 0x08 0x09 0x0A 0x0B 0x0C 0x0D 0x0E 0x0F 0x01


使用 ZeroPadding 填充时，没办法区分真实数据与填充数据，所以只适合以\0结尾的字符串加解密。
为了改善这种不足，ISO10126Padding采用不足的n-1位补随机数，最后一位补n的做法
*/

// ZeroPaddingPack no padding 处理
func ZeroPaddingPack(data []byte, blockSize int) ([]byte, error) {
	// 计算需要补 0 的个数
	padding := blockSize - len(data)%blockSize
	// 数据长度不对齐时使用0填充，否则不填充
	if padding == 0 {
		return data, nil
	}
	paddingText := bytes.Repeat([]byte{byte(0)}, padding)
	return append(data, paddingText...), nil
}

// ZeroPaddingUnPack no padding 处理
func ZeroPaddingUnPack(data []byte) []byte {
	// 去掉左边的 0x0
	return bytes.TrimRight(data, string(byte(0)))
}
