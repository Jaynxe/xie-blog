package menu
import (
	"github.com/Jaynxe/xie-blog/routes"
	"github.com/gin-gonic/gin"
)

type Menu struct {
}

func New() routes.Routes {
	return &Menu{}
}
func (m *Menu) InitRoute(g *gin.RouterGroup) {
	g.POST("/addMenu", m.MenuCreate)
	g.GET("/getMenu/:id", m.GetMenu)
	g.PATCH("/updateMenu/:id", m.MenuUpdate)
	g.DELETE("/deleteMenu", m.MenuDelete)
}
