package userdao

import (
	"github.com/glide-im/api/internal/dao/wrapper/collect"
	"time"
)

var Dao = UserDao{
	Cache:                UserCacheDao{},
	UserInfoDaoInterface: UserInfoDao,
	ContactsDaoInterface: ContactsDao,
}

type Cache interface {
	//GetUserSignState(uid int64) ([]*LoginState, error)
	//IsUserSignIn(uid int64, device int64) (bool, error)
	//DelToken(token string) error
	//DelAllToken(uid int64) error
	//GetTokenInfo(token string) (int64, int64, error)
	//SetSignInToken(uid int64, device int64, token string, expiredAt time.Duration) error

	DelAuthToken(uid int64, device int64) error
	SetTokenVersion(uid int64, device int64, version int64, expiredAt time.Duration) error
	GetTokenVersion(uid int64, device int64) (int64, error)
}

type UserInfoDaoInterface interface {
	AddUser(u *User) error
	DelUser(uid int64) error
	HasUser(uid int64, appId int64) (bool, error)
	AccountExists(account string, excludeIds ...int64) (bool, error)

	UpdateNickname(uid int64, nickname string) error
	UpdateAvatar(uid int64, avatar string) error
	UpdatePassword(uid int64, password string) error
	GetPassword(uid int64) (string, error)

	GetUidInfoByLogin(email string, password string) (User, error)
	GetUser(uid int64) (*User, error)
	GetUserSimpleInfo(uid ...int64) ([]*User, error)
	GetUserCategory(uid []int64, app_od int64) ([]int64, error)
	GetCollectData(uid int64, app_id int64) (collect.CollectData, error)
	GetUserAppId(uid int64) int64
	GetGuestUserAppId(uid int64) int64
	GetUserSimpleOneInfo(uid int64) (User, error)
	GetLastUserId() int64
}

type ContactsDaoInterface interface {
	HasContacts(uid int64, id int64, type_ int8) (bool, error)
	AddContacts(uid int64, id int64, type_ int8) error
	DelContacts(uid int64, id int64, type_ int8) error
	GetContacts(uid int64) ([]*Contacts, error)
	GetContactsByType(uid int64, type_ int) ([]*Contacts, error)
}

type UserDao struct {
	Cache
	UserInfoDaoInterface
	ContactsDaoInterface
}
