package models

import (
	_ "fmt"

	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id        int
	Username  string `form:"username"`
	Password  string `form:"password"`
	Avatar    string
	Motto     string
	CreatedAt string
}

func init() {
	orm.RegisterModel(new(Users))
}

// 判断用户是否存在
func IsUserExists(where map[string]interface{}) bool {
	o := orm.NewOrm()
	o.Using("chat")

	needle := o.QueryTable("users")

	for key, value := range where {
		needle = needle.Filter(key, value)
	}

	return needle.Exist()
}
