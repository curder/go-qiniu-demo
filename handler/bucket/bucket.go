package bucket

// 保存请求
type StoreBucketRequest struct {
	AccountID   uint64 `json:"account_id" form:"account_id" binding:"required"`
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description"`
}

// 更新
type UpdateBucketInfoRequest struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description"`
}

// 账户详情响应
type ShowBucketInfoResponse struct {
	ID          uint64 `json:"id"`
	AccountID   uint64 `json:"account_id" form:"account_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
