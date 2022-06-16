package im

import (
	"github.com/glide-im/glide/pkg/gate"
	"github.com/glide-im/glide/pkg/logger"
	"github.com/glide-im/glide/pkg/messages"
	"github.com/glide-im/glide/pkg/rpc"
	"github.com/glide-im/im-service/pkg/client"
	"strconv"
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
	opt := &rpc.ClientOptions{
		Addr:        addr,
		Port:        port,
		Name:        name,
		EtcdServers: nil,
		Selector:    nil,
	}
	cli, err := client.NewIMServiceClient(opt)
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
	err = c.cli.SetClientID(gate.NewID("", strconv.FormatInt(uid, 10), strconv.FormatInt(device, 10)), id)
	if err != nil {
		return err
	}
	return c.cli.ExitClient(gate.NewID("", strconv.FormatInt(uid, 10), strconv.FormatInt(device, 10)))
}

func (c imServiceRpcClient) EnqueueMessage(uid int64, device int64, message *messages.GlideMessage) error {
	c.cli.ExitClient(gate.NewID2("1"))
	return c.cli.EnqueueMessage(gate.NewID("", strconv.FormatInt(uid, 10), strconv.FormatInt(device, 10)), message)
}
