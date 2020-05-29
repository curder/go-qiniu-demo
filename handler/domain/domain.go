package domain

// 保存请求
type StoreDomainRequest struct {
	BucketID    uint64 `json:"bucket_id" form:"bucket_id" binding:"required"`
	Protocol    string `json:"protocol" form:"protocol" binding:"required"`
	Hostname    string `json:"hostname" form:"hostname" binding:"required"`
	Description string `json:"description" form:"description"`
}
