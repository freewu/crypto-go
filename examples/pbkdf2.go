package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/freewu/crypto-go"
)

func main() {
	password := "admin"
	salt := "12345678"
	keyLength := 32
	hash := ""

	for i := 0; i <= 1024; i++ {
		hash = crypto.PBKDF2(password, salt, i, keyLength, sha256.New)
		fmt.Printf("PBKDF2 iter %4d hash: %v ,len: %v\n", i, hash, len(hash))
	}
	for i := 0; i <= 1024; i++ {
		hash, _ = crypto.PBKDF2MD5(password, salt, i, keyLength)
		fmt.Printf("PBKDF2MD5 iter %4d hash: %v ,len: %v\n", i, hash, len(hash))
	}
	for i := 0; i <= 1024; i++ {
		hash, _ = crypto.PBKDF2SHA1(password, salt, i, keyLength)
		fmt.Printf("PBKDF2SHA1 iter %4d hash: %v ,len: %v\n", i, hash, len(hash))
	}
	for i := 0; i <= 1024; i++ {
		hash, _ = crypto.PBKDF2SHA256(password, salt, i, keyLength)
		fmt.Printf("PBKDF2SHA256 iter %4d hash: %v ,len: %v\n", i, hash, len(hash))
	}
	for i := 0; i <= 1024; i++ {
		hash, _ = crypto.PBKDF2SHA512(password, salt, i, keyLength)
		fmt.Printf("PBKDF2SHA512 iter %4d hash: %v ,len: %v\n", i, hash, len(hash))
	}
}
