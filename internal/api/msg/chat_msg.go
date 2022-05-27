package msg

import (
	comm2 "github.com/glide-im/api/internal/api/comm"
	"github.com/glide-im/api/internal/api/router"
	"github.com/glide-im/api/internal/dao/msgdao"
	"github.com/glide-im/glide/pkg/logger"
	"github.com/glide-im/glide/pkg/messages"
	"math"
	"time"
)

type ChatMsgApi struct{}

//goland:noinspection GoPreferNilSlice
func (*ChatMsgApi) GetRecentChatMessage(ctx *route.Context, request *RecentChatMessageRequest) error {
	ms, err := msgdao.ChatMsgDaoImpl.GetRecentChatMessagesBySession(ctx.Uid, request.Uid, 10)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	msr := []*MessageResponse{}
	for _, m := range ms {
		msr = append(msr, messageModel2MessageResponse(m))
	}
	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, msr))
	return nil
}

//goland:noinspection GoPreferNilSlice
func (*ChatMsgApi) GetChatMessageHistory(ctx *route.Context, request *ChatHistoryRequest) error {

	if request.BeforeMid == 0 {
		request.BeforeMid = math.MaxInt64
	}
	ms, err := msgdao.ChatMsgDaoImpl.GetChatMessagesBySession(ctx.Uid, request.Uid, request.BeforeMid, 20)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	msr := []*MessageResponse{}
	for _, m := range ms {
		msr = append(msr, messageModel2MessageResponse(m))
	}
	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, msr))
	return nil
}

//goland:noinspection GoPreferNilSlice
func (*ChatMsgApi) GetRecentMessage(ctx *route.Context) error {
	ms, err := msgdao.ChatMsgDaoImpl.GetRecentChatMessages(ctx.Uid, time.Now().Unix()-int64(time.Hour*3*24))
	if err != nil {
		return comm2.NewDbErr(err)
	}
	msr := []*MessageResponse{}
	for _, m := range ms {
		msr = append(msr, messageModel2MessageResponse(m))
	}
	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, msr))
	return nil
}

//goland:noinspection GoPreferNilSlice
func (*ChatMsgApi) GetRecentMessageByUser(ctx *route.Context, request *RecentMessageRequest) error {
	resp := []RecentMessagesResponse{}
	var e = 0
	for _, i := range request.Uid {
		ms, err := msgdao.ChatMsgDaoImpl.GetChatMessagesBySession(ctx.Uid, i, 0, 20)
		if err != nil {
			logger.E("GetRecentMessageByUser DB error %v", err)
			e++
			continue
		}
		msr := []*MessageResponse{}
		for _, m := range ms {
			msr = append(msr, messageModel2MessageResponse(m))
		}
		resp = append(resp, RecentMessagesResponse{
			Uid:      i,
			Messages: msr,
		})
	}
	if e == len(request.Uid) {
		return errRecentMsgLoadFailed
	}
	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, resp))
	return nil
}

func (*ChatMsgApi) AckOfflineMessage(ctx *route.Context, request *AckOfflineMessageRequest) error {
	err := msgdao.DelOfflineMessage(ctx.Uid, request.Mid)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	return nil
}

//goland:noinspection GoPreferNilSlice
func (*ChatMsgApi) GetOfflineMessage(ctx *route.Context) error {
	oms, err := msgdao.GetOfflineMessage(ctx.Uid)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	var mid = []int64{}
	for _, m := range oms {
		mid = append(mid, m.MID)
	}
	qms, err := msgdao.GetChatMessage(mid...)
	var ms = []*MessageResponse{}
	for _, m := range qms {
		ms = append(ms, messageModel2MessageResponse(m))
	}
	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, ms))
	return nil
}

func messageModel2MessageResponse(m *msgdao.ChatMessage) *MessageResponse {
	return &MessageResponse{
		Mid:      m.MID,
		CliSeq:   m.CliSeq,
		From:     m.From,
		To:       m.To,
		Type:     m.Type,
		SendAt:   m.SendAt,
		CreateAt: m.CreateAt,
		Content:  m.Content,
		Status:   m.Status,
	}
}
