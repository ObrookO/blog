package models

import (
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id   int
	Name string
}

func init() {
	orm.RegisterModel(new(Category))
}

// 获取所有的栏目
// where map[string]interface{} 查询条件
func GetCategories(where map[string]interface{}) []Category {
	var categories []Category
	o := orm.NewOrm()
	needle := o.QueryTable("category")

	for key, value := range where {
		needle = needle.Filter(key, value)
	}

	needle.All(&categories)

	return categories
}
