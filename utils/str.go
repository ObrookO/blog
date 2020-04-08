package utils

// ObjInIntSlice 判断obj是否在slice中
func ObjInIntSlice(obj int, s []int) bool {
	for _, item := range s {
		if obj == item {
			return true
		}
	}

	return false
}
