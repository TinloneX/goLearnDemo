package controllers

import (
	"github.com/astaxie/beego"
)

//尚未搞懂这是干什么用的，URL构建的流程和执行过程有待验证
type Url4Controller struct {
	beego.Controller
}

func (this *Url4Controller) Get() {

	this.Data["Username"] = "tinlone"
	this.Ctx.Output.Body([]byte("ok"))
}

func (this *Url4Controller) List() {
	this.Ctx.Output.Body([]byte("i am list"))
}
func (this *Url4Controller) Params() {
	var pa string = ""

	for value := range this.Ctx.Input.Params() {
		pa = pa + value
	}

	this.Ctx.Output.Body([]byte(pa))
}

func (this *Url4Controller) Myext() {
	this.Ctx.Output.Body([]byte(this.Ctx.Input.Param(":ext")))
}

func (this *Url4Controller) GetUrl() {
	this.Ctx.Output.Body([]byte(this.URLFor(".Myext")))
}
