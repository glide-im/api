package cs

import (
	"github.com/glide-im/api/internal/api/router"
	"github.com/livekit/protocol/auth"
	"time"
)

type CsApi struct {
}

func (*CsApi) GetRecentChatMessage(ctx *route.Context) error {

	// TODO 2022-4-26
	ctx.ReturnSuccess(&Waiter{
		Uid:      0,
		Nickname: "CustomerService",
		Avatar:   "",
	})
	return nil
}

func (*CsApi) GetJoinToken(ctx *route.Context, request *RoomRequest) error {
	canPublish := true
	canSubscribe := true

	const apiKey = "devkey"
	const apiSecret = "secret"
	const room = "chat"

	at := auth.NewAccessToken(apiKey, apiSecret)
	grant := &auth.VideoGrant{
		RoomJoin:     true,
		Room:         room,
		CanPublish:   &canPublish,
		CanSubscribe: &canSubscribe,
	}
	at.AddGrant(grant).
		SetIdentity(request.Name).
		SetValidFor(time.Hour)

	sign, _ := at.ToJWT()

	ctx.ReturnSuccess(Data{Sign: sign})
	return nil
}
