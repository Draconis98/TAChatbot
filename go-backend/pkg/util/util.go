package util

type HeadInfo struct {
	StatusCode int    `json:"status_code"`
	Username   string `header:"username" json:"username" binding:"required"`
	Email      string `header:"email" json:"email" binding:"required"`
}
