package msgdao

import (
	"github.com/glide-im/api/internal/pkg/db"
)

var Comm CommonDao = commonDao{}

type commonDao struct {
}

func (commonDao) GetMessageID() (int64, error) {
	// TODO 2021-12-17 16:57:04
	//result, err := db.Redis.Incr("im:msg:id:incr").Result()
	//if err != nil {
	//	return 0, err
	//}
	var chatMessage ChatMessage
	db.DB.Model(ChatMessage{}).Order("`m_id` desc").Last(&chatMessage)
	return chatMessage.MID + 1, nil
}
