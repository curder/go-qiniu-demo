package account

// 保存账号
type StoreAccountRequest struct {
	Email       string `json:"email" form:"email" binding:"required" example:"example@example.com"`
	AccessKey   string `json:"access_key" form:"access_key" example:"xxxxxx"`
	SecretKey   string `json:"secret_key" form:"secret_key" binding:"required" example:"xxxxxx"`
	Description string `json:"description" form:"description"`
}

// 更新账号
type UpdateAccountRequest struct {
	Email       string `json:"email" form:"email" binding:"required" example:"example@example.com"`
	AccessKey   string `json:"access_key" form:"access_key" binding:"required" example:"xxxxxx"`
	SecretKey   string `json:"secret_key" form:"secret_key" binding:"required" example:"xxxxxx"`
	Description string `json:"description" form:"description"`
}
