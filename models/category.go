package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id         int
	Name       string
	ShortName  string
	ArticleNum int
	ManagerId  int
	CreatedAt  time.Time `orm:"auto_now_add;type(timestamp)"`
	UpdatedAt  time.Time `orm:"auto_now;type(timestamp)"`
}

func init() {
	orm.RegisterModelWithPrefix("admin_", new(Category))
}

// IsCategoryExists 判断栏目是否存在
func IsCategoryExists(filter map[string]interface{}) bool {
	needle := orm.NewOrm().QueryTable("admin_category")
	for key, value := range filter {
		needle = needle.Filter(key, value)
	}

	return needle.Exist()
}

// AddCategory 添加栏目
func AddCategory(data *Category) (int64, error) {
	return orm.NewOrm().Insert(data)
}

// GetCategories 获取所有的栏目
func GetCategories(filter map[string]interface{}) ([]*Category, error) {
	var categories []*Category

	needle := orm.NewOrm().QueryTable("admin_category")
	for key, value := range filter {
		needle = needle.Filter(key, value)
	}

	_, err := needle.All(&categories)

	return categories, err
}

// GetCategory 获取某个栏目
func GetCategory(filter map[string]interface{}) (Category, error) {
	var category Category

	needle := orm.NewOrm().QueryTable("admin_category")
	for key, value := range filter {
		needle = needle.Filter(key, value)
	}

	err := needle.One(&category)

	return category, err
}

// DeleteCategory 删除栏目
func DeleteCategory(filter map[string]interface{}) (int64, error) {
	needle := orm.NewOrm().QueryTable("admin_category")
	for key, value := range filter {
		needle = needle.Filter(key, value)
	}

	return needle.Delete()
}
