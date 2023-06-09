package main

import (
	"fmt"
	"github.com/freewu/crypto-go"
)

func main() {
	// 20238ad293024e2ea2f505db927cd52e
	fmt.Printf("crypto.HmacMD5(\"admin\",\"123456\") = %v\n", crypto.HmacMD5("admin", "123456"))
	// 3c39afa93e0b12c28f1f08b18488ebd4ad2e5858
	fmt.Printf("crypto.HmacSHA1(\"admin\",\"123456\") = %v\n", crypto.HmacSHA1("admin", "123456"))
	// 69c6fda19329b530a43354615c573bf640de9d59a814d8cf3a9760c2e5c614d8
	fmt.Printf("crypto.HmacSHA256(\"admin\",\"123456\") = %v\n", crypto.HmacSHA256("admin", "123456"))
}
