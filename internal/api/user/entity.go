package user

import (
	"github.com/glide-im/api/internal/dao/msgdao"
	"github.com/glide-im/api/internal/dao/wrapper/collect"
)

type InfoRequest struct {
	Uid []int64
}

type InfoResponse struct {
	Uid         int64
	Nickname    string
	Account     string
	Avatar      string
	CategoryIds []int64
	Collect     collect.CollectData
}

type InfoListResponse struct {
	UserInfo []*InfoResponse
}

type UpdateProfileRequest struct {
	Nickname string `validate:"required,gte=2,lte=16"`
	Password string
	Avatar   string `validate:"required,url"`
}

type UpdateEmailRequest struct {
	Email   string `json:"email" validate:"required|email"`
	Captcha string `json:"captcha" validate:"required"`
}

type ContactResponse struct {
	Id          int64               `json:"id"`
	Uid         int64               `json:"uid"`
	Type        int8                `json:"type"`
	Remark      string              `json:"remark"`
	Nickname    string              `json:"nickname"`
	Account     string              `json:"account"`
	Avatar      string              `json:"avatar"`
	LastMessage msgdao.ChatMessage  `json:"lastMessage"`
	CategoryIds []int64             `json:"categoryIds"`
	Collect     collect.CollectData `json:"collect"`
}

type AddContacts struct {
	Uid    int64 `json:"uid"`
	Remark string
}

type DeleteContactsRequest struct {
	Uid int64
}

type UpdateRemarkRequest struct {
	Uid    int64
	Remark string
}

type UpdateLastMidRequest struct {
	Uid int64
	Mid int64
}

type ContactApproval struct {
	Uid     int64
	Agree   bool
	Comment string
}

type OnlineUser struct {
	Uid    int64
	Before int64
}
