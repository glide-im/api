package auth

import (
	"github.com/glide-im/api/internal/pkg/validate"
)

type AuthTokenRequest struct {
	Token string
}

type SignInRequest struct {
	Device   int64
	Email    string `json:"email" validate:"required,email"`
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
	Avatar   string
	Nickname string
}

type VerifyCodeRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type GuestRegisterV2Request struct {
	FingerprintId string `json:"fingerprint_id" validate:"required"`
	Origin        string `json:"origin"`
}

// AuthResponse login or register result
type AuthResponse struct {
	Token    string
	Uid      int64
	Servers  []string
	NickName string
}

// GuestAuthResponse login or register result
type GuestAuthResponse struct {
	Token    string
	Uid      int64
	Servers  []string
	AppID    int64
	NickName string
}

func (request *GuestRegisterV2Request) Validate() error {
	return validate.ValidateHandle(request)
}

func (request *RegisterRequest) Validate() error {
	return validate.ValidateHandle(request)
}
