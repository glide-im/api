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
	// TODO 2021-11-29
	return nil
}

func (a *UserApi) UpdateContactRemark(ctx *route.Context, request *UpdateRemarkRequest) error {
	// TODO 2021-11-29
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
	hasUser, err := userdao.UserInfoDao.HasUser(request.Uid)
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

	resp := []ContactResponse{}
	for _, contact := range contacts {
		if contact.Type == userdao.ContactsTypeGroup {
			// TODO 2022-4-24 member flag
			_ = group.Interface.UpdateMember(contact.Id, ctx.Uid, 1)
		}
		resp = append(resp, ContactResponse{
			Id:     contact.Id,
			Type:   contact.Type,
			Remark: contact.Remark,
		})
	}

	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, resp))
	return nil
}
