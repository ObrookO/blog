package tool

import "regexp"

// IsValidEmail 判断email是否合法
func IsValidEmail(email string) bool {
	pattern := "^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$"
	reg, _ := regexp.Compile(pattern)

	return reg.MatchString(email)
}

// IsValidPassword 判断密码是否合法
func IsValidPassword(password string) bool {
	pattern := "^[0-9a-zA-Z_]{8,16}$"
	reg, _ := regexp.Compile(pattern)

	return reg.MatchString(password)
}
