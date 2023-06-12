package padding

/**
## PKCS1Padding
```
全名《Public-Key Cryptography Standards (PKCS) #1:
RSA Cryptography Specifications》最新版本2.2 (rfc8017) ，
针对RSA算法的一个规范。里面包含了RSA加密、解密、签名验签等所有的内容，也包含了私钥的格式。PKCS1的1.1版本是1991年发布的

当你选择RSA_PKCS1_PADDING填充模式时，如果明文不够128字节，加密的时候会在明文中随机填充一些数据，所以会导致对同样的明文每次加密后的结果都不一样。

对加密后的密文，用户使用相同的填充方式都能解密。解密后的明文也就是之前加密的明文。

	EB = 00 || BT || PS || 00 || D ，其中D为消息
	BT（The block type块类型）：
	BT=00 or 01 （私钥运算时）
	BT=02 （公钥运算时）
	PS（The padding string填充字符串）：
	BT=00，PS由00组成；
	BT=01，PS由FF组成；
	BT=02，PS由伪随机生成，且非零；
	PS长度为Len(EB) - 3 - Len(D)，最少是8字节
```
*/
