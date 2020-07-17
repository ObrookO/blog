package models

// GetAllTags 获取所有标签
func GetAllTags(filter map[string]interface{}, field ...string) ([]*Tag, error) {
	var tags []*Tag

	_, err := concatFilter("tag", filter).All(&tags, field...)
	return tags, err
}
