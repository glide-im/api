package user

import (
	comm2 "github.com/glide-im/api/internal/api/comm"
	route "github.com/glide-im/api/internal/api/router"
	"github.com/glide-im/api/internal/dao/msgdao"
	"github.com/glide-im/api/internal/dao/userdao"
	"github.com/glide-im/api/internal/group"
	"github.com/glide-im/api/internal/im"
	"github.com/glide-im/glide/pkg/logger"
	"github.com/glide-im/glide/pkg/messages"
	"time"
)

func (a *UserApi) DeleteContact(ctx *route.Context, request *DeleteContactsRequest) error {
	err := userdao.ContactsDao.DelContacts(ctx.Uid, request.Uid, userdao.ContactsTypeUser)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, ""))
	return nil
}

func (a *UserApi) UpdateContactRemark(ctx *route.Context, request *UpdateRemarkRequest) error {
	err := userdao.ContactsDao.UpdateContactRemark(ctx.Uid, request.Uid, userdao.ContactsTypeUser, request.Remark)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, ""))
	return nil
}

func (a *UserApi) UpdateContactLastMid(ctx *route.Context, request *UpdateLastMidRequest) error {
	err := userdao.ContactsDao.UpdateContactLastMid(ctx.Uid, request.Uid, userdao.ContactsTypeUser, request.Mid)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, ""))
	return nil
}

func (a *UserApi) ContactApproval(ctx *route.Context, request *ContactApproval) error {
	// TODO 2021-11-29
	return nil
}

func (a *UserApi) AddContact(ctx *route.Context, request *AddContacts) error {

	if ctx.Uid == request.Uid {
		return errAddSelf
	}
	hasUser, err := userdao.UserInfoDao.HasUser(request.Uid, ctx.AppID)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	if !hasUser {
		return errUserNotExist
	}

	isC, err := userdao.ContactsDao.HasContacts(ctx.Uid, request.Uid, 1)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	if isC {
		return errAlreadyContacts
	}
	// TODO 2021-11-29 use transaction
	err = userdao.ContactsDao.AddContacts(ctx.Uid, request.Uid, userdao.ContactsTypeUser)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	err = userdao.ContactsDao.AddContacts(request.Uid, ctx.Uid, userdao.ContactsTypeUser)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	_, err = msgdao.SessionDaoImpl.CreateSession(ctx.Uid, request.Uid, time.Now().Unix())
	if err != nil {
		logger.E("create session error %v", err)
	}

	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, ""))

	m := comm2.NewContactMessage{
		FromId:   ctx.Uid,
		FromType: 0,
		Id:       ctx.Uid,
		Type:     userdao.ContactsTypeUser,
	}
	im.SendMessage(request.Uid, 0, messages.NewMessage(-1, "", m))
	return nil
}

//goland:noinspection GoPreferNilSlice
func (a *UserApi) GetContactList(ctx *route.Context) error {

	contacts, err := userdao.ContactsDao.GetContacts(ctx.Uid)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	selfContact := userdao.Contacts{
		Fid:     "",
		Uid:     ctx.Uid,
		Id:      ctx.Uid,
		Remark:  "",
		Type:    0,
		Status:  0,
		LastMid: 0,
	}
	resp := []ContactResponse{}
	contacts = append(contacts, &selfContact)
	for _, contact := range contacts {
		if contact.Type == userdao.ContactsTypeGroup {
			// TODO 2022-4-24 member flag
			_ = group.Interface.UpdateMember(contact.Id, ctx.Uid, 1)
		}
		user, _ := userdao.UserInfoDao.GetUserSimpleOneInfo(contact.Id)
		cateIds, _ := userdao.UserInfoDao.GetUserCategory([]int64{contact.Id}, ctx.AppID)
		collect, _ := userdao.UserInfoDao.GetCollectData(contact.Id, ctx.AppID)

		nickname := user.Nickname
		if ctx.Uid == contact.Id {
			user.Nickname = nickname + "(自己)"
		}

		resp = append(resp, ContactResponse{
			Id:           contact.Id,
			Uid:          contact.Id,
			Type:         contact.Type,
			Remark:       contact.Remark,
			Nickname:     user.Nickname,
			Account:      user.Account,
			Avatar:       user.Avatar,
			CategoryIds:  cateIds,
			Collect:      collect,
			LastMessage:  msgdao.ChatMsgDaoImpl.GetChatLastMessage(ctx.Uid, contact.Id),
			MessageCount: msgdao.ChatMsgDaoImpl.GetMessageCount(ctx.Uid, contact.Id),
		})
	}

	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, resp))
	return nil
}
