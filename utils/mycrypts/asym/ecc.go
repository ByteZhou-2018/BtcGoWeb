package asym
//asym 非对称的

import (
	"BtcGoWeb/utils/mycrypts"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
)

//---------------------生成私钥和公钥的密钥对-------------------------

func GenerateECDSAKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

//-----------------私钥签名，公钥验签----------------------
//______私钥签名__________________________________________-______________-
func ECDSASign(pri *ecdsa.PrivateKey, data []byte) (*big.Int, *big.Int, error) {
	return ecdsa.Sign(rand.Reader, pri, mycrypts.Sha256HashBytes(data))
}

//______公钥验签___________________________________________-_______________-____
func ECDSAVerify(pub ecdsa.PublicKey, r, s *big.Int, data []byte) bool {
	return ecdsa.Verify(&pub, mycrypts.Sha256HashBytes(data), r, s)
}
