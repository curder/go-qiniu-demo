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

// 账户详情响应
type ShowAccountResponse struct {
	ID          uint64 `json:"id"`
	AccountKey  string `json:"access_key"`
	SecretKey   string `json:"secret_key"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
