package main

import (
	"github.com/glide-im/api/internal/config"
	"github.com/glide-im/api/internal/dao"
	"github.com/glide-im/api/internal/dao/msgdao"
	"github.com/glide-im/api/internal/dao/userdao"
	"github.com/glide-im/api/internal/pkg/db"
)

func main() {
	config.MustLoad()

	db.Init()
	dao.Init()

	InitDatabase()
}

func InitDatabase() {
	_ = db.DB.AutoMigrate(userdao.User{})
	_ = db.DB.AutoMigrate(userdao.Contacts{})
	_ = db.DB.AutoMigrate(msgdao.ChatMessage{})
	_ = db.DB.AutoMigrate(msgdao.OfflineMessage{})
	_ = db.DB.AutoMigrate(msgdao.GroupMessage{})
	_ = db.DB.AutoMigrate(msgdao.GroupMemberMsgState{})
	_ = db.DB.AutoMigrate(msgdao.GroupMsgSeq{})
	_ = db.DB.AutoMigrate(msgdao.GroupMessageState{})
	_ = db.DB.AutoMigrate(msgdao.Session{})
}
