package msg

import (
	"errors"
	"fmt"
	comm2 "github.com/glide-im/api/internal/api/comm"
	"github.com/glide-im/api/internal/api/router"
	"github.com/glide-im/api/internal/dao/msgdao"
	"github.com/glide-im/api/internal/dao/userdao"
	"github.com/glide-im/api/internal/dao/wrapper/category"
	"github.com/glide-im/api/internal/im"
	"github.com/glide-im/api/internal/pkg/db"
	"github.com/glide-im/glide/pkg/hash"
	"github.com/glide-im/glide/pkg/logger"
	"github.com/glide-im/glide/pkg/messages"
	"math/rand"
	"strconv"
	"time"
)

func (*MsgApi) ReadMessage(ctx *route.Context, request *ReadMessageRequest) error {
	//err := msgdao.SessionDaoImpl.CleanUserSessionUnread(ctx.Uid, request.To, request.To)
	//if err != nil {
	//	return comm.NewDbErr(err)
	//}
	return nil
}

func (*MsgApi) GetSessionTicket(ctx *route.Context, r *GetTicketRequest) error {

	myId := strconv.FormatInt(ctx.Uid, 10)
	inBlackList, err := msgdao.SessionDaoImpl.IsUserInBlackList(r.To, myId)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	if inBlackList {
		return errors.New("your are in black list")
	}

	justContact, err := msgdao.SessionDaoImpl.IsJustReceiveMessageFromContact(r.To)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	if justContact {
		inWhiteList, err := msgdao.SessionDaoImpl.IsUserInWhiteList(r.To, myId)
		if err != nil {
			return comm2.NewDbErr(err)
		}
		if !inWhiteList {
			return errors.New("your are not in white list")
		}
	}

	user, err := userdao.Dao.GetUser(ctx.Uid)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	if user == nil {
		return errors.New("user not found")
	}
	secret := user.MessageDeliverSecret

	sum1 := hash.SHA1(secret + r.To)
	ticket := hash.SHA1(secret + strconv.FormatInt(ctx.Uid, 10) + sum1)

	ctx.ReturnSuccess(&TicketResponse{Ticket: ticket})
	return nil
}

func (*MsgApi) AddToBlackList(ctx *route.Context, request *SessionBlackListRequest) error {
	// TODO
	return nil
}

func (*MsgApi) AddToWhiteList(ctx *route.Context, request *SessionBlackListRequest) error {
	// TODO
	return nil
}
func (*MsgApi) RemoveFromWhiteList(ctx *route.Context, request *SessionBlackListRequest) error {
	// TODO
	return nil
}

func (*MsgApi) RemoveFromBlackList(ctx *route.Context, request *SessionBlackListRequest) error {
	// TODO
	return nil
}

func (*MsgApi) GetBlackList(ctx *route.Context) error {

	myId := strconv.FormatInt(ctx.Uid, 10)
	list, err := msgdao.SessionDaoImpl.GetUserBlackList(myId)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	ctx.ReturnSuccess(list)
	return nil
}

func (*MsgApi) GetWhiteList(ctx *route.Context) error {

	return nil
}

func (*MsgApi) UpdateBlackList(ctx *route.Context, request *SessionBlackListRequest) error {

	var failed []string

	myId := strconv.FormatInt(ctx.Uid, 10)
	err := msgdao.SessionDaoImpl.UpdateUsrBlackList(myId, request.blackList)
	if err != nil {
		return err
	}

	for _, to := range request.blackList {
		secret := randomStr(32)

		id, err := strconv.ParseInt(to, 10, 64)
		if err != nil {
			failed = append(failed, to)
			continue
		}
		err = userdao.Dao.UpdateSecret(id, secret)
		if err != nil {
			failed = append(failed, to)
			continue
		}
		err = im.IM.UpdateClientSecret(to, secret)
		if err != nil {
			logger.E("error on update session permission %v", err)
		}
	}
	return nil
}

func (*MsgApi) GetOrCreateSession(ctx *route.Context, request *SessionRequest) error {
	session, err := msgdao.SessionDaoImpl.GetSession(ctx.Uid, request.To)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	fmt.Println("session", session)
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
	var categoryUsers []category.CategoryUser
	var uids []int64

	for _, s := range session {
		to := s.Uid2
		if s.Uid2 == ctx.Uid {
			to = s.Uid
		}
		uids = append(uids, to)
	}
	db.DB.Model(category.CategoryUser{}).Where("app_id = ?", ctx.AppID).Where("uid in ?", uids).Select("uid, category_id").Find(&categoryUsers)
	for _, s := range session {
		to := s.Uid2
		if s.Uid2 == ctx.Uid {
			to = s.Uid
		}
		var cids []int64
		for _, cate := range categoryUsers {
			fmt.Println("cate.UId == to", cate.UId, to)
			if cate.UId == to {
				cids = append(cids, cate.CategoryId)
			}
		}

		sr = append(sr, &SessionResponse{
			Uid2:        s.Uid,
			Uid1:        s.Uid2,
			To:          to,
			CategoryIds: cids,
			LastMid:     s.LastMID,
			UpdateAt:    s.UpdateAt,
			CreateAt:    s.CreateAt,
		})
		if s.LastMID > 0 {
			mids = append(mids, s.LastMID)
		}
	}

	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, sr))
	return nil
}

func randomStr(n int) string {
	var l = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	length := len(l)
	for i := range b {
		b[i] = l[rand.Intn(length)]
	}
	return string(b)
}
