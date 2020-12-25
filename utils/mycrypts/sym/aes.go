package sym

import (
	"crypto/aes"
	"crypto/cipher"
)
//对数据data进行aes加密 key长度为16、24、32 对应AES-128、AES-192、AES-256
func AesEncrypt(data,key []byte)([]byte,error)  {
	block,err := aes.NewCipher(key)
	if err != nil {
		return nil,err
	}
	data = PKCS5Fill(data,block.BlockSize())
	mode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	dst := make([]byte, len(data))
	mode.CryptBlocks(dst,data)
	return dst,nil
}
func AesDecrypt(data,key []byte)([]byte,error)  {
	block,err := aes.NewCipher(key)
	if err != nil {
		return nil,err
	}
	mode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	dst := make([]byte, len(data))
	mode.CryptBlocks(dst,data)
	dst = PKCS5UnFill(dst,block.BlockSize())
	return dst,nil
}
