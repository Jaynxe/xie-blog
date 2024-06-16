package site

import (
	"github.com/Jaynxe/xie-blog/routes"
	"github.com/gin-gonic/gin"
)

type Site struct {
}

func New() routes.Routes {
	return &Site{}
}

func (s *Site) InitRoute(g *gin.RouterGroup) {
	g.GET("/getSiteInfo", s.GetSiteInfo)
	g.PATCH("/updateSiteInfo", s.UpdateSiteInfo)
}
