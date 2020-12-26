package controllers

import (
	"BtcGoWeb/moudles"
	"github.com/astaxie/beego"
	beeLogger "github.com/beego/bee/logger"
)

type UsersController struct {
	beego.Controller
}

//处理 http://localhost:8080//home 请求，处理用户登录,用户名密码都匹配、登录到主页(比特币节点查询功能)
func (this *UsersController) LoginParseForm() {
	var u moudles.User
	err := this.ParseForm(&u)
	if err != nil {
		beeLogger.Log.Info(err.Error())
		this.Ctx.WriteString("页面解析数据错误，请稍后重试！")
		return
	}
	if u.Name == "" || u.Password == "" {
		this.Ctx.WriteString("请输入用户名和密码")
	}
	beeLogger.Log.Info(u.Name + "     " + u.Password)
	//orm原生 sql语句查询数据库
	//o := orm.NewOrm()
	//var user moudles.User
	//rs := o.Raw("select * "+
	//	//"id,user_name,user_pwd,user_status,user_create_time "+
	//	"form users where user_name = ? and user_pwd = ?", u.Name, u.Password)
	//err = rs.QueryRow(&user)
	//if err != nil {
	//	this.Ctx.WriteString("您的打开方式错误，请稍后重试！")
	//	beeLogger.Log.Info(err.Error())
	//	return
	//}
	//fmt.Println(user)
	//if u.Id == 0 || u.Name == "" || u.Password == "" || u.Status == 1 || u.TimeStamp == 0 {
	//	this.Ctx.WriteString("用户名或密码错误，请重试！")
	//}
	this.TplName = "home.html"
	//beego.Info("hello world !")
	//c.Ctx.Input.Context.Output.Body([]byte("你好世界"))
}
func (this *UsersController) Register() {
	this.TplName = "Register.html"
}
