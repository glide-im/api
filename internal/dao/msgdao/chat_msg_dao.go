package msgdao

import (
	"github.com/glide-im/api/internal/dao"
	"github.com/glide-im/api/internal/dao/common"
	"github.com/glide-im/api/internal/pkg/db"
	"strconv"
)

var ChatMsgDaoImpl ChatMsgDao = chatMsgDaoImpl{}

const (
	ChatMessageStatusDefault  = 0
	ChatMessageStatusRecalled = 1 // 撤回
	ChatMessageStatusDisabled = 2
	ChatMessageStatusRead     = 3 // 已读
)

type chatMsgDaoImpl struct {
}

func (chatMsgDaoImpl) GetRecentChatMessagesBySession(uid1, uid2 int64, pageSize int) ([]*ChatMessage, error) {
	sid, _, _ := getSessionId(uid2, uid1)
	var ms []*ChatMessage
	query := db.DB.Model(&ChatMessage{}).
		Where("`session_id` = ?", sid).
		Order("`send_at` DESC").
		Limit(pageSize).
		Find(&ms)
	if query.Error != nil {
		return nil, query.Error
	}
	return ms, nil
}

func (chatMsgDaoImpl) GetChatMessagesBySession(uid1, uid2 int64, beforeMid int64, pageSize int) ([]*ChatMessage, error) {
	sid, _, _ := getSessionId(uid2, uid1)
	var ms []*ChatMessage
	query := db.DB.Model(&ChatMessage{}).
		Where("`session_id` = ? AND `m_id` < ?", sid, beforeMid).
		Order("`send_at` DESC").
		Limit(pageSize).
		Find(&ms)
	if query.Error != nil {
		return nil, query.Error
	}
	return ms, nil
}

func (chatMsgDaoImpl) GetRecentChatMessages(uid int64, after int64) ([]*ChatMessage, error) {
	var ms []*ChatMessage
	query := db.DB.Model(&ChatMessage{}).Where("`from` = ? OR `to` = ? AND `send_at` > ?", uid, uid, after).Find(&ms)
	if query.Error != nil {
		return nil, query.Error
	}
	return ms, nil
}

func (chatMsgDaoImpl) GetChatMessage(mid ...int64) ([]*ChatMessage, error) {
	//goland:noinspection GoPreferNilSlice
	m := []*ChatMessage{}
	query := db.DB.Model(m).Where("m_id in (?)", mid).Find(&m)
	if err := common.ResolveError(query); err != nil {
		return nil, err
	}
	return m, nil
}

func (chatMsgDaoImpl) UpdateChatMessageStatus(mid int64, from, to int64, status int) error {
	u := ChatMessage{}
	update := db.DB.Model(&u).Where("`m_id` = ? AND `from` = ? AND `to` = ?", mid, from, to).UpdateColumn("status", status)
	return common.JustError(update)
}

func (chatMsgDaoImpl) AddChatMessage(message *ChatMessage) (bool, error) {
	var c int64
	query := db.DB.Table("im_chat_message").Where("m_id = ?", message.MID).Count(&c)
	if err := common.ResolveError(query); err != nil {
		return false, err
	}
	if c > 0 {
		return false, nil
	}
	query = db.DB.Create(message)
	if err := common.ResolveError(query); err != nil {
		return false, err
	}
	return true, nil
}

func (chatMsgDaoImpl) GetChatMessageMidAfter(from, to int64, midAfter int64) ([]*ChatMessage, error) {
	lg, sm := from, to
	if lg < sm {
		lg, sm = sm, lg
	}
	sid := strconv.FormatInt(lg, 10) + "_" + strconv.FormatInt(sm, 10)
	var ms []*ChatMessage
	query := db.DB.Model(&ChatMessage{}).Where("session_id = ? and m_id > ?", sid, midAfter).Find(&ms)
	if err := common.ResolveError(query); err != nil {
		return nil, err
	}
	return ms, nil
}

func (chatMsgDaoImpl) GetChatMessageMidSpan(from, to int64, midStart, midEnd int64) ([]*ChatMessage, error) {
	lg, sm := from, to
	if lg < sm {
		lg, sm = sm, lg
	}
	sid := strconv.FormatInt(lg, 10) + "_" + strconv.FormatInt(sm, 10)
	var ms []*ChatMessage
	query := db.DB.Model(&ChatMessage{}).Where("sid = ? AND m_id >= ? AND m_id < ?", sid, midStart, midEnd).Find(&ms)
	if err := common.ResolveError(query); err != nil {
		return nil, err
	}
	return ms, nil
}

func (chatMsgDaoImpl) AddOfflineMessage(uid int64, mid int64) error {
	offlineMessage := &OfflineMessage{
		MID: mid,
		UID: uid,
	}
	query := db.DB.Create(offlineMessage)
	return common.ResolveError(query)
}

func (chatMsgDaoImpl) GetOfflineMessage(uid int64) ([]*OfflineMessage, error) {
	var m []*OfflineMessage
	query := db.DB.Model(&OfflineMessage{}).Where("uid = ?", uid).Find(&m)
	if query.Error != nil {
		return nil, query.Error
	}
	return m, nil
}

func (chatMsgDaoImpl) DelOfflineMessage(uid int64, mid []int64) error {
	query := db.DB.Where("uid = ? AND m_id IN (?)", uid, mid).Delete(&OfflineMessage{})
	return query.Error
}

func (chatMsgDaoImpl) GetChatLastMessage(from int64, to int64) ChatMessage {
	var message ChatMessage
	db.DB.Model(ChatMessage{}).Where("(`from` = ? AND `to` = ?) or (`from` = ? AND `to` = ?)", from, to, to, from).Order("m_id desc").Last(&message)
	return message
}

func (chatMsgDaoImpl) GetMessageCount(from int64, to int64) int64 {
	session_id := dao.GetSessionId(from, to)
	var count int64
	db.DB.Model(ChatMessage{}).Where("`session_id` = ?", session_id).Where("to = ?", from).Where("`status` = 0").Count(&count)
	return count
}
