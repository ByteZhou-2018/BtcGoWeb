package btc

import (
	"BtcGoWeb/moudles"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)
//获取比特币节点中的 rpcUrl 、rpcAuthorization
var rpcUrl= beego.AppConfig.String("RPCURL")
var rpcAuthorization = beego.AppConfig.String("RPCAuthorization")
//获取比特币节点的字符串形式
/*
method： 调用的具体命令
parms：参数
return 将rpc请求所需的数据打包并序列化为json格式
 */
func GetBTCJsonStr(method string, parms []interface{}) string {
	obj := new(moudles.BTCJson)
	obj.Jsonrpc = "2.0"
	obj.Id = strconv.FormatInt(time.Now().Unix(), 10)
	obj.Method = method
	if parms != nil {
		obj.Params = parms
	}
	objStr, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(objStr)
}
//将 json格式的数据 通过rpc请求方式向比特币节点获取节点数据
/*
jsonStr： json格式的请求数据
BTCResult： rpc请求比特币节点的结果集
error :请求数据或解析遇到的错误
 */
func Excute(jsonStr string) (*moudles.BTCResult, error) {
	clinet := &http.Client{}
	req, err := http.NewRequest("POST", rpcUrl, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Encoding", "UTF-8")

	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(rpcAuthorization)))

	response, err := clinet.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var rpcResult moudles.BTCResult
	err = json.Unmarshal(body, &rpcResult)
	if err != nil {
		return nil, err
	}
	return &rpcResult, nil
}
/*btc命令调用封装函数 命令 [参数1，参数2 ...]
	method： 比特币节点具体命令
	parms ：命令对应的具体参数
	return：比特币 Result
 */
func GetMsgByCommand(method string, parms ...interface{}) (*moudles.BTCResult, error) {
	jsonStr := GetBTCJsonStr(method, parms)
	fmt.Println(jsonStr)
	return Excute(jsonStr)
}

