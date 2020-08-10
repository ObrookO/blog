package models

import (
	"github.com/astaxie/beego/orm"
)

// IsArticleExists 判断文章是否存在
func IsArticleExists(filter map[string]interface{}) bool {
	return concatFilter("article", filter).Exist()
}

// GetRecommendArticles 获取推荐的文章
func GetRecommendArticles(limit int) ([]*Article, error) {
	var articles []*Article

	field := []string{"id", "title"}
	_, err := concatFilter("article", map[string]interface{}{"is_recommend": 1, "status": 1}).OrderBy("-id").Limit(limit).All(&articles, field...)
	return articles, err
}

// GetArticleArchive 获取文章归档以及数量
func GetArticleArchive() ([]*ArticleArchive, error) {
	var archive []*ArticleArchive

	_, err := o.Raw("select date_format(created_at, '%Y-%m') as date, " +
		"count(1) as num from article where status = 1 group by date order by date desc").QueryRows(
		&archive)

	return archive, err
}

// GetAllArticles 获取所有文章
func GetAllArticles(filter map[string]interface{}, offset, limit int, field ...string) ([]*Article, error) {
	var articles []*Article

	_, err := concatFilter("article", filter).
		RelatedSel().
		OrderBy("-id").
		Offset(offset).
		Limit(limit).
		All(&articles, field...)

	for _, a := range articles {
		o.LoadRelated(a, "Tags")
		o.LoadRelated(a, "Comments")
		o.LoadRelated(a, "Favors")
	}

	return articles, err
}

// GetAllArticlesOfTag 获取标签下的所有文章
func GetAllArticlesOfTag(filter map[string]interface{}, offset, limit int, field ...string) ([]*Article, error) {
	var articles []*Article
	var articleIds orm.ParamsList
	var err error
	concatFilter("article_tag", map[string]interface{}{"tag_id": filter["tag_id"]}).ValuesFlat(&articleIds, "article_id")

	if len(articleIds) > 0 {
		_, err = concatFilter("article", map[string]interface{}{"id__in": articleIds, "status": 1}).
			RelatedSel().
			OrderBy("-id").
			Offset(offset).
			Limit(limit).
			All(&articles, field...)

		for _, a := range articles {
			o.LoadRelated(a, "Tags")
			o.LoadRelated(a, "Comments")
			o.LoadRelated(a, "Favors")
		}
	}

	return articles, err
}

// GetAllArticlesOfCategory 获取栏目下的所有文章
func GetAllArticlesOfCategory(filter map[string]interface{}, offset, limit int, field ...string) ([]*Article, error) {
	var articles []*Article

	_, err := concatFilter("article", filter).
		RelatedSel().
		OrderBy("-id").
		Offset(offset).
		Limit(limit).
		All(&articles, field...)

	for _, a := range articles {
		o.LoadRelated(a, "Tags")
		o.LoadRelated(a, "Comments")
		o.LoadRelated(a, "Favors")
	}

	return articles, err
}

// GetAllArticlesOfManager 获取作者的所有文章
func GetAllArticlesOfManager(filter map[string]interface{}, offset, limit int, field ...string) ([]*Article, error) {
	var articles []*Article

	_, err := concatFilter("article", filter).
		RelatedSel().
		OrderBy("-id").
		Offset(offset).
		Limit(limit).
		All(&articles, field...)

	for _, a := range articles {
		o.LoadRelated(a, "Tags")
		o.LoadRelated(a, "Comments")
		o.LoadRelated(a, "Favors")
	}

	return articles, err
}

// GetAllArticlesOfDate 获取日期下的所有文章
func GetAllArticlesOfDate(filter map[string]interface{}, offset, limit int, field ...string) ([]*Article, error) {
	var articles []*Article

	_, err := concatFilter("article", filter).
		RelatedSel().
		OrderBy("-id").
		Offset(offset).
		Limit(limit).
		All(&articles, field...)

	for _, a := range articles {
		o.LoadRelated(a, "Tags")
		o.LoadRelated(a, "Comments")
		o.LoadRelated(a, "Favors")
	}

	return articles, err
}

// GetOneArticle 获取文章详情
func GetOneArticle(filter map[string]interface{}, field ...string) (Article, error) {
	var article Article

	err := concatFilter("article", filter).RelatedSel().One(&article, field...)

	o.LoadRelated(&article, "Tags")
	o.LoadRelated(&article, "Favors")
	o.LoadRelated(&article, "Comments", true, 1000, 0, "-id")

	return article, err
}
