package im

import (
	"github.com/glide-im/glide/pkg/gate"
	"github.com/glide-im/glide/pkg/logger"
	"github.com/glide-im/glide/pkg/messages"
	"github.com/glide-im/glide/pkg/rpc"
	"github.com/glide-im/im-service/pkg/client"
	"strconv"
)

var IM Service

// SendMessage 发送消息到指定用户
func SendMessage(uid int64, device int64, m *messages.GlideMessage) {
	id := strconv.FormatInt(uid, 10)
	d := strconv.FormatInt(device, 10)
	err := IM.EnqueueMessage(id, d, m)
	if err != nil {
		logger.E("SendMessage error: %v", err)
	}
}

// Service IM 服务的接口
type Service interface {
	IsOnline(id string, device string) (bool, error)
	Exit(id string, device string) error
	Logout(uid string, device string) error
	EnqueueMessage(uid string, device string, message *messages.GlideMessage) error
}

// MustSetupClient 初始化 IM 服务 RPC 客户端
func MustSetupClient(addr string, port int, name string) {
	opt := &rpc.ClientOptions{
		Addr:        addr,
		Port:        port,
		Name:        name,
		EtcdServers: nil,
		Selector:    nil,
	}
	cli, err := client.NewClient(opt)
	if err != nil {
		panic(err)
	}
	IM = &imServiceRpcClient{cli}
}

// TODO: optimize 2022-7-18 12:20:03 使用缓存用户连接网关等信息
type imServiceRpcClient struct {
	cli *client.Client
}

func (c imServiceRpcClient) IsOnline(id string, device string) (bool, error) {
	online := c.cli.IsOnline(gate.NewID("", id, device))
	return online, nil
}

func (c imServiceRpcClient) Exit(id string, device string) error {
	return c.cli.ExitClient(gate.NewID("", id, device))
}

func (c imServiceRpcClient) Logout(uid string, device string) error {
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

func (c imServiceRpcClient) EnqueueMessage(uid string, device string, message *messages.GlideMessage) error {
	return c.cli.EnqueueMessage(gate.NewID("", uid, device), message)
}
