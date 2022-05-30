package im

import (
	"github.com/glide-im/glide/pkg/logger"
	"github.com/glide-im/glide/pkg/messages"
)

// ClientInterface 客户端连接相关接口
var ClientInterface ClientManagerInterface = &clientInterface{}

func SendMessage(uid int64, device int64, m *messages.GlideMessage) {
	err := ClientInterface.EnqueueMessage(uid, device, m)
	if err != nil {
		logger.E("SendMessage error: %v", err)
	}
}

type ClientManagerInterface interface {
	Logout(uid int64, device int64) error
	EnqueueMessage(uid int64, device int64, message *messages.GlideMessage) error
}

type clientInterface struct{}

func (c clientInterface) Logout(uid int64, device int64) error {
	// TODO: 调用 IM 服务接口
	return nil
}

func (c clientInterface) EnqueueMessage(uid int64, device int64, message *messages.GlideMessage) error {
	logger.W("clientInterface.EnqueueMessage not implement")
	// TODO: 调用 IM 服务接口
	return nil
}
