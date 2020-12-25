package mycrypts

import (
	"crypto/md5"
	"crypto/sha256"
	"github.com/C0MM4ND/go-ripemd"
)

//该函数对msg进行 md5 哈希加密 返回密文
func Md5HashString(msg string) []byte {
	md5hash := md5.New()
	md5hash.Write([]byte(msg))
	return md5hash.Sum(nil)
}

//改函数对 msg进行 sha256哈希，返回密文
func Sha256HashBytes(msg []byte) []byte {
	sha256hash := sha256.New()
	sha256hash.Write(msg)
	return sha256hash.Sum(nil)
}

//该函数对msg进行ripemd160哈希计算，返回密文
func Ripemd160Hash(msg []byte) []byte {
	ripemdHash := ripemd.New160()
	ripemdHash.Write(msg)
	return ripemdHash.Sum(nil)
}

//该函数对msg进行双重sha256哈希
func Sha256HashDouble(msg []byte) []byte {
	sha256hash := sha256.New()
	sha256hash.Write(msg)
	hash1 := sha256hash.Sum(nil)
	sha256hash.Reset()
	sha256hash.Write(hash1)
	return sha256hash.Sum(nil)
}
