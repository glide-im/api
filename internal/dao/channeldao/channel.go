package channeldao

import (
	"github.com/glide-im/api/internal/pkg/db"
	"time"
)

var Dao = NewChannelDao()

type ChannelDao struct {
}

func NewChannelDao() *ChannelDao {
	return &ChannelDao{}
}

func (c *ChannelDao) AutoMigrate() error {
	err := db.DB.AutoMigrate(&ChannelModel{})
	if err != nil {
		return err
	}
	err = db.DB.AutoMigrate(&ChannelMemberModel{})
	return err
}

func (c *ChannelDao) CreateChannel(id string, name string, avatar string) error {
	model := &ChannelModel{
		Name:      name,
		ChanId:    id,
		Avatar:    avatar,
		Muted:     false,
		Type:      0,
		ReadOnly:  false,
		Access:    0,
		UpdatedAt: time.Now().Unix(),
		CreatedAt: time.Now().Unix(),
	}
	query := db.DB.Create(model)
	return query.Error
}

func (c *ChannelDao) GetChannel(id string) (*ChannelModel, error) {
	model := &ChannelModel{}
	query := db.DB.Where("chan_id = ?", id).First(model)
	return model, query.Error
}

func (c *ChannelDao) CreateChannelMember(chanId string, tp int, perm int64, uid string, memberId string) error {
	model := &ChannelMemberModel{
		MemberId:  memberId,
		ChanId:    chanId,
		Uid:       uid,
		Muted:     false,
		Type:      tp,
		Perm:      perm,
		UpdatedAt: time.Now().Unix(),
		CreatedAt: time.Now().Unix(),
	}
	query := db.DB.Create(model)
	return query.Error
}

func (c *ChannelDao) GetChannelMembers(chanId string) ([]*ChannelMemberModel, error) {
	var models []*ChannelMemberModel
	query := db.DB.Where("chan_id = ?", chanId).Find(&models)
	return models, query.Error
}

func (c *ChannelDao) GetChannelMember(chanId string, uid string) (*ChannelMemberModel, error) {
	model := &ChannelMemberModel{}
	query := db.DB.Where("chan_id = ? AND uid = ?", chanId, uid).First(model)
	return model, query.Error
}

func (c *ChannelDao) GetChannelsByUid(uid string) ([]*ChannelMemberModel, error) {
	var models []*ChannelMemberModel
	query := db.DB.Model(&ChannelMemberModel{}).Where("uid = ?", uid).Find(&models)
	return models, query.Error
}

func (c *ChannelDao) DeleteChannel(chanId string) error {
	query := db.DB.Where("chan_id = ?", chanId).Delete(&ChannelModel{})
	return query.Error
}

func (c *ChannelDao) DeleteChannelMemberByUid(chanId string, uid string) error {
	query := db.DB.Where("chan_id = ? AND uid = ?", chanId, uid).Delete(&ChannelMemberModel{})
	return query.Error
}

func (c *ChannelDao) DeleteMemberByMemberId(memberId string) error {
	query := db.DB.Where("member_id = ?", memberId).Delete(&ChannelMemberModel{})
	return query.Error
}

func (c *ChannelDao) UpdateChannel(chanId string, name string, avatar string) error {
	query := db.DB.Model(&ChannelModel{}).Where("chan_id = ?", chanId).Updates(map[string]interface{}{
		"name":       name,
		"avatar":     avatar,
		"updated_at": time.Now().Unix(),
	})
	return query.Error
}
