package routers

import (
	"firstDemo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/a", &controllers.AController{})
}
