package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id        string `orm:"pk"`
	Title     string
	Cate      string
	Tags      string
	Desc      string
	Content   string
	FavorNum  int64
	IsFavored bool `orm:"-"`
	Status    string
	CreatedAt string
}

type ArticleArchive struct {
	Date  string
	Value string
	Sum   int
}

func init() {
	orm.RegisterModel(new(Article))
}

// 获取所有的文章
// where map[string]interface{} 查询条件
// offset int 偏移量
// limit int 取的条数
func GetArticlesLimit(where map[string]interface{}, offset int, limit int) []Article {
	var articles []Article
	o := orm.NewOrm()
	needle := o.QueryTable("article")
	for key, value := range where {
		needle = needle.Filter(key, value)
	}

	needle.Offset(offset).Limit(limit).OrderBy("-id").All(&articles, "id", "title", "desc", "created_at")
	return articles
}

// 判断文章是否存在
// where map[string]interface{} 查询条件
func IsArticleExists(where map[string]interface{}) bool {
	o := orm.NewOrm()
	needle := o.QueryTable("article")

	for key, value := range where {
		needle = needle.Filter(key, value)
	}

	return needle.Exist()
}

// 获取文章总数
// map[string]interface{} 查询条件
func GetTotal(where map[string]interface{}) (int64, error) {
	o := orm.NewOrm()
	needle := o.QueryTable("article")

	for key, value := range where {
		needle = needle.Filter(key, value)
	}

	return needle.Count()
}

// 文章归档
func Archive() []ArticleArchive {
	var archive []ArticleArchive
	o := orm.NewOrm()

	o.Raw("select date_format(created_at, '%Y年%m月') as date, date_format(created_at, '%Y/%m') as value, count(*) as sum from article group by value").QueryRows(&archive)
	return archive
}

// 使用find_in_set查询文章
func GetArticlesByTag(col string, value string, off int, lim int) []Article {
	var articles []Article
	var sql string
	o := orm.NewOrm()

	if lim > 0 && off >= 0 {
		sql = fmt.Sprintf("select id, title, `desc`, created_at from article where find_in_set('%s', `%s`) order by id desc limit %d offset %d", value, col, off, lim)

	} else {
		sql = fmt.Sprintf("select id, title, `desc`, created_at from article where find_in_set('%s', `%s`) order by id desc", value, col)
	}

	o.Raw(sql).QueryRows(&articles)

	return articles
}

// 获取一篇文章
// where map[string]interface{} 查询条件
func GetOneArticle(where map[string]interface{}) (Article, error) {
	var article Article
	o := orm.NewOrm()
	needle := o.QueryTable("article")

	for key, value := range where {
		needle = needle.Filter(key, value)
	}

	err := needle.One(&article)

	return article, err
}

// 获取上一篇文章的id以及下一篇文章的id
// aid int 当前文章的id
func GetBeforeAndAfter(aid string) (int, int) {
	var article Article
	var before, after int

	o := orm.NewOrm()
	// 获取上一篇文章的id
	if err := o.QueryTable("article").Filter("id__lt", aid).OrderBy("-id").One(&article, "id"); err == nil {
		before, _ = strconv.Atoi(article.Id)
	} else {
		before = 0
	}

	// 获取下一篇文章的id
	if err := o.QueryTable("article").Filter("id__gt", aid).OrderBy("id").One(&article, "id"); err == nil {
		after, _ = strconv.Atoi(article.Id)
	} else {
		after = 0
	}

	return before, after
}
