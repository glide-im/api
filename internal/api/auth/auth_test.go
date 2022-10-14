package auth

import (
	"github.com/glide-im/api/internal/api/router"
	"github.com/glide-im/api/internal/config"
	"github.com/glide-im/api/internal/pkg/db"
	"github.com/glide-im/glide/pkg/logger"
	"github.com/glide-im/glide/pkg/messages"
	"testing"
)

var authApi = AuthApi{}

func init() {
	config.MustLoad()
	db.Init()
}

func getContext(uid int64, device int64) *route.Context {
	return &route.Context{
		Uid:    uid,
		Device: device,
		Seq:    1,
		Action: "",
		R: func(message *messages.GlideMessage) {
			logger.D("Response=%v", message)
		},
	}
}

func logErr(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func TestGuestLogin(t *testing.T) {
	err := authApi.GuestRegister(getContext(0, 0), &GuestRegisterRequest{
		Avatar:   "a",
		Nickname: "asdf",
	})
	logErr(t, err)
}

func TestAuthApi_AuthToken(t *testing.T) {
	err := authApi.AuthToken(getContext(2, 0), &AuthTokenRequest{
		Token: "RN9fXQtAoplDCX8uSiajitgFgCZlrcpX",
	})
	logErr(t, err)
}

func TestAuthApi_Register(t *testing.T) {
	err := authApi.Register(getContext(2, 0), &RegisterRequest{
		Email:    "bb",
		Password: "bb",
	})
	logErr(t, err)
}

func TestAuthApi_SignIn(t *testing.T) {
	err := authApi.SignIn(getContext(2, 0), &SignInRequest{
		Email:    "aa",
		Password: "1234567",
		Device:   1,
	})
	logErr(t, err)
}

func TestAuthApi_Logout(t *testing.T) {
	err := authApi.Logout(getContext(543603, 1))
	logErr(t, err)
}
