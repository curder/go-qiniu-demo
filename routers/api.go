package routers

import (
	"github.com/curder/go-qiniu-demo/handler"
	"github.com/curder/go-qiniu-demo/handler/account"
	"github.com/curder/go-qiniu-demo/handler/users"
	"github.com/curder/go-qiniu-demo/routers/middleware"
	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	var (
		u *gin.RouterGroup
	)
	// 使用中间件
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(middleware.Logging())
	g.Use(middleware.RequestID())
	g.Use(mw...)

	// 404 Handler.
	g.NoRoute(handler.RouteNotFound)
	g.NoMethod(handler.RouteNotFound)

	// The user handlers, requiring authentication
	u = g.Group("/v1/auth")
	u.POST("/login", users.Login)
	u.POST("/register", users.Register)
	u.Use(middleware.AuthMiddleware())
	{
		u.GET("/users", users.Info)
	}

	u = g.Group("/v1/accounts")
	u.Use(middleware.AuthMiddleware())
	{
		u.GET("", account.Index)
		u.POST("", account.Store)
		u.GET("/:id", account.Show)
		u.PUT("/:id", account.Update)
		u.DELETE("/:id", account.Delete)
		u.DELETE("/:id/force", account.ForceDelete)
		u.PUT("/:id/restore", account.Restore)
	}

	return g
}
