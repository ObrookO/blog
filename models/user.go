package models

import (
	_ "fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id           int
	Username     string    // 用户名
	Email        string    // 邮箱
	Password     string    // 密码
	Avatar       string    // 头像
	AllowComment uint8     // 是否允许评论
	Status       uint8     // 状态
	CreatedAt    time.Time `orm:"auto_now_add;type(timestamp)"`
	UpdatedAt    time.Time `orm:"auto_now;type(timestamp)"`
}

func init() {
	orm.RegisterModelWithPrefix("home_", new(User))
}

// IsUserExists 判断用户是否存在
func IsUserExists(filter map[string]interface{}) bool {
	needle := orm.NewOrm().QueryTable("home_user")
	for key, value := range filter {
		needle = needle.Filter(key, value)
	}

	return needle.Exist()
}

// AddUser 添加用户
func AddUser(data *User) (int64, error) {
	return orm.NewOrm().Insert(data)
}

// UpdateUser 更新用户信息
func UpdateUser(filter, values map[string]interface{}) (int64, error) {
	needle := orm.NewOrm().QueryTable("home_user")
	for key, value := range filter {
		needle = needle.Filter(key, value)
	}

	return needle.Update(values)
}

// GetOneUser 获取某个用户的信息
func GetOneUser(filter map[string]interface{}) (User, error) {
	var user User

	needle := orm.NewOrm().QueryTable("home_user")
	for key, value := range filter {
		needle = needle.Filter(key, value)
	}

	err := needle.One(&user)
	return user, err
}

// GetUsers 获取多个用户信息
func GetUsers(filter map[string]interface{}) ([]*User, error) {
	var users []*User

	needle := orm.NewOrm().QueryTable("home_user")
	for key, value := range filter {
		needle = needle.Filter(key, value)
	}

	_, err := needle.All(&users)
	return users, err
}
