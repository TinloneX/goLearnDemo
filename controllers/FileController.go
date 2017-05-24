package controllers

import (
	"log"
	"strconv"

	"github.com/astaxie/beego"
)

type FileController struct {
	beego.Controller
}

func (this *FileController) Prepare() {

}

//通过数据绑定的方式获取Get请求中的参数
func (this *FileController) Get() {

	var id int
	this.Ctx.Input.Bind(&id, "id")

	var user string
	this.Ctx.Input.Bind(&user, "user")

	this.Ctx.WriteString(strconv.Itoa(id) + user)
}

/*
获取上传的文件，文件标记参数为‘file’
示例：---------------------------acebdf13572468
Content-Disposition: form-data; name="（此处为标记）file"; filename="liteidex31.windows-qt4.zip"
Content-Type: application/octet-stream

<@INCLUDE *E:\liteidex31.windows-qt4.zip*@>
---------------------------acebdf13572468--
*/
func (this *FileController) Post() {
	f, h, err := this.GetFile("file")
	defer f.Close()
	if err != nil {
		log.Println(err)
	} else {
		this.SaveToFile("file", "static/upload/"+h.Filename)
		this.Ctx.WriteString("接收文件成功")
	}
}
