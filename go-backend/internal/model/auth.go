package model

// UserInfo 结构体用于解析GitHub API响应
type UserInfo struct {
	Login string `json:"login"` // GitHub用户名
	Email string `json:"email"` // GitHub邮箱
	// ... 其他可能的字段 ...
}
