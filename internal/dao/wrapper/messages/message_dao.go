package messages

import (
	"github.com/glide-im/api/internal/dao/msgdao"
	"github.com/glide-im/api/internal/pkg/db"
)

type MessageDao struct {
}

var MessageDaoH MessageDao

func (receiver *MessageDao) GetMessages(from int64, to int64, pageSize int, page int, end_mid int64, start_mid int64) []msgdao.ChatMessage {
	var chatMessages []msgdao.ChatMessage
	query := db.DB.Model(msgdao.ChatMessage{}).Where("(`from` = ? AND `to` = ?) or (`from` = ? AND `to` = ?)", from, to, to, from).Order("m_id desc")
	if pageSize == 0 {
		pageSize = 200
		page = 1
	}

	query = query.Limit(pageSize).Offset(page)
	if end_mid > 0 {
		query = query.Where("mid > ?", end_mid)
	}

	if start_mid > 0 {
		query = query.Where("m_id < ?", start_mid)
	}

	query.Find(&chatMessages)
	return chatMessages
}

func (receiver *MessageDao) MessageRead(m_ids []int64) {
	db.DB.Model(msgdao.ChatMessage{}).Where("m_id in (?)", m_ids).UpdateColumn("status", 1)
}
