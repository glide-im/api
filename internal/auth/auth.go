package auth

import (
	"github.com/glide-im/api/internal/api/comm"
	"github.com/glide-im/glide/pkg/auth"
	"github.com/glide-im/glide/pkg/auth/jwt_auth"
	"strconv"
	"time"
)

var jwtAuth *jwt_auth.JwtAuthorize

func ParseToken(token string) (*AuthInfo, error) {
	return parseJwt(token)
}

func Auth(from int64, device int64, t string) (*jwt_auth.Response, error) {

	jwtAuthInfo := &jwt_auth.JwtAuthInfo{
		UID:    strconv.FormatInt(from, 10),
		Device: strconv.FormatInt(device, 10),
	}
	result, err2 := jwtAuth.Auth(jwtAuthInfo, &auth.Token{Token: t})

	resp, ok := result.Response.(*jwt_auth.Response)
	if !ok {
		return nil, err2
	}
	return resp, err2
}

func GenerateTokenExpire(uid int64, device int64, expireHour int64) (string, error) {
	jt := AuthInfo{
		Uid:    uid,
		Device: device,
		Ver:    genJwtVersion(),
	}
	expir := time.Now().Add(time.Hour * time.Duration(expireHour))
	token, err := genJwtExp(jt, expir)
	if err != nil {
		return "", comm.NewUnexpectedErr("login failed", err)
	}
	return token, nil
}
