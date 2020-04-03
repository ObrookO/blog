package models

import "github.com/astaxie/beego/orm"

type Permission struct {
	Id int
}

func init() {
	orm.RegisterModelWithPrefix("admin_", new(Permission))
}
