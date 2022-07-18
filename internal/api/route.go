package api

import (
	"github.com/glide-im/api/internal/api/app"
	"github.com/glide-im/api/internal/api/auth"
	"github.com/glide-im/api/internal/api/cs"
	"github.com/glide-im/api/internal/api/groups"
	"github.com/glide-im/api/internal/api/msg"
	"github.com/glide-im/api/internal/api/user"
	platformApp "github.com/glide-im/api/internal/api/wrapper/app"
	"github.com/glide-im/api/internal/api/wrapper/articles"
	"github.com/glide-im/api/internal/api/wrapper/category"
)

func initRoute() {

	appApi := app.AppApi{}
	authApi := auth.AuthApi{}
	groupApi := groups.GroupApi{}
	userApi := user.UserApi{}
	toolApi := cs.ToolApi{}
	articleApi := articles.ArticleApi{}
	platformAppApi := platformApp.PlatFromApi{}

	CategoryApi := category.CategoryApi{}
	msgApi := msg.MsgApi{}
	csApi := cs.CsApi{}
	appGuestApi := auth.AuthApi{}

	routes := group("api",
		group("app",
			get("release", appApi.GetReleaseInfo),
		),
		group("auth",
			post("register", authApi.Register),
			post("forget", authApi.Forget),
			post("guest", authApi.GuestRegister),
			post("signin", authApi.SignIn),
			post("token", authApi.AuthToken),
			post("verifyCode", authApi.VerifyCode),
			post("logout", authApi.Logout),
			group("guest",
				post("/signin", appGuestApi.GuestRegisterV2),
			),
		),
		group("guest",
			use(guestMiddleware,
				post("/guest-id", platformAppApi.GetGuestToId),
				get("/articles/:id", articleApi.GuestShow),
				get("/articles/list", articleApi.GuestList),
			),
		),
		use(authMiddleware,
			group("group",
				post("info", groupApi.GetGroupInfo),
				post("members", groupApi.GetGroupMember),
				post("create", groupApi.CreateGroup),
				post("join", groupApi.JoinGroup),
				post("members/invite", groupApi.AddGroupMember),
				post("members/remove", groupApi.RemoveMember),
			),
			group("contacts",
				post("add", userApi.AddContact),
				post("list", userApi.GetContactList),
				post("approval", userApi.ContactApproval),
				post("del", userApi.DeleteContact),
				post("update/remark", userApi.UpdateContactRemark),
				post("update/mid", userApi.UpdateContactLastMid),
			),
			group("user",
				post("info", userApi.GetUserInfo),
				post("profile", userApi.UserProfile),
				post("profile/update", userApi.UpdateUserProfile),
				post("profile/email", userApi.UpdateUserEmail),
				post("profile/auth", userApi.UserAuthProfile),
			),
			group("tool",
				post("get-qiniu-token", toolApi.GetQiniuToken),
			),
			group("articles",
				post("store", articleApi.Store),
				post(":id", articleApi.Update),
				delete_("delete/:id", articleApi.Delete),
				get("show/:id", articleApi.Show),
				post("order", articleApi.Order),
				get("list", articleApi.GuestList),
			),
			group("category",
				post("store", CategoryApi.Store),
				post("updates", CategoryApi.Updates),
				post(":id", CategoryApi.Update),
				delete_("delete/:id", CategoryApi.Delete),
				post("order", CategoryApi.Order),
				get("list", CategoryApi.List),
				post("user/:uid", CategoryApi.SetUserCategory),
			),
			group("app",
				post("store", platformAppApi.Store),
				post(":id", platformAppApi.Update),
				delete_("delete/:id", platformAppApi.Delete),
				get("list", platformAppApi.List),
				post("host", platformAppApi.UpdateHost),
			),
			group("msg",
				post("id", msgApi.GetMessageID),
				post("group/history", msgApi.GetGroupMessageHistory),
				post("group/recent", msgApi.GetRecentGroupMessage),
				post("group/state", msgApi.GetGroupMessageState),
				post("group/state/all", msgApi.GetUserGroupMessageState),
				post("chat/history", msgApi.GetChatMessageHistory),
				post("chat/user", msgApi.GetRecentMessageByUser),
				post("chat/recent", msgApi.GetRecentMessage),
				post("chat/offline", msgApi.GetOfflineMessage),
				post("chat/offline/ack", msgApi.AckOfflineMessage),
			),
			group("session",
				post("recent", msgApi.GetRecentSessions),
				post("get", msgApi.GetOrCreateSession),
			),
			group("cs",
				post("get", csApi.GetRecentChatMessage),
			),
		))
	routes.setup(g)
}
