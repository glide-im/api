package auth

import (
	"github.com/glide-im/api/internal/dao/wrapper/app"
	"github.com/glide-im/api/internal/pkg/validate"
	"github.com/glide-im/glide/pkg/gate"
)

type AuthTokenRequest struct {
	Token string `json:"token"`
}

type SignInRequest struct {
	Device   int64  `json:"device"`
	Email    string `json:"email"  validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LogoutRequest struct {
}

type RegisterRequest struct {
	Account  string `json:"account"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=16,min=6"`
	Captcha  string `json:"captcha" validate:"required"`
	Nickname string `json:"nickname" validate:"required"`
}

type ForgetRequest struct {
	Email   string `json:"email" validate:"required,email"`
	Captcha string `json:"captcha" validate:"required"`
}

type GuestRegisterRequest struct {
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
}

type VerifyCodeRequest struct {
	Email string `json:"email" validate:"required,email"`
	Mode  string `json:"mode"`
}

type GuestRegisterV2Request struct {
	FingerprintId string `json:"fingerprint_id" validate:"required"`
	Origin        string `json:"origin"`
}

// AuthResponse login or register result
type AuthResponse struct {
	Token      string                    `json:"token"`
	Uid        int64                     `json:"uid"`
	Servers    []string                  `json:"servers"`
	NickName   string                    `json:"nick_name"`
	App        app.App                   `json:"app"`
	Email      string                    `json:"email"`
	Phone      string                    `json:"phone"`
	Device     int64                     `json:"device"`
	Credential *gate.EncryptedCredential `json:"credential,omitempty"`
}

// GuestAuthResponse login or register result
type GuestAuthResponse struct {
	Token      string                    `json:"token"`
	Uid        int64                     `json:"uid"`
	Servers    []string                  `json:"servers"`
	AppID      int64                     `json:"app_id"`
	NickName   string                    `json:"nick_name"`
	Credential *gate.EncryptedCredential `json:"credential,omitempty"`
}

func (request *GuestRegisterV2Request) Validate() error {
	return validate.ValidateHandle(request)
}

func (request *RegisterRequest) Validate() error {
	return validate.ValidateHandle(request)
}
