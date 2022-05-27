package main

import (
	"github.com/glide-im/api/config"
	"github.com/glide-im/api/internal/api"
	"github.com/glide-im/api/internal/auth"
	"github.com/glide-im/api/internal/dao"
	"github.com/glide-im/api/internal/pkg/db"
)

func main() {
	config.MustLoad()

	db.Init()
	dao.Init()

	secret := config.ApiHttp.JwtSecret
	auth.SetJwtSecret([]byte(secret))

	err := api.RunHttpServer(config.ApiHttp.Addr, config.ApiHttp.Port)

	if err != nil {
		panic(err)
	}
}
