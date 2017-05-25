package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
)

type User struct {
	Name string
	Age  int
}

type Dog struct {
	Name  string `form:"name",valid:"Required"`
	Color string `form:"color"`
}

func CheckUser() {
	user := User{"man", 16}
	valid := validation.Validation{}
	valid.Required(user.Name, "name")
	valid.MaxSize(user.Name, 15, "nameMax")
	valid.Range(user.Age, 0, 180, "age")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logs.Info(err.Key, err.Message)
		}

	}

	if v := valid.Max(user.Age, 140, "age"); !v.Ok {
		logs.Info(v.Error.Key, v.Error.Message)
	}
	// 定制错误信息
	minAge := 18
	valid.Min(user.Age, minAge, "age").Message("少儿不宜！")
	// 错误信息格式化
	valid.Min(user.Age, minAge, "age").Message("%d不禁", minAge)
}
