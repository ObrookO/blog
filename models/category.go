package models

// GetAllCategories 获取所有栏目
func GetAllCategories(filter map[string]interface{}, field ...string) ([]*Category, error) {
	var categories []*Category

	_, err := concatFilter("category", filter).All(&categories, field...)
	return categories, err
}
