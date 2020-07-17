package models

// IsFavorRecordExists 判断点赞记录是否存在
func IsFavorRecordExists(filter map[string]interface{}) bool {
	return concatFilter("favor_record", filter).Exist()
}

// AddFavorRecord 添加点赞记录
func AddFavorRecord(record FavorRecord) (int64, error) {
	return o.Insert(&record)
}
