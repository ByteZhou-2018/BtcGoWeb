package routers

import (
	"BtcGoWeb/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//登录页面展示
	beego.Router("/", &controllers.UsersController{}, "GET:Login")
	beego.Router("/login", &controllers.UsersController{}, "GET:Login")
	//处理注册账号请求 进入登录页面
	beego.Router("/login", &controllers.UsersController{}, "POST:ParseRegister")


	//处理登录请求 进入主页
	beego.Router("/home", &controllers.UsersController{}, "POST:LoginParseForm")
	//beego.Router("/home", &controllers.MainController{}, "POST:LoginParseForm")


	////注册页面展示
	//beego.Router("/register.html", &controllers.MainController{}, "POST:LoginParseForm")
	beego.Router("/register", &controllers.UsersController{}, "GET:Register")
	////处理register页面表单请求
	//beego.Router("/home.html", &controllers.MainController{}, "POST:Home")
	//
	////主页展示
	//beego.Router("/home", &controllers.MainController{}, "POST:HomeParseForm")
	////主页数据处理
	//beego.Router("/register", &controllers.MainController{}, "POST:LoginParseForm")
	//
	//beego.Router("/register", &controllers.MainController{}, "POST:LoginParseForm")
}

//基础路由 get方式
//beego.Get("/", func(context *context.Context) {
//
//})
//固定路由 post请求匹配POST方法 get请求匹配GET方法
//beego.Router("/GetInfo",&controllers.MainController{})

//正则路由 ：全匹配: *    /getUser:id  匹配 getUser/xxx     *.* 匹配 xxx.xxx  /api/:id([0-9]+) 匹配 0~9数字
//beego.Router("/*",&controllers.MainController{})
//自定义路由 ：新添第三个参数，方法映射 "请求方式：处理方法名"
