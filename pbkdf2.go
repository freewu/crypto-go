package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"hash"
)

// PBKDF2
// str 	string 需要 hash 的值
// salt string
// iter int 迭代次数
// keyLen int 返回的 hash 值长度
// h	func() hash.Hash  使用用的 Hash 算法  md5.New / sha1.New / sha256.New
func PBKDF2(str string, salt string, iter int, keyLen int, h func() hash.Hash) string {
	hash := pbkdf2.Key([]byte(str), []byte(salt), iter, keyLen/2, h)
	return fmt.Sprintf("%x", hash)
}

// PBKDF2MD5
// str 	string 需要 hash 的值
// salt string
// iter int 迭代次数 0-1一样 2次起步
// keyLen int 返回的 hash 值长度
func PBKDF2MD5(str string, salt string, iter int, keyLen int) string {
	hash := pbkdf2.Key([]byte(str), []byte(salt), iter, keyLen/2, md5.New)
	return fmt.Sprintf("%x", hash)
}

// PBKDF2SHA256
// str 	string 需要 hash 的值
// salt string
// iter int 迭代次数 0-1一样 2次起步
// keyLen int 返回的 hash 值长度
func PBKDF2SHA256(str string, salt string, iter int, keyLen int) string {
	hash := pbkdf2.Key([]byte(str), []byte(salt), iter, keyLen/2, sha256.New)
	return fmt.Sprintf("%x", hash)
}

// PBKDF2SHA1
// str 	string 需要 hash 的值
// salt string
// iter int 迭代次数 0-1一样 2次起步
// keyLen int 返回的 hash 值长度
func PBKDF2SHA1(str string, salt string, iter int, keyLen int) string {
	hash := pbkdf2.Key([]byte(str), []byte(salt), iter, keyLen/2, sha1.New)
	return fmt.Sprintf("%x", hash)
}

// PBKDF2SHA512
// str 	string 需要 hash 的值
// salt string
// iter int 迭代次数 0-1一样 2次起步
// keyLen int 返回的 hash 值长度
func PBKDF2SHA512(str string, salt string, iter int, keyLen int) string {
	hash := pbkdf2.Key([]byte(str), []byte(salt), iter, keyLen/2, sha512.New)
	return fmt.Sprintf("%x", hash)
}
