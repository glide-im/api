package user

import (
	"errors"
	comm2 "github.com/glide-im/api/internal/api/comm"
	"github.com/glide-im/api/internal/api/router"
	"github.com/glide-im/api/internal/dao/userdao"
	"github.com/glide-im/glide/pkg/messages"
)

type UserApi struct{}

func (a *UserApi) GetUserProfile(msg *route.Context) error {
	// TODO 2021-11-29 我的详细信息
	return nil
}

func (a *UserApi) UpdateUserProfile(msg *route.Context, request *UpdateProfileRequest) error {
	// TODO 2021-11-29 更新我的信息
	return nil
}

func (a *UserApi) GetUserInfo(ctx *route.Context, request *InfoRequest) error {
	info, err := userdao.UserInfoDao.GetUserSimpleInfo(request.Uid...)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	//goland:noinspection ALL
	resp := []InfoResponse{}
	for _, i := range info {
		resp = append(resp, InfoResponse{
			Uid:      i.Uid,
			Nickname: i.Nickname,
			// Account:  i.Account,
			Avatar: i.Avatar,
		})
	}
	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, resp))
	return nil
}

func (a *UserApi) GetOnlineUser(msg *route.Context) error {

	type u struct {
		Uid      int64
		Account  string
		Avatar   string
		Nickname string
	}
	//goland:noinspection GoPreferNilSlice
	allClient := []u{}
	users := make([]u, len(allClient))

	m := messages.NewMessage(msg.Seq, comm2.ActionSuccess, users)
	msg.Response(m)
	return nil
}

func (a *UserApi) UserProfile(ctx *route.Context) error {

	info, err := userdao.UserInfoDao.GetUserSimpleInfo(ctx.Uid)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	//goland:noinspection ALL
	resp := []InfoResponse{}
	for _, i := range info {
		resp = append(resp, InfoResponse{
			Uid:      i.Uid,
			Nickname: i.Nickname,
			Account:  i.Account,
			Avatar:   i.Avatar,
		})
	}
	if len(resp) != 1 {
		return comm2.NewUnexpectedErr("no such user", errors.New("user info is empty"))
	}
	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, resp[0]))
	return nil
}
