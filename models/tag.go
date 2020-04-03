package models

import (
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Tag struct {
	Id         int
	Name       string
	ManagerId  int
	ArticleNum int
	CreatedAt  time.Time `orm:"auto_now_add;type(timestamp)"`
	UpdatedAt  time.Time `orm:"auto_now;type(timestamp)"`
}

func init() {
	orm.RegisterModelWithPrefix("admin_", new(Tag))
}

// 获取所有标签
// where map[string]interface{} 查询条件
func GetTags(where map[string]interface{}) []Tag {
	var tags []Tag
	needle := orm.NewOrm().QueryTable("admin_tag")

	for key, value := range where {
		needle = needle.Filter(key, value)
	}

	needle.All(&tags)

	return tags
}

// 获取标签的某一列，返回列表
// where map[string]interface{} 查询条件
// col ... string 查询的列
func GetTagFields(where map[string]interface{}, col ...string) []string {
	var tags orm.ParamsList
	var list []string
	needle := orm.NewOrm().QueryTable("admin_tag")

	for key, value := range where {
		needle = needle.Filter(key, value)
	}

	needle.ValuesFlat(&tags, strings.Join(col, ","))
	for _, ele := range tags {
		list = append(list, ele.(string))
	}

	return list
}
