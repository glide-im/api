package auth

import (
	"github.com/glide-im/api/internal/pkg/validate"
)

type AuthTokenRequest struct {
	Token string
}

type SignInRequest struct {
	Device   int64
	Account  string
	Password string
}

type LogoutRequest struct {
}

type RegisterRequest struct {
	Account  string
	Nickname string
	Password string
}

type GuestRegisterRequest struct {
	Avatar   string
	Nickname string
}

type GuestRegisterV2Request struct {
	FingerprintId string `json:"fingerprint_id" validate:"required"`
}

// AuthResponse login or register result
type AuthResponse struct {
	Token   string
	Uid     int64
	Servers []string
}

func (request *GuestRegisterV2Request) Validate() error {
	return validate.ValidateHandle(request)
}
