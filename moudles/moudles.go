package moudles

import "github.com/astaxie/beego/orm"

//用户名结构体
type User struct {
	Id       int    `form:"id" orm:"column(id)" json:"id"`
	Username string `form:"username" orm:"column(user_name)" json:"username"`
	Password string `form:"password" orm:"column(user_pwd)" json:"password"`
}
//注册orm实体模型
func init() {
	orm.RegisterModel(new(User))
}
