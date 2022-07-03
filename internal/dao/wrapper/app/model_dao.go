package app

import (
	"github.com/glide-im/api/internal/pkg/db"
	"gorm.io/gorm"
)

type App struct {
	AppID   int64  `json:"app_id,omitempty"`
	Name    int64  `json:"name"`
	Uid     int64  `json:"uid"`
	License string `json:"license"`
	Status  int    `json:"status"`
	Logo    string `json:"logo"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

var AppDao = &AppH{}

type AppH struct {
}

func (a *AppH) GetModel(app_id int64, uid int64) *gorm.DB {
	return db.DB.Model(&App{}).Where("app_id = ? and uid = ?", app_id, uid)
}
