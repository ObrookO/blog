package tool

import "strings"

// GetCharAmount 获取字符数
func GetCharAmount(str string) int {
	return len([]rune(strings.TrimSpace(str)))
}
