package controllers

import (
	"firstDemo/models"

	"github.com/astaxie/beego"
)

type ResponseController struct {
	beego.Controller
}

//猫咪
type Cat struct {
	Name  string
	Color string
	age   string
}

//初始化一只猫咪
var cat = Cat{Name: "Tom", Color: "black", age: "3"}

//使用JSON格式返回
func (this *ResponseController) Get() {
	models.CheckUser()

	this.Data["json"] = &cat
	this.ServeJSON()
}

//使用XML格式返回
func (this *ResponseController) Post() {
	this.Data["xml"] = &cat
	this.ServeXML()
}
