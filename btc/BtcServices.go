package btc

import (
	"BtcGoWeb/moudles"
	"github.com/mitchellh/mapstructure"
)
//比特币节点命令 getblockcount 的封装函数
func GetBlockCount() (interface{}, error) {
	result, err := GetMsgByCommand("getblockcount")
	if err != nil {
		return "", err
	}
	return result.Result, err
}
//比特币节点命令 getbestblockhash 的封装函数
func GetBestBlockHash() (interface{}, error) {
	result, err := GetMsgByCommand("getbestblockhash")
	if err != nil {
		return "", err
	}
	return result.Result, err
}

//比特币节点命令 getblockhash 的封装函数
func GetBlockHashByHeight(height int64) (interface{}, error) {
	result, err := GetMsgByCommand("getblockhash",height)
	if err != nil {
		return "", err
	}
	return result.Result, err
}

//比特币节点命令 getblock 的封装函数
func GetBlockByHash(hash string) (*moudles.Blcok,error) {
	result, err:= GetMsgByCommand("getblock", hash)
	if err != nil {
		return nil, err
	}
	var block moudles.Blcok
	err = mapstructure.Decode(result.Result, &block)
	if err != nil {
		return nil,err
	}
	return &block, nil
}

//比特币节点命令 getblockchaininfo 的封装函数
func GetBlockChainInfo()(*moudles.BlockChainInfo,error)  {
	result, err := GetMsgByCommand("getblockchaininfo")
	if err != nil {
		return nil,err
	}
	var blockchain moudles.BlockChainInfo
	err = mapstructure.Decode(result.Result, &blockchain)
	if err != nil{
		return nil,err
	}
	return &blockchain,nil
}

//比特币节点命令 getblockheader 的封装函数
func GetBlockHeaderByHash(hash string) (*moudles.BlockHeader,error) {
	result, err:= GetMsgByCommand("getblockheader", hash)
	if err != nil {
		return nil, err
	}
	var block moudles.BlockHeader
	err = mapstructure.Decode(result.Result, &block)
	if err != nil {
		return nil,err
	}
	return &block, nil
}
