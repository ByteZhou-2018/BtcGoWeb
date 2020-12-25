package sym

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

/**对数据进行des加密 返回加密后的密文
*data 要加密的数据
*key des的密钥，8个字节
return data加密后的密文
*/
func DesEncrypt(data []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	data = PKCS5Fill(data,block.BlockSize())
	mode := cipher.NewCBCEncrypter(block, key)
	dst := make([]byte, len(data))
	mode.CryptBlocks(dst, data)
	return dst, err
}

/**对数据进行des解密 返回解密后的明文
*data 要解密的密文
*key des的密钥，8个字节
return data解密后的明文
*/
func DesDecrypt(data []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, key)
	dst := make([]byte, len(data))
	mode.CryptBlocks(dst, data)
	dst = PKCS5UnFill(dst,block.BlockSize())
	return dst, err
}

/**
*对数据进行 pkcs5 填充
 */
func PKCS5Fill(data []byte, blocksize int) []byte {
	fill := blocksize - len(data)%blocksize
	fillText := bytes.Repeat([]byte{byte(fill)}, fill)
	return append(data, fillText...)
}
func PKCS5UnFill(data []byte, blocksize int) []byte {
	length := len(data)
	fill := int(data[length-1])
	return data[:length-fill]
}

