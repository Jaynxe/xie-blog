package tag
import (
	"github.com/Jaynxe/xie-blog/routes"
	"github.com/gin-gonic/gin"
)

type Tag struct {
}

func New() routes.Routes {
	return &Tag{}
}
func (t *Tag) InitRoute(g *gin.RouterGroup) {
	g.POST("/addTag", t.TagCreate)

	g.PATCH("/updateTag/:id", t.TagUpdate)

	g.GET("/getTag/:id", t.GetTag)
	g.GET("/getAllTags", t.GetAllTags)
	g.GET("/paginatedTags", t.GetPaginatedTags)

	g.DELETE("/deleteTag", t.TagDelete)

}
