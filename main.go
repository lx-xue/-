package main

import (
	_ "beegoProject/database"
	_ "beegoProject/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
