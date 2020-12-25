package sym

import (
	"crypto/cipher"
	"crypto/des"
)
//对data进行 3des 加密 返回密文 key的大小为24字节
func TripleDesEncrypt(data, key []byte) ([]byte, error) {

	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	data = PKCS5Fill(data, block.BlockSize())

	mode := cipher.NewCBCEncrypter(block, key[:8])
	dst := make([]byte, len(data))
	mode.CryptBlocks(dst, data)
	return dst, nil
}
//对3des加密的密文进行解密 返回明文
func TripleDesDecrypt(data, key []byte) ([]byte, error) {

	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, key[:8])
	dst := make([]byte, len(data))
	mode.CryptBlocks(dst, data)

	dst = PKCS5UnFill(dst,block.BlockSize())
	return dst, nil
}