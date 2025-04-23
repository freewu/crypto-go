package main

import (
	"fmt"
	"github.com/freewu/crypto-go"
)

func main() {
	// 21232f297a57a5a743894a0e4a801fc3
	fmt.Printf("crypto.MD5(\"admin\") = %v\n", crypto.MD5("admin"))
	// d033e22ae348aeb5660fc2140aec35850c4da997
	fmt.Printf("crypto.SHA1(\"admin\") = %v\n", crypto.SHA1("admin"))
	// 8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918
	fmt.Printf("crypto.SHA256(\"admin\") = %v\n", crypto.SHA256("admin"))
	// dc1fd00e3eeeb940ff46f457bf97d66ba7fcc36e0b20802383de142860e76ae6
	fmt.Printf("crypto.SM3(\"admin\") = %v\n", crypto.SM3("admin"))
}
