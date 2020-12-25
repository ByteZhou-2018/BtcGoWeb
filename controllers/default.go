package controllers

import (
	"BtcGoWeb/btc"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type MainController struct {
	beego.Controller
}

//处理 http://localhost:8080/ 请求，返回登录页面展示
func (c *MainController) Login() {
	c.TplName = "login.html"
}


func (c *MainController) Home() {
	c.TplName = "home.html"

}
func (c *MainController) HomeParseForm() {

	method := c.GetString("method")
	arguments := c.GetString("arguments")


	var result interface{}
	var err error

	switch method {
	case "getblockcount":
		result, err = btc.GetBlockCount()
	case "getblockbesthash":
		result, err = btc.GetBestBlockHash()
	case "getblockchaininfo":
		result,err = btc.GetBlockChainInfo()
	case "getblock":
		result, err = btc.GetBlockByHash(arguments)
	case "getblockhash":
		height, _ := strconv.Atoi(arguments)
		result, err = btc.GetBlockHashByHeight(int64(height))
	case "getblockheader":
		result, err = btc.GetBlockHeaderByHash(arguments)
	default:
		result,err = btc.GetMsgByCommand(method,arguments)

	}
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
	fmt.Println(result)

	c.Data["Result"] = result
	c.TplName = "data.html"
}
