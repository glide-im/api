package user

import (
	"github.com/glide-im/api/internal/api/router"
	"github.com/glide-im/api/internal/apidep"
	"github.com/glide-im/api/internal/im"
	"github.com/glide-im/api/internal/pkg/db"
	"github.com/glide-im/glide/pkg/logger"
	"testing"
)

var api = UserApi{}

func init() {
	db.Init()
	im.ClientInterface = apidep.MockClientManager{}
}

func getContext(uid int64, device int64) *route.Context {
	return &route.Context{
		Uid:    uid,
		Device: device,
		Seq:    1,
		Action: "",
		R: func(message *messages.GlideMessage) {
			logger.D("Response=%v", message)
		},
	}
}

func TestUserApi_AddContact(t *testing.T) {
	err := api.AddContact(getContext(543603, 1), &AddContacts{
		Uid: 543602,
	})
	if err != nil {
		t.Error(err)
	}
}

func TestUserApi_GetContactList(t *testing.T) {
	err := api.GetContactList(getContext(543603, 1))
	if err != nil {
		t.Error(err)
	}
}
