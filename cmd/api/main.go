package main

import (
	"github.com/glide-im/api/internal/api"
	"github.com/glide-im/api/internal/auth"
	"github.com/glide-im/api/internal/config"
	"github.com/glide-im/api/internal/dao"
	"github.com/glide-im/api/internal/im"
	"github.com/glide-im/api/internal/pkg/db"
	"github.com/glide-im/api/internal/pkg/validate"
)

func main() {
	config.MustLoad()

	db.Init()
	dao.Init()
	validate.Init()

	secret := config.ApiHttp.JwtSecret
	auth.SetJwtSecret([]byte(secret))

	im.MustSetupClient(config.IMRpcServer.Addr, config.IMRpcServer.Port, config.IMRpcServer.Name)
	err := api.Run(config.ApiHttp.Addr, config.ApiHttp.Port)

	if err != nil {
		panic(err)
	}
}
