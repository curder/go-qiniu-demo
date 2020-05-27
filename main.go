package main

import (
	"github.com/curder/go-qiniu-demo/config"
	v "github.com/curder/go-qiniu-demo/pkg/version"
	"github.com/spf13/pflag"
)

var (
	cfg     = pflag.StringP("config", "c", "", "snake config file path.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	var (
		err   error
	)
	// 加载版本信息
	pflag.Parse()
	v.Init(*version)

	// 加载配置
	if err = config.Init(*cfg); err != nil {
		panic(err)
	}

	// 加载数据库

	// 加载Redis缓存

	// 加载Gin框架

	//

}
