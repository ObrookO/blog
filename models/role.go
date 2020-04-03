package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Role struct {
	Id        int
	Name      string
	CreatedAt time.Time `orm:"auto_now_add;type(timestamp)"`
	UpdatedAt time.Time `orm:"auto_now;type(timestamp)"`
}

func init() {
	orm.RegisterModelWithPrefix("admin_", new(Role))
}
