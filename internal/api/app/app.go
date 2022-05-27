package app

import (
	"fmt"
	"github.com/glide-im/api/internal/api/comm"
	"github.com/glide-im/api/internal/api/router"
	"github.com/glide-im/api/internal/dao/appdao"
	"github.com/glide-im/glide/pkg/messages"
)

type Interface interface {
	Echo(req *route.Context) error
}

type AppApi struct {
}

func (*AppApi) Echo(req *route.Context) error {
	req.Response(messages.NewMessage(req.Seq, "api.app.echo", fmt.Sprintf("seq=%d, uid=%d", req.Seq, req.Uid)))
	return nil
}

func (*AppApi) GetReleaseInfo(ctx *route.Context) error {

	info, err := appdao.Impl.GetReleaseInfo()
	if err != nil {
		return err
	}
	ctx.Response(messages.NewMessage(0, comm.ActionSuccess, info))
	return nil
}
