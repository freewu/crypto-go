package padding

/**
## PKCS1Padding
```
填充格式如下:

    Padding = 00 + BT + PS + 00 + D

    00为固定字节
    BT为处理模式
    PS为填充字节，填充数量为k - 3 - D，k表示密钥长度, D表示原文长度。PS的最小长度为8个字节。填充的值根据BT值来定：
    BT = 00时，填充全00
    BT = 01时，填充全FF
    BT = 02时，随机填充，但不能为00。
```
*/
