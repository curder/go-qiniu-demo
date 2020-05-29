package domain

// 保存请求
type StoreDomainRequest struct {
	BucketID    uint64 `json:"bucket_id" form:"bucket_id" binding:"required"`
	Protocol    string `json:"protocol" form:"protocol" binding:"required"`
	Hostname    string `json:"hostname" form:"hostname" binding:"required"`
	Description string `json:"description" form:"description"`
}

type ShowDomainInfoResponse struct {
	ID          uint64 `json:"id"`
	BucketID    uint64 `json:"bucket_id"`
	Protocol    string `json:"protocol"`
	Hostname    string `json:"hostname"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// 更新域名请求
type UpdateDomainRequest struct {
	BucketID    uint64 `form:"bucket_id" binding:"required"`
	Protocol    string `form:"protocol" binding:"required"`
	Hostname    string `form:"hostname" binding:"required"`
	Description string `form:"description"`
}
