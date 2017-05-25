package routers

import (
	"firstDemo/controllers"

	"github.com/astaxie/beego"
)

//分发由main.go入口处的请求
func init() {
	//第一个参数为路由下的子路由，第二个参数为beego.Controller的子类（即你定义的处理请求的类）
	beego.Router("/", &controllers.MainController{})
	beego.Router("/a", &controllers.AController{})
	beego.Router("/file", &controllers.FileController{})
	beego.Router("/api/list", &controllers.Url4Controller{}, "*:List")
	beego.Router("/person/:last/:first", &controllers.Url4Controller{})
	beego.AutoRouter(&controllers.Url4Controller{})
	beego.Router("/response", &controllers.ResponseController{})
}
