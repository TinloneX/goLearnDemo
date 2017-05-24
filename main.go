package main

import (
	//请前往查看routers，import时会执行router.go中的init
	_ "firstDemo/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
