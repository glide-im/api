package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/glide-im/api/internal/auth"
	"github.com/glide-im/api/internal/dao/userdao"
	"net/http"
	"strings"
)

var authRouteGroup gin.IRoutes

const CtxKeyAuthInfo = "CTX_KEY_AUTH_INFO"

func useAuth() gin.IRoutes {
	if authRouteGroup == nil {
		authRouteGroup = g.Use(authMiddleware).Use(crosMiddleware())
	}
	return authRouteGroup
}

func authMiddleware(context *gin.Context) {
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

	authInfo.AppId = userdao.UserInfoDao.GetUserAppId(authInfo.Uid)
	if authInfo.AppId == 0 {
		authInfo.AppId = userdao.UserInfoDao.GetGuestUserAppId(authInfo.Uid)
	}
	hasUser, err := userdao.UserInfoDao.HasUser(authInfo.Uid, authInfo.AppId)
	if err != nil {
		context.Status(http.StatusUnauthorized)
		context.Abort()
		return
	}
	if hasUser == false {
		context.Status(http.StatusUnauthorized)
		context.Abort()
		return
	}

	context.Set(CtxKeyAuthInfo, authInfo)
	fmt.Println("(authInfo.Uid", authInfo)
	context.Next()
}
