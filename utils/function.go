package utils

import (
	"crypto/md5"
	"fmt"
)

// 使用md5加密字符串
func Md5Str(str string) string {
	md5Str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5Str
}
