package controllers

import (
	"BtcGoWeb/moudles"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	beeLogger "github.com/beego/bee/logger"
)

type BlanCnotroller struct {
	beego.Controller
}

func (this *BlanCnotroller) Login() {
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
	//orm原生 sql语句查询数据库
	o := orm.NewOrm()
	err = o.Raw("select "+
		"id,user_name,user_pwd,user_status,user_create_time "+
		"form user where user_name = ? and user_pwd = ?", u.Name, u.Password).QueryRow(&u)
	if err != nil {
		this.Ctx.WriteString("您的打开方式错误，请稍后重试！")
		beeLogger.Log.Info(err.Error())
		return
	}
	if u.Id == 0 || u.Name == "" || u.Password == "" || u.Status == 1 || u.TimeStamp == 0 {
		this.Ctx.WriteString("用户名或密码错误，请重试！")
	}
	this.TplName = "home.html"
}
