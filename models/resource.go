package models

// GetAllResource 获取所有干货收藏
func GetAllResource(filter map[string]interface{}) ([]*Resource, error) {
	var list []*Resource

	_, err := concatFilter("resource", filter).All(&list)
	return list, err
}
