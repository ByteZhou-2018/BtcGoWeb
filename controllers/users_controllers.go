package controllers

import (
	"BtcGoWeb/db"
	"BtcGoWeb/moudles"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	beeLogger "github.com/beego/bee/logger"
	"time"
)

type UsersController struct {
	beego.Controller
}

//处理  http://localhost:8080/  http://localhost:8080/login 请求，返回登录页面展示
func (this *UsersController) Login() {
	this.TplName = "Login.html"
}

//处理 http://localhost:8080/home 请求，处理用户登录,用户名密码都匹配、登录到主页(比特币节点查询功能)
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

//处理进入注册页面在请求  http://localhost:8080/register
func (this *UsersController) Register() {
	this.TplName = "Register.html"
}
func (this *UsersController) ParseRegister() {
	var u moudles.User
	err := this.ParseForm(&u)
	if err != nil {
		beeLogger.Log.Info(err.Error())
		this.Ctx.WriteString("页面解析数据错误，请稍后重试！")
		return
	}
	if u.Name == "" || u.Password == "" {
		this.Ctx.WriteString("请输入用户名和密码")
		return
	}
	beeLogger.Log.Info(u.Name + "     " + u.Password)

	//如果数据中已经该用户，提示：注册失败，用户已存在！
	err = db.O.Raw("select * from user where user_name = ? ", u.Name).QueryRow(&u)
	if err != nil {//没查到 注册用户， 插入数据库
			//this.Ctx.WriteString(err.Error())
		u.TimeStamp = time.Now().Unix()
		u.Status = 0
		_, err = db.O.Insert(&u)
		if err != nil {
			this.Ctx.WriteString("注册用户失败" + fmt.Sprint(err.Error()))
			return
		}
		this.TplName = "Login.html"
		return
	}

	bytes, _ := json.Marshal(&u)
	this.Ctx.WriteString(string(bytes))

	return

}

