package models

// AddComment 添加评论
func AddComment(comment Comment) (int64, error) {
	return o.Insert(&comment)
}
