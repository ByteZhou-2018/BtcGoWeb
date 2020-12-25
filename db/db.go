package db

import (
	"database/sql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //导入驱动包
	"io"
	"log"
)

//var O orm.Ormer
/*
* NewOrm 的同时会执行 orm.BootStrap (整个 app 只执行一次)，用以验证模型之间的定义并缓存。
切换数据库，或者，进行事务处理，都会作用于这个 Ormer 对象，以及其进行的任何查询。

所以：需要 切换数据库 和 事务处理 的话，不要使用全局保存的 Ormer 对象。
 */
var Db *sql.DB

//初始化连接mysql
func init() {
	var (
		driverName         = beego.AppConfig.String("DriverName")
		MysqlAuthorization = beego.AppConfig.String("MysqlAuthorization")
		MysqlUrl           = beego.AppConfig.String("MysqlUrl")
		MysqlDatabaseName  = beego.AppConfig.String("MysqlDatabaseName")

		dbConnectStr = MysqlAuthorization + "@tcp(" + MysqlUrl + ")/" + MysqlDatabaseName + "?charset=utf8"
	)

	//1、注册数据库驱动
	err := orm.RegisterDriver(driverName, orm.DRMySQL)
	if err != nil {
		log.Println(err.Error())
		return
	}
	//2、注册数据库连接
	// 参数1        数据库的别名，用来在 ORM 中切换数据库使用
	// 参数2        driverName  驱动名
	// 参数3        数据库连接的字符串 user:password@tcp(ip:port)/databaseName?charset=utf8
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)

	err = orm.RegisterDataBase("default", driverName, dbConnectStr, 30, 30)
	if err != nil {
		log.Println(err.Error())
		return
	}

	//3、通过orm.GetDB 方法 拿到db对象 给全局db对象赋值
	o := orm.NewOrm()
	err = o.Using("default")//默认使用 deault，但可以指定为其它数据库别名用以切换数据库

	//获取orm中注册的default *sql.Db
	db, err := orm.GetDB("default")
	if err != nil {
		log.Println(err.Error())
		return
	}
	Db = db
	//4、调试模式打印查询语句，可能存在性能问题，不建议用在产品模式
	/*
	默认使用 os.Stderr 输出日志信息

	改变输出到你自己的 io.Writer


	var w io.Writer
	...
	// 设置为你的 io.Writer
	...
	orm.DebugLog = orm.NewLog(w)
	 */
	orm.Debug = true
	//
	var w io.Writer

	orm.DebugLog = orm.NewLog(w)
	beego.Info("连接数据成功！")
	/**
	RegisterModel
		orm.RegisterModel(new(User))

	RegisterModel 也可以同时注册多个 model

	orm.RegisterModel(new(User), new(Profile), new(Post))
	将你定义的 Model 进行注册，最佳设计是有单独的 models.go 文件，在他的 init 函数中进行注册。
	 */
}



