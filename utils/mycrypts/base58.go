package mycrypts

import (
	"bytes"
	"math/big"
)

/*
* Base58 是一种基于文本的二进制编码方式。这种编码方式不仅实现了数据压缩，保持了易读性，还具有错误诊断功能
*
* Base58 是Base64的子集 同样使用大小写字母和10个数字，
*  但舍弃了数字 0、大写字母 O 、小写字母 l 、大写字母 I 以及 + 和 / 两个字符。
*
*
*
*
*
*
*

 */
//base58的字母表
const BASE58ALPHABET = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

var base58Alphabet = []byte(BASE58ALPHABET)

func Base58Encode(input []byte) []byte {
	var result []byte
	x := big.NewInt(0).SetBytes(input)             //输入的字符
	base := big.NewInt(int64(len(base58Alphabet))) //基数 58
	zero := big.NewInt(0)
	mod := &big.Int{} //余数
	//不断将数值对58取模，如果商大于58、则对商继续取模
	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, base58Alphabet[mod.Int64()])
	}
	//https://en.bitcon.it/wiki/Base58Check_encoding#Version_bytes
	if input[0] == 0x00 {
		result = append(result, base58Alphabet[0])
	}
	ReverseBytes(result)
	return result
}
//base58解码
func Base58Decode(input []byte)[]byte  {
	result :=big.NewInt(0)
	for _,b:= range input{
		charIndex := bytes.IndexByte(base58Alphabet,b)
		result.Mul(result,big.NewInt(58))
		result.Add(result,big.NewInt(int64(charIndex)))
	}
	decoded:= result.Bytes()
	if input[0] == base58Alphabet[0] {
		decoded = append([]byte{0x00},decoded...)
	}
	return decoded
}
//字节逆转
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i],data[j] = data[j],data[i]
	}
}
