package collect

import (
	"github.com/glide-im/api/internal/pkg/db"
	"gorm.io/gorm"
)

type CollectData struct {
	AppID   int64  `json:"app_id,omitempty"`
	Uid     int64  `json:"uid"`
	Ip      string `json:"ip"`
	Region  string `json:"region"`
	Browser string `json:"browser"`
	Device  string `json:"device"`
	Origin  string `json:"origin"`
}

var CollectDataDao = &CollectDataDaoH{}

type CollectDataDaoH struct {
}

func (a *CollectDataDaoH) GetModel(app_id int64, uid int64) *gorm.DB {
	return db.DB.Model(&CollectData{}).Where("app_id = ? and uid = ?", app_id, uid)
}

func (a *CollectDataDaoH) updateOrCreate(CollectData CollectData) *gorm.DB {
	_collectData := CollectData{}
	db.DB.Model(&CollectData{}).Where("app_id = ?", CollectData.AppID, uid)
}
