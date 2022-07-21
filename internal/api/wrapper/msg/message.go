package msg

import (
	route "github.com/glide-im/api/internal/api/router"
	"github.com/glide-im/api/internal/dao/wrapper/messages"
)

type MessageApi struct {
}

func (m *MessageApi) GetMessageList(ctx *route.Context, request *MessageQueryRequest) error {
	list := messages.MessageDaoH.GetMessages(ctx.Uid, request.To, request.PageSize, request.Page, request.EndMid, request.StartMid)
	ctx.ReturnSuccess(list)
	return nil
}

func (m *MessageApi) MessageRead(ctx *route.Context, request *MessageQueryRequest) error {
	list := messages.MessageDaoH.GetMessages(ctx.Uid, request.To, request.PageSize, request.Page, request.EndMid, request.StartMid)
	ctx.ReturnSuccess(list)
	return nil
}
