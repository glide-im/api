package auth

import (
	"errors"
	"fmt"
	"github.com/glide-im/glide/pkg/auth"
	"github.com/glide-im/glide/pkg/auth/jwt_auth"
	"strconv"
)

var jwtAuth *jwt_auth.JwtAuthorize

var (
	GUEST_DEVICE   = int64(3)
	MOBILE_DEVICE  = int64(2)
	DEFAULT_DEVICE = int64(1)
)

func ParseToken(token string) (*AuthInfo, error) {
	var a = &jwt_auth.JwtAuthInfo{}
	result, err := jwtAuth.Auth(a, &auth.Token{Token: token})
	if err != nil {
		return nil, err
	}
	resp, ok := result.Response.(*jwt_auth.Response)
	if !ok {
		return nil, errors.New("invalid auth info")
	}

	fmt.Println("resp", resp)
	parseInt, _ := strconv.ParseInt(resp.Uid, 10, 64)
	i, _ := strconv.ParseInt(resp.Device, 10, 64)
	return &AuthInfo{
		Uid:    parseInt,
		Device: i,
	}, nil
}

func Auth(from int64, device int64, t string) (*jwt_auth.Response, error) {

	jwtAuthInfo := &jwt_auth.JwtAuthInfo{
		UID:    strconv.FormatInt(from, 10),
		Device: strconv.FormatInt(device, 10),
	}
	result, err2 := jwtAuth.Auth(jwtAuthInfo, &auth.Token{Token: t})

	if err2 != nil {
		return nil, err2
	}
	resp, ok := result.Response.(*jwt_auth.Response)
	if !ok {
		return nil, err2
	}
	return resp, err2
}

func GenerateTokenExpire(uid int64, device int64, expireHour int64) (string, error) {
	jai := &jwt_auth.JwtAuthInfo{
		UID:    strconv.FormatInt(uid, 10),
		Device: strconv.FormatInt(device, 10),
	}
	fmt.Println("strconv.FormatInt(device, 10)", strconv.FormatInt(device, 10))
	token, err := jwtAuth.GetToken(jai)
	return token.Token, err
}
