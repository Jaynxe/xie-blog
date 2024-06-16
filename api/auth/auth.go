package auth

import (
	"github.com/Jaynxe/xie-blog/routes"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Auth struct {
}

func New() routes.Routes {
	return &Auth{}
}

func (a *Auth) InitRoute(g *gin.RouterGroup) {
}

// InitGlobalRoute 全局api,什么角色都可以调用
func (a *Auth) InitGlobalRoute(g *gin.RouterGroup) {
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge:   3600,
		HttpOnly: true,
		Domain:   "127.0.0.1", // 要发送Cookie的后端域名
		Path:     "/api",      // 要发送Cookie的后端路径
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	})
	g.Use(sessions.Sessions("verificationCode", store))

	g.GET("/getAllArticles", a.GetAllArticles)
	g.GET("/getAllMenus", a.GetAllMenus)
	g.GET("/isValid", a.IsValidSession)

	g.POST("/login", a.UserLogin)
	g.POST("/refresh", a.UserLoginRefresh)
	g.POST("/register", a.UserRegister)
	g.POST("/loginWithEmail", a.LoginWithEmail)
	//g.POST("/loginWithQQ", a.LoginWithQQ)

	g.POST("/resetPassword", a.ResetPassword)

}
