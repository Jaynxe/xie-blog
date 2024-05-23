package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Jaynxe/xie-blog/app"
	"github.com/Jaynxe/xie-blog/core"
	"github.com/Jaynxe/xie-blog/flag"
	"github.com/Jaynxe/xie-blog/global"
	"github.com/Jaynxe/xie-blog/utils/token"
	"github.com/golang-jwt/jwt/v5"
)

func init() {
	// 初始化日志和配置信息
	core.InitConfig()
	core.InitLog()
	core.InitGorm()
	core.InitRedis()
	token.TK = token.NewJWTAccessGenerate(global.GVB_REDIS, jwt.SigningMethodHS256).(*token.JWTAccessGenerate)
}

// @title            Gin-Vue-Blog Swagger API接口文档
// @version          v1.0.0
// @description      使用gin+vue进行开发的博客平台
// @host			 127.0.0.1:8888
// @BasePath         /
func main() {
	//命令行参数绑定
	option := flag.Parse()
	//先停止web服务
	if flag.IsWebStop(option) {
		//根据命令行参数执行响应的方法
		flag.SwitchOption(option)
		return
	}
	server := app.NewServer()

	sigCh := make(chan os.Signal, 1)
	/*
		使用 signal.Notify 函数将 SIGINT 和 SIGTERM 信号注册到 sigCh 通道。
		这样，当程序收到这两个信号之一时，操作系统会将对应的信号发送到 sigCh 通道。
	*/
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	server.Close()

}
