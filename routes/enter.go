package routes

import "github.com/gin-gonic/gin"

type GlobalRoutes interface {
	InitGlobalRoute(*gin.RouterGroup)
}


type Routes interface {
	InitRoute(*gin.RouterGroup)
}
