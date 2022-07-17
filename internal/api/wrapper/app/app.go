package app

import (
	comm2 "github.com/glide-im/api/internal/api/comm"
	route "github.com/glide-im/api/internal/api/router"
	"github.com/glide-im/api/internal/dao/common"
	"github.com/glide-im/api/internal/dao/wrapper/app"
	"github.com/glide-im/api/internal/pkg/db"
	"math/rand"
	"time"
)

type PlatFromApi struct {
}

// 我名下的平台
func (a *PlatFromApi) List(ctx *route.Context) error {
	model := db.DB.Model(&app.App{})
	appList := []app.App{}
	model.Where("uid = ?", ctx.Uid).Find(&model)

	ctx.ReturnSuccess(appList)
	return nil
}

// 平台新增
func (a *PlatFromApi) Store(ctx *route.Context, request *AppStoreRequest) error {
	model := db.DB.Model(&app.App{})
	rand.Seed(time.Now().UnixNano())
	app_id := app.GenerateAppId()

	platformStore := app.App{
		AppID: app_id,
		Uid:   ctx.Uid,
		//License: request.License,
		Logo:  request.Logo,
		Email: request.Email,
		Name:  request.Name,
		Phone: request.Phone,
		Host:  request.Host,
	}

	appModel := app.AppDao.CheckExistHost(request.Host)
	if appModel.Uid > 0 {
		return comm2.NewApiBizError(3001, "域名已被配置")
	}

	_db := model.Create(platformStore)
	if err := common.JustError(_db); err != nil {
		return err
	}
	ctx.ReturnSuccess(platformStore)
	return nil
}

// 平台更新
func (a *PlatFromApi) Update(ctx *route.Context, request *AppStoreRequest) error {
	model := db.DB.Model(&app.App{})
	platformUpdate := app.App{
		//License: request.License,
		Logo:  request.Logo,
		Email: request.Email,
		Name:  request.Name,
		Phone: request.Phone,
	}
	id := ctx.Context.Param("id")
	model.Where("id = ? and uid = ?", id, ctx.Uid).Updates(platformUpdate)
	ctx.ReturnSuccess(platformUpdate)
	return nil
}

// 平台删除
func (a *PlatFromApi) Delete(ctx *route.Context) error {
	model := db.DB.Model(&app.App{})
	id := ctx.Context.Param("id")
	model.Where("id = ? and uid = ?", id, ctx.Uid).Delete(&app.App{})
	ctx.ReturnSuccess(nil)
	return nil
}

// 获取联络人
func (a *PlatFromApi) GetGuestToId(ctx *route.Context) error {
	model := db.DB.Model(&app.App{})
	var appModel app.App
	appid := ctx.AppID
	model.Where("id = ?", appid).Find(&appModel)
	ctx.ReturnSuccess(map[string]int64{
		"uid": appModel.Uid,
	})
	return nil
}
