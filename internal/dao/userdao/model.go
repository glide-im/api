package userdao

import "github.com/glide-im/api/internal/dao/wrapper/category"

type User struct {
	Uid      int64  `gorm:"primaryKey"`
	Account  string `gorm:"unique"`
	Nickname string
	Password string
	Avatar   string
	CreateAt int64

	CategoryUser []category.CategoryUser
}

type Contacts struct {
	Fid     string `gorm:"primaryKey"`
	Uid     int64
	Id      int64
	Remark  string
	Type    int8
	Status  int8
	LastMid int64
}

type LoginState struct {
	Device int64
	Token  string
}
