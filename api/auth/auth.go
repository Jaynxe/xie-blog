package auth

import (
	"github.com/Jaynxe/xie-blog/routes"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Auth struct {
}

func New() routes.Routes {
	return &Auth{}
}

func (a *Auth) InitRoute(g *gin.RouterGroup) {
}

// 全局api,什么角色都可以调用
func (a *Auth) InitGlobalRoute(g *gin.RouterGroup) {
	store := cookie.NewStore([]byte("secret"))
	g.Use(sessions.Sessions("mysession", store))
	
	g.GET("/isvalid", a.IsValidSession)
	g.POST("/login", a.UserLogin)
	g.POST("/refresh", a.UserLoginRefresh)
	g.POST("/register", a.UserRegister)
	g.GET("/getallArticles", a.GetAllArticles)
	g.GET("/getAllMenus", a.GetAllMenus)
	g.POST("/loginWithEmail", a.LoginWithEmail)
	g.POST("/loginWithQQ",a.LoginWithQQ)
}
