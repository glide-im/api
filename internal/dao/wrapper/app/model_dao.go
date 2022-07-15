package app

import (
	"github.com/glide-im/api/internal/dao/common"
	"github.com/glide-im/api/internal/pkg/db"
	"gorm.io/gorm"
)

type App struct {
	Id    int64  `json:"id"`
	AppID string `json:"app_id,omitempty"`
	Name  string `json:"name"`
	Uid   int64  `json:"uid"`
	//License string `json:"license,omitempty"`
	Status int    `json:"status"`
	Logo   string `json:"logo"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Host   string `json:"host"`
}

var AppDao = &AppH{}

type AppH struct {
}

func (a *AppH) GetModel(app_id int64, uid int64) *gorm.DB {
	return db.DB.Model(&App{}).Where("app_id = ? and uid = ?", app_id, uid)
}

func (a *AppH) CheckExistHost(host string) *App {
	appModel := &App{}
	db.DB.Model(&App{}).Where("host = ?", host).First(&appModel)

	return appModel
}

func (a *AppH) AddApp(appModel *App) error {
	app_id := GenerateAppId()
	appModel.AppID = app_id
	query := db.DB.Model(&App{}).Create(&appModel)

	return common.JustError(query)
}

func (a *AppH) GetAppProfile(uid int64) App {
	app := App{}
	db.DB.Model(&App{}).Where("uid = ?", uid).Find(&app)
	return app
}

// 检查域名是否存在
func (a *AppH) GetAppID(host string) int64 {
	if len(host) == 0 {
		return 0
	}

	appModel := &App{}
	db.DB.Model(&App{}).Where("host = ?", host).First(&appModel)

	return appModel.Id
}
