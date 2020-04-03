package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Manager struct {
	Id        int
	Username  string
	Email     string
	Password  string
	Avatar    string
	CreatedAt time.Time `orm:"auto_now_add;type(timestamp)"`
	UpdatedAt time.Time `orm:"auto_now;type(timestamp)"`
}

func init() {
	orm.RegisterModelWithPrefix("admin_", new(Manager))
}
