package user

import (
	"github.com/glide-im/api/internal/dao/msgdao"
	"github.com/glide-im/api/internal/dao/wrapper/collect"
)

type InfoRequest struct {
	Uid []int64 `json:"uid"`
}

type InfoResponse struct {
	Uid         int64               `json:"uid"`
	Nickname    string              `json:"nick_name"`
	Account     string              `json:"account"`
	Avatar      string              `json:"avatar"`
	CategoryIds []int64             `json:"category_ids"`
	Collect     collect.CollectData `json:"collect"`
}

type InfoListResponse struct {
	UserInfo []*InfoResponse
}

type UpdateProfileRequest struct {
	Nickname string `validate:"required,gte=2,lte=16" json:"nick_name"`
	Password string `json:"password"`
	Avatar   string `validate:"required,url" json:"avatar"`
}

type UpdateEmailRequest struct {
	Email   string `json:"email" validate:"required|email"`
	Captcha string `json:"captcha" validate:"required"`
}

type ContactResponse struct {
	Id           int64               `json:"id"`
	Uid          int64               `json:"uid"`
	Type         int8                `json:"type"`
	Remark       string              `json:"remark"`
	Nickname     string              `json:"nickname"`
	Account      string              `json:"account"`
	Avatar       string              `json:"avatar"`
	LastMessage  msgdao.ChatMessage  `json:"lastMessage"`
	CategoryIds  []int64             `json:"categoryIds"`
	Collect      collect.CollectData `json:"collect"`
	MessageCount int64               `json:"message_count"`
}

type UserProfileResponse struct {
	AppID    int64  `json:"app_id"`
	Uid      int64  `json:"uid"`
	Account  string `json:"account"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Device   int64  `json:"device"`
}

type AddContacts struct {
	Uid    int64  `json:"uid"`
	Remark string `json:"remark"`
}

type DeleteContactsRequest struct {
	Uid int64 `json:"uid"`
}

type UpdateRemarkRequest struct {
	Uid    int64  `json:"uid"`
	Remark string `json:"remark"`
}

type UpdateLastMidRequest struct {
	Uid int64 `json:"uid"`
	Mid int64 `json:"mid"`
}

type ContactApproval struct {
	Uid     int64  `json:"uid"`
	Agree   bool   `json:"agree"`
	Comment string `json:"comment"`
}

type OnlineUser struct {
	Uid    int64 `json:"uid"`
	Before int64 `json:"before"`
}
