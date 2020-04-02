package models

import (
	_ "fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id        int
	Username  string
	Email     string
	Password  string
	Avatar    string
	CreatedAt time.Time `orm:"auto_now_add;type(timestamp)"`
	UpdatedAt time.Time `orm:"auto_now;type(timestamp)"`
}

func init() {
	orm.RegisterModel(new(User))
}

// 判断用户是否存在
func IsUserExists(filter map[string]interface{}) bool {
	needle := orm.NewOrm().QueryTable("user")

	for key, value := range filter {
		needle = needle.Filter(key, value)
	}

	return needle.Exist()
}

// 获取某个用户的信息
func GetOneUser(filter map[string]interface{}) (User, error) {
	var user User
	needle := orm.NewOrm().QueryTable("user")

	for key, value := range filter {
		needle = needle.Filter(key, value)
	}

	err := needle.One(&user)
	return user, err
}

// 添加用户
func AddUser(data *User) (int64, error) {
	return orm.NewOrm().Insert(data)
}
