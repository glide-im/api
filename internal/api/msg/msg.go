package msg

import (
	comm2 "github.com/glide-im/api/internal/api/comm"
	"github.com/glide-im/api/internal/api/router"
	"github.com/glide-im/api/internal/dao/msgdao"
	"github.com/glide-im/glide/pkg/messages"
)

type MsgApi struct {
	*GroupMsgApi
	*ChatMsgApi
}

func (MsgApi) GetMessageID(ctx *route.Context) error {
	id, err := msgdao.GetMessageID()
	if err != nil {
		return comm2.NewDbErr(err)
	}
	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, MessageIDResponse{id}))
	return nil
}
