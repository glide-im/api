package articles

import (
	"github.com/glide-im/api/internal/pkg/db"
	"gorm.io/gorm"
)

type Article struct {
	AppID     int64  `json:"app_id,omitempty"`
	Uid       int64  `json:"uid"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	PublishAt string `json:"publish_at"`
	Weight    int64  `json:"weight"`
}

var ArticleDao = &ArticleH{}

type ArticleH struct {
}

func (a *ArticleH) GetModel(app_id int64, uid int64) *gorm.DB {
	return db.DB.Model(&Article{}).Where("app_id = ? and uid = ?", app_id, uid)
}
