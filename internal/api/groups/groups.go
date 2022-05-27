package groups

import (
	comm2 "github.com/glide-im/api/internal/api/comm"
	"github.com/glide-im/api/internal/api/router"
	"github.com/glide-im/api/internal/dao/common"
	"github.com/glide-im/api/internal/dao/groupdao"
	"github.com/glide-im/api/internal/dao/msgdao"
	"github.com/glide-im/api/internal/dao/userdao"
	"github.com/glide-im/api/internal/group"
	"github.com/glide-im/api/internal/im"
	"github.com/glide-im/glide/pkg/logger"
	"github.com/glide-im/glide/pkg/messages"
)

type Interface interface {
}

type GroupApi struct {
}

func (m *GroupApi) CreateGroup(ctx *route.Context, request *CreateGroupRequest) error {

	dbGroup, err := groupdao.Dao.CreateGroup(request.Name, 1)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	_, err = msgdao.GroupMsgDaoImpl.CreateGroupMessageState(dbGroup.Gid)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	err = userdao.Dao.AddContacts(ctx.Uid, dbGroup.Gid, userdao.ContactsTypeGroup)
	if err != nil {
		return comm2.NewDbErr(err)
	}

	err = groupdao.Dao.AddMember(dbGroup.Gid, ctx.Uid, groupdao.GroupMemberTypeOwner, groupdao.GroupFlagDefault)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	//err = groupdao.Dao.AddMembers(dbGroup.Gid, MemberFlagDefault, MemberTypeNormal, request.Member...)
	//if err != nil {
	//	return comm.NewDbErr(err)
	//}
	err = group.Interface.CreateGroup(dbGroup.Gid)
	if err != nil {
		return comm2.NewUnexpectedErr("create group failed", err)
	}
	err = group.Interface.PutMember(dbGroup.Gid, []int64{ctx.Uid})
	if err != nil {
		return comm2.NewUnexpectedErr("add group member failed", err)
	}
	err = group.Interface.UpdateMember(dbGroup.Gid, ctx.Uid, 1)
	if err != nil {
		return comm2.NewUnexpectedErr("create group failed", err)
	}
	//n := messages.NewMessage(0, comm.ActionInviteToGroup, InviteGroupMessage{Gid: dbGroup.Gid})
	//for _, uid := range request.Member {
	//	apidep.SendMessageIfOnline(uid, 0, n)
	//}
	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, CreateGroupResponse{Gid: dbGroup.Gid}))
	return nil
}

func (m *GroupApi) GetGroupMember(ctx *route.Context, request *GetGroupMemberRequest) error {

	mbs, err := groupdao.Dao.GetMembers(request.Gid)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	ms := make([]*GroupMemberResponse, 0, len(mbs))
	for _, member := range mbs {
		ms = append(ms, &GroupMemberResponse{
			Uid:        member.Uid,
			RemarkName: member.Remark,
			Type:       int(member.Type),
		})
	}
	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, ms))
	return nil
}

func (m *GroupApi) GetGroupInfo(ctx *route.Context, request *GroupInfoRequest) error {
	groups, err := groupdao.Dao.GetGroups(request.Gid...)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, groups))
	return nil
}

func (m *GroupApi) RemoveMember(ctx *route.Context, request *RemoveMemberRequest) error {
	typ, err := groupdao.Dao.GetMemberType(request.Gid, ctx.Uid)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	if typ == groupdao.GroupMemberTypeAdmin || typ == groupdao.GroupMemberTypeOwner {
		//goland:noinspection GoPreferNilSlice
		notFind := []int64{}
		for _, id := range request.Uid {
			err = userdao.ContactsDao.DelContacts(id, request.Gid, userdao.ContactsTypeGroup)
			if err != nil {
				return comm2.NewDbErr(err)
			}
			err = groupdao.Dao.RemoveMember(request.Gid, id)
			if err == common.ErrNoRecordFound {
				notFind = append(notFind, id)
				continue
			} else if err != nil {
				return comm2.NewDbErr(err)
			}
			err = dispatchGroupNotify(request.Gid, 1, id)
			if err != nil {
				logger.E("remove member error:%v", err)
			}
		}
		if len(notFind) == 0 {
			ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, ""))
		} else {
			ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionFailed, notFind))
		}
	} else {
		return ErrGroupNotExit
	}
	return nil
}

