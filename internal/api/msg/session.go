package msg

import (
	"fmt"
	comm2 "github.com/glide-im/api/internal/api/comm"
	"github.com/glide-im/api/internal/api/router"
	"github.com/glide-im/api/internal/dao/msgdao"
	"github.com/glide-im/glide/pkg/messages"
	"time"
)

func (*MsgApi) ReadMessage(ctx *route.Context, request *ReadMessageRequest) error {
	//err := msgdao.SessionDaoImpl.CleanUserSessionUnread(ctx.Uid, request.To, request.To)
	//if err != nil {
	//	return comm.NewDbErr(err)
	//}
	return nil
}

func (*MsgApi) GetOrCreateSession(ctx *route.Context, request *SessionRequest) error {
	session, err := msgdao.SessionDaoImpl.GetSession(ctx.Uid, request.To)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	if session.SessionId == "" {
		se, err := msgdao.SessionDaoImpl.CreateSession(ctx.Uid, request.To, time.Now().Unix())
		if err != nil {
			return comm2.NewDbErr(err)
		}
		session = se
	}

	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, SessionResponse{
		Uid1:     session.Uid,
		Uid2:     session.Uid2,
		LastMid:  session.LastMID,
		UpdateAt: session.UpdateAt,
		CreateAt: session.CreateAt,
	}))

	return nil
}

func (a *MsgApi) GetRecentSessions(ctx *route.Context) error {
	now := time.Now().Unix() + 100
	session, err := msgdao.SessionDaoImpl.GetRecentSession(ctx.Uid, now, 100)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	//goland:noinspection GoPreferNilSlice
	sr := []*SessionResponse{}
	var mids []int64
	for _, s := range session {
		to := s.Uid2
		if s.Uid2 == ctx.Uid {
			to = s.Uid
		}

		sr = append(sr, &SessionResponse{
			Uid2:     s.Uid,
			Uid1:     s.Uid2,
			To:       to,
			LastMid:  s.LastMID,
			UpdateAt: s.UpdateAt,
			CreateAt: s.CreateAt,
		})
		mids = append(mids, s.LastMID)
	}
	fmt.Println(mids)
	err, _messages := msgdao.SessionDaoImpl.GetMessagesByMids(mids)
	if err != nil {
		return err
	}
	for _, s := range sr {
		s.Message = _messages[s.LastMid]
	}
	fmt.Println(sr)

	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, sr))
	return nil
}
