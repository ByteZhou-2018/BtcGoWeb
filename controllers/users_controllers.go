package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type UsersController struct {
	beego.Controller
}
//处理 http://localhost:8080/login 请求，处理用户登录,用户名密码都匹配、登录到主页(比特币节点查询功能)
func (this *MainController) LoginParseForm() {
	username := this.GetString("username")
	password := this.GetString("password")
	fmt.Println("hello ," + username + password)
	this.TplName = "home.html"
	//beego.Info("hello world !")
	//c.Ctx.Input.Context.Output.Body([]byte("你好世界"))
}
