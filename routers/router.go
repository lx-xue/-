package routers

import (
	"beegoProject/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/index/:id/:name", &controllers.UserController{}, "get:Index") //参数id
	beego.Router("/user/adduser", &controllers.UserController{}, "post:AddUser")
}
