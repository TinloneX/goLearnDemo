package controllers

import (
	"encoding/json"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type AController struct {
	beego.Controller
}

func (this *AController) Prepare() {

}

type Book struct {
	Name   string `form:"name"`
	Price  string `form:"price"`
	Number string `form:"number"`
}

func (book *Book) toString() string {
	return "Book:{name=" + book.Name + ",price=" + book.Price + ",number=" + book.Number + "}\n"
}

func (this *AController) Get() {
	var s string
	if err := this.Ctx.Input.Bind(&s, "s"); err != nil {
		log.Println("get s error:", err)
		this.Ctx.WriteString(string(s) + "\n")
	} else {
		log.Println("get [s] = ", s)
		this.Ctx.WriteString(string(s) + "\n")
	}
	this.Ctx.WriteString("This is a interface!\n")
}

/*
将post请求的参数直接解析在struct中
*/
func (this *AController) Post() {
	var book Book
	json.Unmarshal(this.Ctx.Input.RequestBody, &book)
	this.Ctx.WriteString(book.toString())
	logs.Info(book.toString())
}
