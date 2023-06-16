package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/ripemd160"
	"testing"
)

func TestPBKDF2(t *testing.T) {
	password := "admin"
	salt := "12345678"
	iter := 1024
	l := 32
	hash := PBKDF2(password, salt, iter, l, md5.New)
	fmt.Printf("md5 hash: %v ,len: %v\n", hash, len(hash)) // d0c1f2effad421667f40d0a8156c9cd5
	if hash != "d0c1f2effad421667f40d0a8156c9cd5" {
		t.Errorf("PBKDF2(%v, %v, %v, %v, md5.New) = %v", password, salt, iter, l, hash)
	}
	hash = PBKDF2(password, salt, iter, 32, sha1.New)
	fmt.Printf("sha1 hash: %v ,len: %v\n", hash, len(hash)) // 56ab671eca5913a7b38d26758e86da75
	if hash != "56ab671eca5913a7b38d26758e86da75" {
		t.Errorf("PBKDF2(%v, %v, %v, %v, sha1.New) = %v", password, salt, iter, l, hash)
	}
	hash = PBKDF2(password, salt, iter, 32, sha256.New)
	fmt.Printf("sha256 hash: %v ,len: %v\n", hash, len(hash)) // dc99ffb0a94c9e848dbfe5b9c6c14b4a
	if hash != "dc99ffb0a94c9e848dbfe5b9c6c14b4a" {
		t.Errorf("PBKDF2(%v, %v, %v, %v, sha256.New) = %v", password, salt, iter, l, hash)
	}
	hash = PBKDF2(password, salt, iter, 32, ripemd160.New)
	fmt.Printf("ripemd160 hash: %v ,len: %v\n", hash, len(hash)) // 7f2699c1f4d6f649a31478c1c2ae6970
	if hash != "7f2699c1f4d6f649a31478c1c2ae6970" {
		t.Errorf("PBKDF2(%v, %v, %v, %v, ripemd160.New) = %v", password, salt, iter, l, hash)
	}
}

func TestPBKDF2EmptySalt(t *testing.T) {
	password := "admin"
	salt := ""
	iter := 1024
	hash := PBKDF2(password, salt, iter, 32, md5.New)
	fmt.Printf("hash: %v ,len: %v\n", hash, len(hash))
}

func TestPBKDF2Iter(t *testing.T) {
	password := "admin"
	salt := ""
	iter := 0
	hash := PBKDF2(password, salt, iter, 32, md5.New)
	fmt.Printf("iter 0 hash: %v ,len: %v\n", hash, len(hash)) // 和 迭代 1 次一样
	//hash, _ = PBKDF2(password, salt, 1, 32, md5.New)
	//fmt.Printf("iter 1 hash: %v ,len: %v\n", hash, len(hash)) // 和 迭代 0 次一样
	//hash, _ = PBKDF2(password, salt, 2, 32, md5.New)
	//fmt.Printf("iter 2 hash: %v ,len: %v\n", hash, len(hash))
	//
	//hash, _ = PBKDF2(password, salt, 5, 32, md5.New)
	//fmt.Printf("iter 5 hash: %v ,len: %v\n", hash, len(hash))
	//
	//hash, _ = PBKDF2(password, salt, 1024, 32, md5.New)
	//fmt.Printf("iter 1024 hash: %v ,len: %v\n", hash, len(hash))

	for i := 0; i <= 1024; i++ {
		hash = PBKDF2(password, salt, i, 32, md5.New)
		fmt.Printf("iter %4d hash: %v ,len: %v\n", i, hash, len(hash))
	}
}

func TestPBKDF2MD5(t *testing.T) {
	password := "admin"
	salt := "12345678"
	keyLength := 32
	hash := ""

	for i := 0; i <= 1024; i++ {
		hash = PBKDF2MD5(password, salt, i, keyLength)
		fmt.Printf("PBKDF2MD5 iter %4d hash: %v ,len: %v\n", i, hash, len(hash))
	}
}

func TestPBKDF2SHA1(t *testing.T) {
	password := "admin"
	salt := "12345678"
	keyLength := 32
	hash := ""

	for i := 0; i <= 1024; i++ {
		hash = PBKDF2SHA1(password, salt, i, keyLength)
		fmt.Printf("PBKDF2SHA1 iter %4d hash: %v ,len: %v\n", i, hash, len(hash))
	}
}

func TestPBKDF2SHA256(t *testing.T) {
	password := "admin"
	salt := "12345678"
	keyLength := 32
	hash := ""

	for i := 0; i <= 1024; i++ {
		hash = PBKDF2SHA256(password, salt, i, keyLength)
		fmt.Printf("PBKDF2SHA256 iter %4d hash: %v ,len: %v\n", i, hash, len(hash))
	}
}

func TestPBKDF2SHA512(t *testing.T) {
	password := "admin"
	salt := "12345678"
	keyLength := 32
	hash := ""

	for i := 0; i <= 1024; i++ {
		hash = PBKDF2SHA512(password, salt, i, keyLength)
		fmt.Printf("PBKDF2SHA512 iter %4d hash: %v ,len: %v\n", i, hash, len(hash))
	}
}
