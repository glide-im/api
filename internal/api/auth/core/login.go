package core

import (
	comm2 "github.com/glide-im/api/internal/api/comm"
	route "github.com/glide-im/api/internal/api/router"
)

var (
	VerifyCodeError = comm2.NewApiBizError(2001, "验证码错误请重试")
	UserNoFound     = comm2.NewApiBizError(2003, "账户不存在")
	PasswordError   = comm2.NewApiBizError(2004, "用户名或则密码错误")
	VerifyCodeLimit = comm2.NewApiBizError(2004, "获取验证码频率太高了，请稍后重试")
)

type LoginDao interface {
	Login(ctx *route.Context, data map[string]interface{}) error
}
