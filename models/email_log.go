package models

const (
	NewManager           = iota // 添加管理员成功
	ResetManagerPassword        // 管理员重置密码
	RegisterAccount             // 前台用户注册
	NewAccount                  // 用户注册成功
	ResetAccountPassword        // 前台用户重置密码
)

// AddEmailLog 添加发送邮件日志
func AddEmailLog(log EmailLog) (int64, error) {
	return o.Insert(&log)
}
