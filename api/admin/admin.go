package admin

import (
	"io"

	"github.com/Jaynxe/xie-blog/routes"
	"github.com/gin-gonic/gin"
)

type Admin struct {
}

var _ io.Closer = &Admin{}

func New() routes.Routes {
	return &Admin{}
}

func (ad *Admin) Close() error {
	return nil
}

func (ad *Admin) InitRoute(g *gin.RouterGroup) {

	admin := g.Group("/admin")

	admin.POST("/admin/new", ad.RegisterAdmin)

	admin.GET("/paginatedUsers", ad.GetPaginatedUsers)
	admin.GET("/getAllUsers", ad.GetAllUsers)

	admin.PATCH("/modifyAdmin", ad.ModifyAdmin)
	admin.PATCH("/modifyUser", ad.ModifyUser)
	admin.PATCH("/modifyAdminPassword", ad.ModifyAdminPassword)
	admin.PATCH("/modifyUserPassword", ad.ModifyUserPassword)

	admin.DELETE("/deleteUser", ad.DeleteUser)
	admin.DELETE("/deleteAdmin", ad.DeleteAdmin)
}
