package main

import (
	"context"
	"github.com/curder/go-qiniu-demo/config"
	"github.com/curder/go-qiniu-demo/model"
	"github.com/curder/go-qiniu-demo/pkg/log"
	"github.com/curder/go-qiniu-demo/pkg/redis"
	v "github.com/curder/go-qiniu-demo/pkg/version"
	"github.com/curder/go-qiniu-demo/routers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	cfg     = pflag.StringP("config", "c", "", "snake config file path.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	var (
		router *gin.Engine
		srv    *http.Server
		err    error
	)

	// 加载版本信息
	pflag.Parse()
	v.Init(*version)

	// 加载配置
	if err = config.Init(*cfg); err != nil {
		panic(err)
	}

	// 加载数据库
	model.DB.Init()
	defer model.DB.Close()

	// 加载Redis缓存
	redis.Init()

	// 加载Gin框架
	// 设置Gin模式
	gin.SetMode(viper.GetString("run_mode"))
	router = gin.Default()
	routers.Load(router)

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	srv = &http.Server{
		Addr:    viper.GetString("addr"),
		Handler: router,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s", err.Error())
		}
	}()

	gracefulStop(srv)
}

// gracefulStop 优雅退出
// 等待中断信号以超时 5 秒正常关闭服务器
// 官方说明：https://github.com/gin-gonic/gin#graceful-restart-or-stop
func gracefulStop(srv *http.Server) {
	quit := make(chan os.Signal)
	// kill 命令发送信号 syscall.SIGTERM
	// kill -2 命令发送信号 syscall.SIGINT
	// kill -9 命令发送信号 syscall.SIGKILL
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// 5 秒后捕获 ctx.Done() 信号
	select {
	case <-ctx.Done():
		log.Info("timeout of 5 seconds.")
	default:
	}
	log.Info("Server exiting")
}
