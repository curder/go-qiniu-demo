package users

// CreateRequest 创建用户请求
type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserLoginCredentials 用户登录
type UserLoginCredentials struct {
	Name     string `json:"name" form:"name" binding:"required" example:"john/john@example.com"`
	Password string `json:"password" form:"password" binding:"required" example:"your own password"`
}

// UserRegisterCredentials用户注册
type UserRegisterCredentials struct {
	Name     string `json:"name" form:"name" binding:"required" example:"john"`
	Email    string `json:"email" form:"email" binding:"required" example:"john@example.com"`
	Password string `json:"password" form:"password" binding:"required" example:"your own password"`
}
