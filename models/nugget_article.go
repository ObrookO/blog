package models

import (
	"github.com/astaxie/beego/orm"
)

type NuggetArticles struct {
	Id       int
	Title    string
	Category string
	Tags     string
	Content  string
	CatchAt  string
}

func init() {
	orm.RegisterModel(new(NuggetArticles))
}

// 获取Nugget的所有文章
func GetAllArticles(where map[string]interface{}, offest int, limit int) []NuggetArticles {
	var articles []NuggetArticles
	o := orm.NewOrm()

	needle := o.QueryTable("nugget_articles")
	for key, value := range where {
		needle = needle.Filter(key, value)
	}

	needle.Offset(offest).Limit(limit).OrderBy("id").All(&articles, "id", "title", "catch_at")
	return articles
}

// 获取Nugget文章总数
func GetNuggetTotal(where map[string]interface{}) (int64, error) {
	o := orm.NewOrm()
	needle := o.QueryTable("nugget_articles")

	for key, value := range where {
		needle = needle.Filter(key, value)
	}

	return needle.Count()
}

// 获取Nugget文章详情
func GetArticleInfo(where map[string]interface{}) NuggetArticles {
	var article NuggetArticles
	o := orm.NewOrm()
	needle := o.QueryTable("nugget_articles")

	for key, value := range where {
		needle = needle.Filter(key, value)
	}

	needle.One(&article)
	return article
}

// 获取上一篇文章id以及下一篇文章的id
func GetLastAndNext(aid int) (int, int) {
	var article NuggetArticles
	var last, next int

	o := orm.NewOrm()
	// 获取上一篇的文章id
	if err := o.QueryTable("nugget_articles").Filter("id__lt", aid).OrderBy("-id").One(&article, "id"); err == nil {
		last = article.Id
	} else {
		last = 0
	}

	if err := o.QueryTable("nugget_articles").Filter("id__gt", aid).OrderBy("id").One(&article, "id"); err == nil {
		next = article.Id
	} else {
		next = 0
	}

	return last, next
}
