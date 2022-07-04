package api

import (
	"github.com/gin-gonic/gin"
	"github.com/glide-im/api/internal/auth"
	"net/http"
	"strings"
)

const CtxGuestKeyAuthInfo = "CTX_KEY_GUEST_INFO"

func useGuestAuth() gin.IRoutes {
	if authRouteGroup == nil {
		authRouteGroup = g.Use(guestMiddleware).Use(crosMiddleware())
	}
	return authRouteGroup
}

func guestMiddleware(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	if authHeader == "" {
		context.Status(http.StatusUnauthorized)
		context.Abort()
		return
	}
	authHeader = strings.TrimPrefix(authHeader, "Bearer ")
	authInfo, err := auth.ParseToken(authHeader)
	if err != nil {
		context.Status(http.StatusUnauthorized)
		context.Abort()
		return
	}
	context.Set(CtxGuestKeyAuthInfo, authInfo)
	// 这里通过检测 设置他的 app_id

	context.Next()
}
