package messages

import (
	"github.com/glide-im/api/internal/dao"
	"github.com/glide-im/api/internal/dao/msgdao"
	"github.com/glide-im/api/internal/pkg/db"
)

type MessageDao struct {
}

var MessageDaoH MessageDao

func (receiver *MessageDao) GetMessages(from int64, to int64, pageSize int, page int, end_mid int64, start_mid int64) []msgdao.ChatMessage {
	var chatMessages []msgdao.ChatMessage
	query := db.DB.Model(msgdao.ChatMessage{}).Order("m_id desc")
	if pageSize == 0 {
		pageSize = 200
		page = 1
	}
	if to > 0 {
		session_id := dao.GetSessionId(to, from)
		query = query.Where("`session_id` = ? ", session_id)
	} else {
		query = query.Where("`from` = ? or `to` = ?", from, from)
	}

	query = query.Limit(pageSize).Offset(page - 1)
	if end_mid > 0 {
		query = query.Where("`m_id` > ?", end_mid)
	}

	if start_mid > 0 {
		query = query.Where("`m_id` < ?", start_mid)
	}

	query.Find(&chatMessages)
	return chatMessages
}

func (receiver *MessageDao) MessageRead(session_id string, uid int64) {
	db.DB.Model(msgdao.ChatMessage{}).Where("`session_id` = ?", session_id).Where("`to` = ?", uid).UpdateColumn("status", 1)
}
