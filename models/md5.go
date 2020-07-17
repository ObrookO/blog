package models

// IsMd5RecordExists 判断MD5记录是否存在
func IsMd5RecordExists(filter map[string]interface{}) bool {
	return concatFilter("md5", filter).Exist()
}

// GetOneMd5Record 获取一条MD5记录
func GetOneMd5Record(filter map[string]interface{}) (Md5, error) {
	var record Md5

	err := concatFilter("md5", filter).One(&record)
	return record, err
}

// AddMd5Record 插入一条MD5记录
func AddMd5Record(record Md5) (int64, error) {
	return o.Insert(&record)
}
