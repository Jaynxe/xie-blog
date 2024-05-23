package app

import (
	"net/http"

	"github.com/Jaynxe/xie-blog/global"
	docs "github.com/Jaynxe/xie-blog/docs"
	"github.com/Jaynxe/xie-blog/middleware"
	"github.com/Jaynxe/xie-blog/routes"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) initAuthRequiedRoutes(g *gin.RouterGroup) {
	for _, r := range s.routes {
		r.InitRoute(g)
	}
}

func (s *Server) initGlobalRoutes(g *gin.RouterGroup) {
	for _, r := range s.routes {
		if route, ok := r.(routes.GlobalRoutes); ok {
			route.InitGlobalRoute(g)
		}
	}
}

// 调度路由
func (s *Server) dispatchRoute() {
	if global.GVB_CONFIG.System.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	docs.SwaggerInfo.BasePath = "/api"
	e := gin.Default()
	e.Use(middleware.UseCORS())
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := e.Group("/api")
	// 初始化全局api
	s.initGlobalRoutes(api)

	requiredAuth := api.Group("/authrequired")
	requiredAuth.Use(middleware.UseTokenVerify())
	// 初始化需要授权的api
	s.initAuthRequiedRoutes(requiredAuth)

	// 启动服务器
	s.setupHTTPServer(e)

}

// 开启服务器
func (s *Server) setupHTTPServer(e *gin.Engine) {
	s.srv = &http.Server{
		Addr:    global.GVB_CONFIG.System.Addr(),
		Handler: e,
	}

	go s.srv.ListenAndServe()

	global.GVB_LOGGER.Infof("HTTP Server Starts At %s", s.srv.Addr)

}
