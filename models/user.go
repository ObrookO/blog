package models

import (
	_ "fmt"

	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id        int
	Username  string `form:"username" valid:"Required"`
	Password  string `form:"password" valid:"Required"`
	Avatar    string
	Motto     string
	CreatedAt int64
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

// 添加用户
func AddUser(data *Users) (int64, error) {
	o := orm.NewOrm()
	o.Using("chat")

	return o.Insert(data)
}