func (m *GroupApi) AddGroupMember(ctx *route.Context, request *AddMemberRequest) error {
	for _, uid := range request.Uid {
		err := addGroupMemberDb(request.Gid, uid, groupdao.GroupMemberNormal)
		if err != nil {
			return err
		}
		err = userdao.ContactsDao.AddContacts(uid, request.Gid, userdao.ContactsTypeGroup)
		if err != nil {
			return comm2.NewDbErr(err)
		}
		n := messages.NewMessage(0, "", comm2.NewContactMessage{
			FromId:   ctx.Uid,
			FromType: 0,
			Id:       request.Gid,
			Type:     userdao.ContactsTypeGroup,
		})
		im.SendMessage(uid, 0, n)
		err = group.Interface.PutMember(request.Gid, []int64{uid})
		if err != nil {
			return comm2.NewUnexpectedErr("add group failed", err)
		}
		err = dispatchGroupNotify(request.Gid, 1, uid)
		if err != nil {
			logger.E("notify add group member error: %v", err)
		}
	}
	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, ""))
	return nil
}

func (m *GroupApi) ExitGroup(ctx *route.Context, request *ExitGroupRequest) error {

	err := group.Interface.RemoveMember(request.Gid, ctx.Uid)
	if err != nil {
		return comm2.NewUnexpectedErr("exit group failed", err)
	}
	err = groupdao.Dao.RemoveMember(request.Gid, ctx.Uid)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	err = dispatchGroupNotify(request.Gid, 1, ctx.Uid)
	if err != nil {
		logger.E("exit group error: %v", err)
	}
	resp := messages.NewMessage(ctx.Seq, comm2.ActionSuccess, " group success")
	ctx.Response(resp)
	return err
}

func (m *GroupApi) JoinGroup(ctx *route.Context, request *JoinGroupRequest) error {

	isC, err := userdao.ContactsDao.HasContacts(ctx.Uid, request.Gid, userdao.ContactsTypeGroup)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	if isC {
		return ErrMemberAlreadyExist
	}
	// TODO 2021-11-29 use transaction
	err = userdao.ContactsDao.AddContacts(ctx.Uid, request.Gid, userdao.ContactsTypeGroup)
	if err != nil {
		return comm2.NewDbErr(err)
	}

	err = addGroupMemberDb(request.Gid, ctx.Uid, groupdao.GroupMemberNormal)
	if err != nil {
		return err
	}
	err = group.Interface.PutMember(request.Gid, []int64{ctx.Uid})
	if err != nil {
		return comm2.NewUnexpectedErr("add group failed", err)
	}
	err = dispatchGroupNotify(request.Gid, 1, ctx.Uid)
	if err != nil {
		logger.E("join group error:%v", err)
	}
	ctx.Response(messages.NewMessage(ctx.Seq, comm2.ActionSuccess, ""))
	return nil
}

func addGroupMemberDb(gid int64, uid int64, typ int64) error {
	hasGroup, err := groupdao.Dao.HasGroup(gid)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	if !hasGroup {
		return ErrGroupNotExit
	}
	hasMember, err := groupdao.Dao.HasMember(gid, uid)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	if hasMember {
		return ErrMemberAlreadyExist
	}
	err = groupdao.Dao.AddMember(gid, uid, typ, groupdao.GroupFlagDefault)
	if err != nil {
		return comm2.NewDbErr(err)
	}
	return nil
}

func dispatchGroupNotify(gid int64, typ int64, uid int64) error {
	//id, err := msgdao.GetMessageID()
	//if err != nil {
	//	logger.E("get message id error:%v", err)
	//	return err
	//}
	//n := message.NewGroupNotifyAdded([]int64{uid})
	//notify := message.NewGroupNotify(id, gid, 0, typ, time.Now().Unix(), &n)
	//return apidep.Interface.DispatchNotifyMessage(gid, notify)
	return nil
}
