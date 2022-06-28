package api

import (
	"github.com/glide-im/api/internal/api/app"
	"github.com/glide-im/api/internal/api/auth"
	"github.com/glide-im/api/internal/api/cs"
	"github.com/glide-im/api/internal/api/groups"
	"github.com/glide-im/api/internal/api/msg"
	"github.com/glide-im/api/internal/api/user"
	"github.com/glide-im/api/internal/api/wrapper/articles"
)

func initRoute() {

	appApi := app.AppApi{}
	getNoAuth("api/app/release", appApi.GetReleaseInfo)

	authApi := auth.AuthApi{}
	postNoAuth("/api/auth/register", authApi.Register)
	postNoAuth("/api/auth/guest", authApi.GuestRegister)
	postNoAuth("/api/auth/signin", authApi.SignIn)
	postNoAuth("/api/auth/token", authApi.AuthToken)
	post("/api/auth/logout", authApi.Logout)

	groupApi := groups.GroupApi{}
	post("/api/group/info", groupApi.GetGroupInfo)
	post("/api/group/members", groupApi.GetGroupMember)
	post("/api/group/create", groupApi.CreateGroup)
	post("/api/group/join", groupApi.JoinGroup)
	post("/api/group/members/invite", groupApi.AddGroupMember)
	post("/api/group/members/remove", groupApi.RemoveMember)

	userApi := user.UserApi{}
	post("/api/contacts/add", userApi.AddContact)
	post("/api/contacts/list", userApi.GetContactList)
	post("/api/contacts/approval", userApi.ContactApproval)
	post("/api/contacts/del", userApi.DeleteContact)
	post("/api/contacts/update/remark", userApi.UpdateContactRemark)
	post("/api/contacts/update/mid", userApi.UpdateContactLastMid)

	post("/api/user/info", userApi.GetUserInfo)
	post("/api/user/profile", userApi.UserProfile)
	post("/api/user/profile/update", userApi.UpdateUserProfile)

	toolApi := cs.ToolApi{}
	post("/api/tool/get-qiniu-token", toolApi.GetQiniuToken)

	articleApi := articles.ArticleApi{}
	post("/api/articles/store", articleApi.Store)
	post("/api/articles/:id", articleApi.Update)
	_delete("/api/articles/delete/:id", articleApi.Delete)
	post("/api/articles/order", articleApi.Order)
	get("/api/articles/list", articleApi.List)

	msgApi := msg.MsgApi{}
	post("/api/msg/id", msgApi.GetMessageID)
	post("/api/msg/group/history", msgApi.GetGroupMessageHistory)
	post("/api/msg/group/recent", msgApi.GetRecentGroupMessage)
	post("/api/msg/group/state", msgApi.GetGroupMessageState)
	post("/api/msg/group/state/all", msgApi.GetUserGroupMessageState)

	post("/api/msg/chat/history", msgApi.GetChatMessageHistory)
	post("/api/msg/chat/user", msgApi.GetRecentMessageByUser)
	post("/api/msg/chat/recent", msgApi.GetRecentMessage)
	post("/api/msg/chat/offline", msgApi.GetOfflineMessage)
	post("/api/msg/chat/offline/ack", msgApi.AckOfflineMessage)

	post("/api/session/recent", msgApi.GetRecentSessions)
	post("/api/session/get", msgApi.GetOrCreateSession)

	csApi := cs.CsApi{}
	post("/api/cs/get", csApi.GetRecentChatMessage)
}

func postNoAuth(path string, fn interface{}) {
	rt.POST(path, getHandler(path, fn))
}
func getNoAuth(path string, fn interface{}) {
	rt.GET(path, getHandler(path, fn))
}

func post(path string, fn interface{}) {
	useAuth().POST(path, getHandler(path, fn))
}

func _delete(path string, fn interface{}) {
	useAuth().DELETE(path, getHandler(path, fn))
}

func get(path string, fn interface{}) {
	useAuth().GET(path, getHandler(path, fn))
}
