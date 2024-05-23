package user

import (
	"github.com/Jaynxe/xie-blog/routes"
	"github.com/gin-gonic/gin"
)

type User struct{}

func New() routes.Routes {
	return &User{}
}
func (u *User) InitRoute(g *gin.RouterGroup) {
	user := g.Group("/user")
	user.GET("/getUserInfo", u.GetUserInfo)

	user.POST("/logout", u.Logout)

	user.PATCH("/modifyUserPassword", u.ModifyUserPassword)
	user.PATCH("/modifyUser", u.ModifyUser)

	user.DELETE("/deleteUser", u.DeleteUser)

}
