package user

import (
	"errors"
	"fmt"
	comm2 "github.com/glide-im/api/internal/api/comm"
	"github.com/glide-im/api/internal/api/router"
	"github.com/glide-im/api/internal/dao/userdao"
	"github.com/glide-im/api/internal/dao/wrapper/tm"
	"github.com/glide-im/glide/pkg/messages"
)

type UserApi struct{}

func (a *UserApi) GetUserProfile(msg *route.Context) error {
	// TODO 2021-11-29 我的详细信息
	return nil
}

func (a *UserApi) UpdateUserProfile(ctx *route.Context, request *UpdateProfileRequest) error {
	// TODO 2021-11-29 更新我的信息
	user := userdao.UpdateProfile{
		Nickname: request.Nickname,
		//Password: request.Password,
		Avatar: request.Avatar,
	}
	err := userdao.UserInfoDao.UpdateProfile(ctx.Uid, user)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, ""))
	return nil
}

func (a *UserApi) UpdateUserEmail(ctx *route.Context, request *UpdateEmailRequest) error {
	err := tm.VerifyCodeU.ValidateVerifyCode(request.Email, request.Captcha)
	if err != nil {
		return err
	}

	// 账户是否存在
	exist, err := userdao.UserInfoDao.AccountExists(request.Email, ctx.Uid)
	if err != nil {
		return err
	}
	fmt.Println("exist", exist)

	// 账户是否存在
	if exist {
		return errors.New("邮箱已被占用")
	}

	// TODO 2021-11-29 更新我的信息
	user := userdao.UpdateProfile{
		Email: request.Email,
	}
	err = userdao.UserInfoDao.UpdateProfile(ctx.Uid, user)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, ""))
	return nil
}

func (a *UserApi) GetUserInfo(ctx *route.Context, request *InfoRequest) error {
	//goland:noinspection ALL
	resp := []InfoResponse{}
	for _, u := range request.Uid {
		si, err := userdao.UserInfoDao.GetUserSimpleInfo(u)
		cateIds, err := userdao.UserInfoDao.GetUserCategory([]int64{u}, ctx.AppID)
		collect, err := userdao.UserInfoDao.GetCollectData(u, ctx.AppID)
		if err != nil {
			continue
		}
		for _, i := range si {
			resp = append(resp, InfoResponse{
				Uid:         i.Uid,
				Nickname:    i.Nickname,
				Account:     i.Account,
				Avatar:      i.Avatar,
				CategoryIds: cateIds,
				Collect:     collect,
			})
		}
		fmt.Println(resp)
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

func (a *UserApi) UserAuthProfile(ctx *route.Context) error {
	info, err := userdao.UserInfoDao.GetUserSimpleInfo(ctx.Uid)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	user := info[0]
	profile := UserProfileResponse{
		AppID:    ctx.AppID,
		Uid:      ctx.Uid,
		Account:  user.Account,
		Email:    user.Email,
		Phone:    user.Phone,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Device:   ctx.Device,
	}
	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, profile))
	return nil
}
