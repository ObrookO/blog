package models

import (
	"github.com/astaxie/beego/orm"
)

type FavorRecord struct {
	Id  int
	Aid string
	Ip  string
}

func init() {
	orm.RegisterModel(new(FavorRecord))
}

// 获取所有的点赞记录
func GetAllFavorRecord() ([]FavorRecord, error) {
	var records []FavorRecord
	o := orm.NewOrm()

	_, err := o.QueryTable("favor_record").All(&records)
	return records, err
}
