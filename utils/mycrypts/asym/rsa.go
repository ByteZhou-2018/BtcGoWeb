package asym

import (
	"BtcGoWeb/utils/mycrypts"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"

)
//生成指定长度的rsq私钥
func GenRsaKey(keysize int)  (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, keysize)
}
/**
* 根据给定的私钥数据，生成私钥文件
 */
func GenrsaKeyPairFiles(key *rsa.PrivateKey,filename string) error {
	//根据PCKS1规则，序列化的私钥
	priStream := x509.MarshalPKCS1PrivateKey(key)
	privateFile, err := os.Create("rsa_"+filename+"pri.pem") //new一个存私钥的文件
	if err != nil {
		return err
	}
	pubStream := x509.MarshalPKCS1PublicKey(&key.PublicKey)
	publicFile, err := os.Create("rsa_"+filename+"pub.pem")
	if err != nil {
		return err
	}
	block1 := &pem.Block{
		Type:    " RSA Public Key ",
		Bytes:   pubStream,
	}
	err = pem.Encode(publicFile, block1)
	if err != nil {
		return err
	}
	//pem文件中的格式 结构体   pem: 证书文件后缀
	block := &pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   priStream,
	}
	//调用pem包下的encode函数，将我们定义好的pem文件格式和内容写入到privateFile文件中
	err = pem.Encode(privateFile, block)
	if err != nil {
		return err
	}
	return nil
}
//______________________________________读取pem文件格式的秘钥————————————————--

func ReadPemPriKey(filename string)(*rsa.PrivateKey,error){
	blockBytes,err := ioutil.ReadFile(filename)
	if err != nil {
		return	nil, err
	}
	block,_ :=pem.Decode(blockBytes)
	return  x509.ParsePKCS1PrivateKey(block.Bytes)
}
//______________________________________读取pem文件格式的公钥————————————————--
func  ReadPemPubKey(filename string)(*rsa.PublicKey,error) {
	fileBytes,err := ioutil.ReadFile(filename)
	if err != nil {
		return nil,err
	}
	block,_ := pem.Decode(fileBytes)
	return x509.ParsePKCS1PublicKey(block.Bytes)
}

/*+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++加密
RSA算法公钥对数据进行加密，返回加密的密文
*/
func RSAEncrypt(publicKey *rsa.PublicKey, data []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, data)
}

//RSA算法私钥对密文进行解密，返回解密后的数据
func RSADecrypt(privateKey *rsa.PrivateKey, data []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, data)
}

/**
*RSA私钥对数据进行签名 返回签名
*privateKey rsa私钥指针
*data 要签名的数据
return 签名，error
*/
func RSASign(privatKey *rsa.PrivateKey, data []byte) ([]byte, error) {

	return rsa.SignPKCS1v15(rand.Reader, privatKey, crypto.SHA256, mycrypts.Sha256HashBytes(data))
}

/***
*使用RSA算法对数据进行签名验证，并返回验证签名的结果
	验证通过，返回true，nil
	验证不通过，返回false，error
*publicKey 公钥指针
*data 数据
*signText 签名
*/
func RSAVerify(publicKey *rsa.PublicKey, data, signText []byte) (bool, error) {
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, mycrypts.Sha256HashBytes(data), signText, )
	return err == nil, err
}