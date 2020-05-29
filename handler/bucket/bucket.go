package bucket

// 保存请求
type StoreBucketRequest struct {
	AccountID   uint64 `json:"account_id" form:"account_id" binding:"required"`
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description"`
}
