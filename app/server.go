package app

import (
	"context"
	"net/http"
	"time"

	"github.com/Jaynxe/xie-blog/api/admin"
	"github.com/Jaynxe/xie-blog/api/auth"
	"github.com/Jaynxe/xie-blog/api/common/image"
	"github.com/Jaynxe/xie-blog/api/common/menu"
	"github.com/Jaynxe/xie-blog/api/common/site"
	"github.com/Jaynxe/xie-blog/api/common/tag"
	"github.com/Jaynxe/xie-blog/api/user"
	"github.com/Jaynxe/xie-blog/routes"

	"github.com/Jaynxe/xie-blog/global"
)

type Server struct {
	srv    *http.Server
	routes []routes.Routes
}

func initModules() []routes.Routes {
	return []routes.Routes{
		auth.New(),
		admin.New(),
		image.New(),
		menu.New(),
		user.New(),
		tag.New(),
		site.New(),
	}
}

// NewServer 创建服务
func NewServer() *Server {
	s := &Server{
		routes: initModules(),
	}
	s.dispatchRoute()
	return s
}

// 优雅关闭服务器
func (s *Server) shutdownHTTPServer() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		global.GVB_LOGGER.Fatal(err)
	}
}

func (s *Server) Close() {
	s.shutdownHTTPServer()
}
