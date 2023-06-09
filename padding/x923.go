package padding

import (
	"bytes"
)

/**
## X923Padding
```
填充至符合块大小的整数倍，填充值最后一个字节为填充的数量数，其他字节填 0

# 填充之前:

    0x01 0x02 0x03 0x04

# 填充之后:

    0x01 0x02 0x03 0x04 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0C

    0x0C = 12 需要填充的数量

```
*/

// X923PaddingPack x923 padding pack 处理
func X923PaddingPack(data []byte, blockSize int) ([]byte, error) {
	// 计算需要填补的个数
	padding := blockSize - len(data)%blockSize
	// 如果刚开始已经完整分组了也需要补一整个分组
	if padding == 0 {
		padding = blockSize
	}
	paddingText := bytes.Repeat([]byte{byte(padding)}, padding-1)
	result := append(data, paddingText...)
	// 最后一位是需要填补的数量
	return append(result, byte(padding)), nil
}

// X923PaddingUnPack x923 padding unpack 处理
func X923PaddingUnPack(data []byte) []byte {
	l := len(data)
	removeNum := int(data[l-1])
	return data[:(l - removeNum)]
}
