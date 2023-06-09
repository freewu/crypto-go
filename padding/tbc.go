package padding

import "bytes"

/**
## TBCPadding(Trailling-Bit-Compliment)
```
填充至符合块大小的整数倍，原文最后一位为1时填充0x00，最后一位为0时填充0xFF。

# 填充之前:

    0x01 0x02 0x03 0x04

# 填充之后:

    0x01 0x02 0x03 0x04 0xFF 0xFF 0xFF 0xFF 0xFF 0xFF 0xFF 0xFF 0xFF 0xFF 0xFF 0xFF

    0x04 = 0000 0100 最后一位是 0
```
*/

// TBCPaddingPack tbc padding pack 处理
func TBCPaddingPack(data []byte, blockSize int) ([]byte, error) {
	// 计算需要填补的个数
	padding := blockSize - len(data)%blockSize
	// 如果刚开始已经完整分组了也需要补一整个分组
	if padding == 0 {
		padding = blockSize
	}
	// 判断内容最后1位  0 or 1
	// 奇数：i&1 == 1
	// 偶数：i&1 == 0
	// 原文最后一位为1时填充0x00，最后一位为0时填充0xFF。
	paddingValue := 0x00
	if data[len(data)-1]&1 == 0 {
		paddingValue = 0xFF
	}

	paddingText := bytes.Repeat([]byte{byte(paddingValue)}, padding-1)
	result := append(data, paddingText...)
	// 最后一位是需要填补的数量
	return append(result, byte(padding)), nil
}

// TBCPaddingUnPack tbc padding unpack 处理
func TBCPaddingUnPack(data []byte) []byte {
	return bytes.TrimRight(data, string(data[len(data)-1])) // trim 掉最后个类型
}
