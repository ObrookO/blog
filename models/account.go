package models

// IsAccountExist 判断账号是否存在
func IsAccountExist(filter map[string]interface{}) bool {
	return concatFilter("account", filter).Exist()
}

// GetOneAccount 获取账号信息
func GetOneAccount(filter map[string]interface{}, field ...string) (Account, error) {
	var account Account

	err := concatFilter("account", filter).RelatedSel().One(&account, field...)
	return account, err
}

// AddAccount 添加账号
func AddAccount(account Account) (int64, error) {
	return o.Insert(&account)
}

// UpdateAccount 修改账号信息
func UpdateAccount(filter, value map[string]interface{}) (int64, error) {
	return concatFilter("account", filter).Update(value)
}
