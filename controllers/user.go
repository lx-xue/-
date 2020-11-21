package controllers

import (
	"github.com/astaxie/beego"
)

//定义控制器(声明结构体)
type UserController struct {
	beego.Controller
}

func (user *UserController) Index() {
	id := user.Ctx.Input.Param(":id")
	name := user.Ctx.Input.Param(":name")

	//输出相应
	mysqluser := "数据库用户名" + beego.AppConfig.String("mysqluser") + "\n"
	mysqlpass := "数据库用户密码" + beego.AppConfig.String("mysqlpass") + "\n"
	mysqlurls := "数据库主机" + beego.AppConfig.String("mysqlurls") + "\n"
	mysqldb := "数据库名" + beego.AppConfig.String("mysqldb") + "\n"
	mysql := mysqluser + mysqlpass + mysqlurls + mysqldb
	//user.Ctx.WriteString("数据库配置:\n" + mysql + "\n" + id + name)
	user.Data["id"] = id
	user.Data["name"] = name
	user.Data["mysql"] = mysql
	user.TplName = "user/index.html"
}

//添加用户
func (user *UserController) AddUser() {
	sex := user.GetString("sex")
	height := user.GetString("height")
	weight := user.GetString("weight")
	user.Ctx.WriteString("post表单提交成功:\n" + sex + height + weight)
}
