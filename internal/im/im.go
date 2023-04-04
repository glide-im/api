package im

import (
	"github.com/glide-im/glide/pkg/gate"
	"github.com/glide-im/glide/pkg/logger"
	"github.com/glide-im/glide/pkg/messages"
	"github.com/glide-im/glide/pkg/rpc"
	"github.com/glide-im/im-service/pkg/client"
	"strconv"
)

// IM 消息服务客户端实例
var IM Service

// DeviceTypes 登录设备类型
var DeviceTypes = []string{"1", "2", "3"}

// SendMessage 发送消息到指定用户
func SendMessage(uid int64, device int64, m *messages.GlideMessage) {
	id := strconv.FormatInt(uid, 10)
	d := strconv.FormatInt(device, 10)
	err := IM.EnqueueMessage(id, d, m)
	if err != nil {
		logger.E("SendMessage error: %v", err)
	}
}

// SendMessageToAllDevice 发送消息到指定用户所有设备
func SendMessageToAllDevice(uid int64, m *messages.GlideMessage) error {
	for _, deviceType := range DeviceTypes {
		err := IM.EnqueueMessage(strconv.FormatInt(uid, 10), deviceType, m)
		if err != nil {
			return err
		}
	}
	return nil
}

// Service IM 服务的接口
type Service interface {

	// IsOnline 获取指定 id 用户指定设备是否在线, device 目前有四个值 1,2,3,4
	IsOnline(id string, device string) bool

	// Exit 退出特定用户特定设备, 多个设备需要多次调用, 用户不在线会返回错误
	Exit(id string, device string) error

	// EnqueueMessage 给特定用户发送消息, 如果要给多个设备发需要调用多次传入设备 id
	// 如果用户不在线, 会直接丢弃消息不会推送, 如果需要确定必达, 先判断用户是否在线, 不在线则保存到数据库, 用户上线时拉取
	EnqueueMessage(uid string, device string, message *messages.GlideMessage) error

	// SetID 设置用户 ID 老 id 不存在或新 id 已存在都会返回错误
	SetID(old string, new string) error
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
	// cli 消息服务客户端, 包含群聊, 消息网关接口
	cli *client.Client
}

func (c *imServiceRpcClient) SetID(old string, new string) error {
	return c.cli.SetClientID(gate.NewID2(old), gate.NewID2(new))
}

func (c *imServiceRpcClient) IsOnline(id string, device string) bool {
	online := c.cli.IsOnline(gate.NewID("", id, device))
	return online
}

func (c *imServiceRpcClient) Exit(id string, device string) error {
	return c.cli.ExitClient(gate.NewID("", id, device))
}

func (c *imServiceRpcClient) EnqueueMessage(uid string, device string, message *messages.GlideMessage) error {
	return c.cli.EnqueueMessage(gate.NewID("", uid, device), message)
}
