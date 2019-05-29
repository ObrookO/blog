package models

import (
	"github.com/astaxie/beego/orm"
)

type NuggetTags struct {
	Id    int
	Name  string
	OutId string
}

func init() {
	orm.RegisterModel(new(NuggetTags))
}

// 获取Nugeet的所有标签
func GetAllTags() []NuggetTags {
	var tags []NuggetTags
	o := orm.NewOrm()

	o.QueryTable("nugget_tags").All(&tags)
	return tags
}
