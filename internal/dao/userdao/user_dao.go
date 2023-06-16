package userdao

import (
	"github.com/glide-im/api/internal/dao/common"
	"github.com/glide-im/api/internal/dao/wrapper/app"
	"github.com/glide-im/api/internal/dao/wrapper/category"
	"github.com/glide-im/api/internal/dao/wrapper/collect"
	"github.com/glide-im/api/internal/pkg/db"
	"time"
)

const (
	ContactsTypeUser  = 1
	ContactsTypeGroup = 2
)

const (
	ContactsStatusKeep     = 1 // 正常关系
	ContactsStatusApproval = 2 // 等待同意
	ContactsStatusInvalid  = 3 // 双方不存在好友关系
)

var UserInfoDao = &UserInfoDaoImpl{}

type UserInfoDaoImpl struct{}

func (d *UserInfoDaoImpl) AddUser(u *User) error {
	u.Uid = 0
	u.CreateAt = time.Now().Unix()
	query := db.DB.Create(&u)
	return common.ResolveError(query)
}

func (d *UserInfoDaoImpl) AddGuestUser(u *User) error {
	uid := d.GetLastUserId()
	uid += 1
	u.Uid = uid
	u.CreateAt = time.Now().Unix()
	query := db.DB.Create(u)
	return common.ResolveError(query)
}

func (d *UserInfoDaoImpl) DelUser(uid int64) error {
	query := db.DB.Where("uid = ?", uid).Delete(&User{})
	return common.ResolveError(query)
}

func (d *UserInfoDaoImpl) AccountExists(email string, excludeIds ...int64) (bool, error) {
	var count int64
	query := db.DB.Model(&User{})
	if len(excludeIds) > 0 {
		query.Where("uid not in (?)", excludeIds)
	}
	query.Where("email = ?", email).Count(&count)
	if err := common.ResolveError(query); err != nil {
		return false, err
	}
	return count > 0, nil
}

func (d *UserInfoDaoImpl) HasUser(uid int64, selfId int64) (bool, error) {
	var count int64
	query := db.DB.Model(&User{}).Where("uid = ? and app_id = ?", uid, selfId).Count(&count)
	if err := common.ResolveError(query); err != nil {
		return false, err
	}
	return count > 0, nil
}

func (d *UserInfoDaoImpl) UpdateNickname(uid int64, nickname string) error {
	return d.update(uid, "nickname", nickname)
}

func (d *UserInfoDaoImpl) UpdateSecret(uid int64, secret string) error {
	return d.update(uid, "message_deliver_secret", secret)
}

func (d *UserInfoDaoImpl) UpdateAvatar(uid int64, avatar string) error {
	return d.update(uid, "avatar", avatar)
}

func (d *UserInfoDaoImpl) UpdateAppId(uid int64, appId int64) error {
	return d.update(uid, "app_id", appId)
}

func (d *UserInfoDaoImpl) UpdatePassword(uid int64, password string) error {
	return d.update(uid, "password", password)
}

func (d *UserInfoDaoImpl) GetUidInfoByLogin(account string, password string) (User, error) {
	var user User
	query := db.DB.Model(&User{}).
		Where("email = ?", account).
		Select("uid, nickname, email, password").
		Find(&user)
	if query.Error != nil {
		return user, query.Error
	}
	if query.RowsAffected == 0 {
		return user, common.ErrNoRecordFound
	}
	if !PasswordVerify(password, user.Password) {
		return user, common.ErrNoRecordFound
	}
	return user, nil
}

func (d *UserInfoDaoImpl) GetPassword(uid int64) (string, error) {
	var password string
	query := db.DB.Model(&User{}).Where("uid = ?", uid).Select("password").Find(&password)
	if err := common.ResolveError(query); err != nil {
		return "", err
	}
	return password, nil
}

func (d *UserInfoDaoImpl) GetUser(uid int64) (*User, error) {
	user := &User{}
	query := db.DB.Model(user).Where("uid = ?", uid).Find(user)
	if err := common.ResolveError(query); err != nil {
		return nil, err
	}
	return user, nil
}

func (d *UserInfoDaoImpl) GetUserSimpleInfo(uid ...int64) ([]*User, error) {
	var us []*User
	query := db.DB.Model(&User{}).Where("uid IN (?)", uid).Select("uid, account, nickname, avatar, email, role").Find(&us)
	if err := common.MustFind(query); err != nil {
		return nil, err
	}
	return us, nil
}

func (d *UserInfoDaoImpl) GetUserSimpleOneInfo(uid int64) (User, error) {
	var us User
	query := db.DB.Model(&User{}).Where("uid IN (?)", uid).Select("uid, account, nickname, avatar, email, role").Find(&us)
	if err := common.MustFind(query); err != nil {
		return us, err
	}
	return us, nil
}

func (d *UserInfoDaoImpl) update(uid int64, field string, value interface{}) error {
	query := db.DB.Model(&User{}).Where("uid = ?", uid).Update(field, value)
	return common.ResolveError(query)
}

func (d *UserInfoDaoImpl) GetUserCategory(uids []int64, app_id int64) ([]int64, error) {
	var categoryUsers []category.CategoryUser
	query := db.DB.Model(category.CategoryUser{}).Where("app_id = ?", app_id).Where("uid in ?", uids).Select("uid, category_id").Find(&categoryUsers)
	if err := common.JustError(query); err != nil {
		return nil, err
	}
	var category_ids []int64
	for _, cate := range categoryUsers {
		category_ids = append(category_ids, cate.CategoryId)
	}
	return category_ids, nil
}

func (d *UserInfoDaoImpl) GetCollectData(uid int64, app_id int64) (collect.CollectData, error) {
	var collectData collect.CollectData
	query := db.DB.Model(collect.CollectData{}).Where("app_id = ?", app_id).Where("uid = ?", uid).Find(&collectData)
	if err := common.JustError(query); err != nil {
		return collectData, err
	}

	return collectData, nil
}

func (d *UserInfoDaoImpl) UpdateProfile(uid int64, profile UpdateProfile) error {
	query := db.DB.Model(&User{}).Where("uid = ?", uid).Updates(User{
		Avatar:   profile.Avatar,
		Nickname: profile.Nickname,
		//Password: profile.Password,
		//Email:    profile.Email,
	})
	return common.JustError(query)
}

func (d *UserInfoDaoImpl) GetUserAppId(uid int64) int64 {
	var appModel app.App
	db.DB.Model(&app.App{}).Where("uid = ?", uid).Find(&appModel)
	return appModel.Id
}

func (d *UserInfoDaoImpl) GetGuestUserAppId(uid int64) int64 {
	var user User
	db.DB.Model(&User{}).Where("uid = ?", uid).Find(&user)
	return user.AppID
}

func (d *UserInfoDaoImpl) GetLastUserId() int64 {
	var user User
	db.DB.Model(&User{}).Order("uid desc").Select("uid").Last(&user)
	return user.Uid
}
