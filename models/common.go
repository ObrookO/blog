package models

import "github.com/astaxie/beego/orm"

// concatFilter 拼接查询条件
func concatFilter(table string, filter map[string]interface{}) orm.QuerySeter {
	needle := o.QueryTable(table)
	for k, v := range filter {
		needle = needle.Filter(k, v)
	}

	return needle
}
