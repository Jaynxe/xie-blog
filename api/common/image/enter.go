package image

import (
	"github.com/Jaynxe/xie-blog/routes"
	"github.com/gin-gonic/gin"
)

type Image struct{}

func New() routes.Routes {
	return &Image{}
}
func (i *Image) InitRoute(g *gin.RouterGroup) {
	g.GET("/getAllImages", i.GetAllImages)
	g.GET("/paginatedImages", i.GetPaginatedImages)

	g.POST("/uploadImages", i.UploadFile)

	g.DELETE("/deleteImages", i.ImageDelete)

	g.PATCH("/updateImage", i.ImageUpdate)

}
