package padding

import "bytes"

/**
## ISO7816-4Padding
```
填充至符合块大小的整数倍，填充值第一个字节为0x80，其他字节填 0x00

# 填充之前:

    0x01 0x02 0x03 0x04

# 填充之后:

    0x01 0x02 0x03 0x04 0x80 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0

```
*/

// ISO78164PaddingPack iso7816-4 padding pack 处理
func ISO78164PaddingPack(data []byte, blockSize int) ([]byte, error) {
	// 计算需要填补的个数
	padding := blockSize - len(data)%blockSize
	// 如果刚开始已经完整分组了也需要补一整个分组
	if padding == 0 {
		padding = blockSize
	}
	paddingText := bytes.Repeat([]byte{byte(0)}, padding-1)
	result := append(data, byte(0x80)) // 内容结束加上 0x80 后面补 0
	// 最后一位是需要填补的数量
	return append(result, paddingText...), nil
}

// ISO78164PaddingUnPack iso7816-4 padding unpack 处理
func ISO78164PaddingUnPack(data []byte) []byte {
	retsult := bytes.TrimRight(data, string(byte(0)))
	return retsult[:(len(retsult) - 1)] // 去牛内容结尾的 0x80
}
