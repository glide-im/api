package group

// Interface 群相关接口
var Interface Service = &groupInterface{}

type Service interface {
	PutMember(gid int64, mb []int64) error
	RemoveMember(gid int64, uid ...int64) error
	CreateGroup(gid int64) error
	UpdateMember(gid int64, uid int64, flag int64) error
}

type groupInterface struct {
}

func (g groupInterface) PutMember(gid int64, mb []int64) error {
	//TODO implement me
	panic("implement me")
}

func (g groupInterface) RemoveMember(gid int64, uid ...int64) error {
	//TODO implement me
	panic("implement me")
}

func (g groupInterface) CreateGroup(gid int64) error {
	//TODO implement me
	panic("implement me")
}

func (g groupInterface) UpdateMember(gid int64, uid int64, flag int64) error {
	//TODO implement me
	panic("implement me")
}
