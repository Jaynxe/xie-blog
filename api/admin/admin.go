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

func (a *Admin) Close() error {
	return nil
}

func (a *Admin) InitRoute(g *gin.RouterGroup) {

	admin := g.Group("/admin")

	admin.POST("/admin/new", a.RegisterAdmin)

	admin.GET("/paginatedUsers", a.GetPaginatedUsers)
	admin.GET("/getAllUsers", a.GetAllUsers)

	admin.PATCH("/modifyAdmin", a.ModifyAdmin)
	admin.PATCH("/modifyUser", a.ModifyUser)
	admin.PATCH("/modifyAdminPassword", a.ModifyAdminPassword)
	admin.PATCH("/modifyUserPassword", a.ModifyUserPassword)

	admin.DELETE("/deleteUser", a.DeleteUser)
	admin.DELETE("/deleteAdmin",a.DeleteAdmin)
}
