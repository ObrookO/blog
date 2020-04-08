package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Log struct {
	Id        int
	UserId    int
	Ip        string
	Url       string
	Method    string
	Query     string
	Headers   string
	Body      string
	Response  string
	Content   string
	Reason    string
	CreatedAt time.Time `orm:"auto_now_add;type(timestamp)"`
	UpdatedAt time.Time `orm:"auto_now;type(timestamp)"`
}

func init() {
	orm.RegisterModelWithPrefix("home_", new(Log))
}

// 记录日志
func AddHomeLog(data Log) (int64, error) {
	return orm.NewOrm().Insert(&data)
}
