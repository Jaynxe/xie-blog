package middleware

import (
	"context"
	"strings"

	"github.com/Jaynxe/xie-blog/model"
	"github.com/Jaynxe/xie-blog/utils/errhandle"
	"github.com/Jaynxe/xie-blog/utils/token"
	"github.com/gin-gonic/gin"
)

func UseTokenVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")
		prefix := "Bearer "
		tk := ""

		if auth != "" && strings.HasPrefix(auth, prefix) {
			tk = auth[len(prefix):]
		}

		if tk == "" {
			model.Throw(ctx, errhandle.TokenError)
			return
		}

		userinfo, ok := token.TK.Verify(context.Background(), tk)
		if !ok {
			model.Throw(ctx, errhandle.PermissionDenied)
			return
		}
		path := strings.Split(ctx.Request.URL.Path, "/")
		if len(path) >= 4 {
			switch path[3] {
			case "admin":
				if userinfo.Role != "admin" {
					model.Throw(ctx, errhandle.PermissionDenied)
					return
				}
			}
		}
		ctx.Set("info", userinfo)
		ctx.Next()
	}
}
