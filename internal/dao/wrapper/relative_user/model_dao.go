package relative_user

import (
	"github.com/glide-im/api/internal/dao/userdao"
	"github.com/glide-im/api/internal/pkg/db"
	"gorm.io/gorm"
)

type RelativeUser struct {
	gorm.Model
	Id          int64            `json:"id"`
	PrimaryUid  string           `json:"primary_uid"`
	RelativeUid string           `json:"relative_id"`
	Type        int8             `json:"type"`
	UserInfo    userdao.TripUser `json:"user_info" gorm:"foreignKey:RelativeUid;references:Uid"`
}

const (
	WHITE_TYPE = 1
	BLACK_TYPE = 2
)

type RelativeUserH struct {
}

func (m *RelativeUserH) SetBlackLists(primaryUid string, relativeUids []string) error {
	var relativeUsers []*RelativeUser
	for _, relativeUid := range relativeUids {
		var relativeUser = &RelativeUser{
			PrimaryUid:  primaryUid,
			RelativeUid: relativeUid,
			Type:        BLACK_TYPE,
		}
		if (db.DB.Model(&RelativeUser{}).Where(&relativeUser).RowsAffected > 0) {
			continue
		}

		relativeUsers = append(relativeUsers, relativeUser)
	}
	if err := db.DB.Model(&RelativeUser{}).Create(relativeUsers).Error; err != nil {
		return err
	}

	return nil
}

func (m *RelativeUserH) RemoveBlackLists(primaryUid string, relativeUids []string) {
	for _, relativeUid := range relativeUids {
		db.DB.Model(&RelativeUser{}).Where("primary_uid = ? and relative_uid = ? and type = ?", primaryUid, relativeUid, BLACK_TYPE).Delete(&RelativeUser{})
	}
}

func (m *RelativeUserH) GetBlackLists(primaryUid string) []RelativeUser {
	var relativeUsers []RelativeUser
	db.DB.Model(&RelativeUser{}).Where("primary_uid = ? and type = ?", primaryUid, BLACK_TYPE).Preload("UserInfo").Find(&relativeUsers)
	return relativeUsers
}

func (m *RelativeUserH) IsUserInBlackList(primaryUid string, relativeUid string) bool {
	return db.DB.Model(&RelativeUser{}).Where("primary_uid = ? and relative_id = ? and type = ?", primaryUid, relativeUid, BLACK_TYPE).RowsAffected > 0
}
