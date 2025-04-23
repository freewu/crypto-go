package main

import (
	"fmt"
	"github.com/freewu/crypto-go"
)

func main() {
	// dc1fd00e3eeeb940ff46f457bf97d66ba7fcc36e0b20802383de142860e76ae6
	fmt.Printf("crypto.SM3(\"admin\") = %v\n", crypto.SM3("admin"))
}
