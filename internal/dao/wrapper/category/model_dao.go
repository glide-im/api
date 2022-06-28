package category

import (
	"github.com/glide-im/api/internal/pkg/db"
	"gorm.io/gorm"
)

type Category struct {
	AppID  int64  `json:"app_id,omitempty"`
	Name   string `json:"title"`
	Weight int64  `json:"weight"`
	Icon   string `json:"icon"`
}

var CategoryDao = &CategoryH{}

type CategoryH struct {
}

func (a *CategoryH) GetModel(app_id int64) *gorm.DB {
	return db.DB.Model(&Category{}).Where("app_id = ?", app_id)
}
