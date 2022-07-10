package category

import (
	"github.com/glide-im/api/internal/dao/common"
	"github.com/glide-im/api/internal/pkg/db"
	"gorm.io/gorm"
)

type Category struct {
	Id     int64  `json:"id"`
	AppID  int64  `json:"app_id,omitempty"`
	Name   string `json:"name"`
	Weight int64  `json:"weight"`
	Icon   string `json:"icon"`
}

var CategoryDao = &CategoryH{}

type CategoryH struct {
}

func (a *CategoryH) GetModel(app_id int64) *gorm.DB {
	return db.DB.Model(&Category{}).Where("app_id = ?", app_id)
}

type CategoryUser struct {
	AppID      int64 `json:"app_id,omitempty"`
	CategoryId int64 `json:"category_id"`
	UId        int64 `json:"uid"`
	Form       int64 `json:"from"`
}

var CategoryUserDao = &CategoryUserH{}

type CategoryUserH struct {
}

func (s *CategoryUserH) GetModel(app_id int64) *gorm.DB {
	return db.DB.Model(&CategoryUser{}).Where("app_id = ?", app_id)
}

func (s *CategoryUserH) Updates(uid int64, category_ids []int64, selfId int64) error {
	var _db = db.DB
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		_db = s.GetModel(1).Where("uid = ?", uid).Delete(&Category{})
		if err := common.JustError(_db); err != nil {
			return err
		}

		var categories = []CategoryUser{}
		for _, category_id := range category_ids {
			categories = append(categories, CategoryUser{
				AppID:      1,
				CategoryId: category_id,
				UId:        uid,
				Form:       selfId,
			})
		}
		_db = db.DB.CreateInBatches(categories, 100)
		if err := common.JustError(_db); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
