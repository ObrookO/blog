package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var o orm.Ormer

// Account 账号
type Account struct {
	Id           int
	Username     string    // 用户名
	Email        string    // 邮箱
	Password     string    // 密码
	Avatar       string    // 头像
	AllowComment int8      `orm:"default(1)"` // 是否允许评论
	Status       int8      `orm:"default(1)"` // 状态
	CreatedAt    time.Time `orm:"auto_now;type(timestamp)"`
	UpdatedAt    time.Time `orm:"auto_now_add;type(timestamp)"`
}

// Article 文章
type Article struct {
	Id           int
	Title        string         // 标题
	Keyword      string         // 关键词
	Category     *Category      `orm:"rel(one)"`
	Tags         []*Tag         `orm:"rel(m2m);rel_through(blog/models.ArticleTag)"` // 标签
	Description  string         // 描述
	Cover        string         // 封面地址
	CoverUrl     string         // 封面地址
	Content      string         // 内容
	Comments     []*Comment     `orm:"reverse(many)"` // 评论列表
	Favors       []*FavorRecord `orm:"reverse(many)"` // 点赞记录
	IsScroll     int8           // 是否轮播
	IsRecommend  int8           // 是否推荐
	AllowComment int8           // 是否允许评论
	Manager      *Manager       `orm:"rel(one)"` // 管理员信息，即作者信息
	Status       int8           // 状态
	CreatedAt    time.Time      // 添加时间
	UpdatedAt    time.Time      // 修改时间
}

// ArticleArchive 文章归档
type ArticleArchive struct {
	Date string
	Num  int64
}

// ArticleTag 文章标签
type ArticleTag struct {
	Id      int
	Article *Article `orm:"rel(one)"`
	Tag     *Tag     `orm:"rel(one)"`
}

// Category 栏目
type Category struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Comment 评论
type Comment struct {
	Id              int
	Account         *Account `orm:"rel(one)"`
	Article         *Article `orm:"rel(fk)"`
	OriginalContent string
	ShowContent     string
	Keyword         string // 违规关键词
	Ip              string
	CreatedAt       time.Time `orm:"auto_now_add;type(timestamp)"`
	UpdatedAt       time.Time `orm:"auto_now;type(timestamp)"`
}

// EmailLog 邮件日志
type EmailLog struct {
	Id        int
	EmailType int
	Address   string
	Content   string
	Result    string
	Reason    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// FavorRecord 点赞记录
type FavorRecord struct {
	Id        int
	Article   *Article `orm:"rel(fk)"`
	Ip        string
	CreatedAt time.Time `orm:"auto_now_add;type(timestamp)"`
	UpdatedAt time.Time `orm:"auto_now;type(timestamp)"`
}

// Manager 管理员
type Manager struct {
	Id        int
	Username  string
	Nickname  string
	Email     string
	Avatar    string
	Password  string
	IsAdmin   int8
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Md5 MD5记录
type Md5 struct {
	Id          int
	RawData     string
	EncryptData string
	CreatedAt   time.Time `orm:"auto_now_add;type(timestamp)"`
	UpdatedAt   time.Time `orm:"auto_now;type(timestamp)"`
}

// Resource 干货收藏
type Resource struct {
	Id          int
	Title       string
	Description string
	Url         string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Tag 标签
type Tag struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	if beego.BConfig.RunMode == "dev" {
		//orm.Debug = true
	}

	host := beego.AppConfig.String("db_host")
	db := beego.AppConfig.String("db_name")
	user := beego.AppConfig.String("db_user")
	pass := beego.AppConfig.String("db_password")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	// 注册默认数据库
	orm.RegisterDataBase("default", "mysql", user+":"+pass+"@tcp("+host+")/"+db+"?charset=utf8&loc=Asia%2FShanghai")

	orm.RegisterModel(
		new(Account),
		new(Article),
		new(ArticleTag),
		new(Category),
		new(Comment),
		new(EmailLog),
		new(FavorRecord),
		new(Manager),
		new(Md5),
		new(Resource),
		new(Tag),
	)

	o = orm.NewOrm()
}
