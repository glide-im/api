package im

import (
	"github.com/glide-im/glide/pkg/gate"
	"github.com/glide-im/glide/pkg/logger"
	"github.com/glide-im/glide/pkg/messages"
	"github.com/glide-im/im-service/pkg/client"
	"github.com/glide-im/im-service/pkg/rpc"
)

// ClientInterface 客户端连接相关接口
var ClientInterface ClientManagerInterface

func SendMessage(uid int64, device int64, m *messages.GlideMessage) {
	err := ClientInterface.EnqueueMessage(uid, device, m)
	if err != nil {
		logger.E("SendMessage error: %v", err)
	}
}

type ClientManagerInterface interface {
	Logout(uid int64, device int64) error
	EnqueueMessage(uid int64, device int64, message *messages.GlideMessage) error
}

func MustSetupClient(addr string, port int, name string) {
	options := &rpc.ClientOptions{
		Addr: addr,
		Port: port,
		Name: name,
	}
	cli, err := client.NewIMServiceClient(options)
	if err != nil {
		panic(err)
	}
	ClientInterface = &imServiceRpcClient{cli}
}

type imServiceRpcClient struct {
	cli *client.IMServiceClient
}

func (c imServiceRpcClient) Logout(uid int64, device int64) error {
	id, err := gate.GenTempID("")
	if err != nil {
		return err
	}
	err = c.cli.SetClientID(gate.NewID("", uid, device), id)
	if err != nil {
		return err
	}
	return c.cli.ExitClient(gate.NewID("", uid, device))
}

func (c imServiceRpcClient) EnqueueMessage(uid int64, device int64, message *messages.GlideMessage) error {
	return c.cli.EnqueueMessage(gate.NewID("", uid, device), message)
}
