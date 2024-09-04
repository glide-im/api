package im

import (
	"github.com/glide-im/glide/im_service/client"
	"github.com/glide-im/glide/pkg/gate"
	"github.com/glide-im/glide/pkg/logger"
	"github.com/glide-im/glide/pkg/messages"
	"github.com/glide-im/glide/pkg/rpc"
	"github.com/glide-im/glide/pkg/subscription"
	"github.com/glide-im/glide/pkg/subscription/subscription_impl"
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

	// Exit 退出特定用户特定设备, 多个设备需要多次调用, 用户不在线会返回错误
	Exit(id string, device string) error

	// EnqueueMessage 给特定用户发送消息, 如果要给多个设备发需要调用多次传入设备 id
	// 如果用户不在线, 会直接丢弃消息不会推送, 如果需要确定必达, 先判断用户是否在线, 不在线则保存到数据库, 用户上线时拉取
	EnqueueMessage(uid string, device string, message *messages.GlideMessage) error

	// UpdateClientSecret 更新用户 message deliver secret
	UpdateClientSecret(id string, secret string) error

	SubscribeChannel(uid string, channels []string) error

	CreateChannel(id string) error

	UpdateSubscriber(chanId string, uid string, perm subscription_impl.Permission) error

	PublishChannel(channel string, message *messages.ChatMessage) error

	Close()
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
	cli, err := client.NewGatewayRpcImpl(opt)
	if err != nil {
		panic(err)
	}
	sub, err := client.NewSubscriptionRpcImpl(opt)
	if err != nil {
		panic(err)
	}
	IM = &imServiceRpcClient{cli, sub}
}

// TODO: optimize 2022-7-18 12:20:03 使用缓存用户连接网关等信息
type imServiceRpcClient struct {
	// cli 消息服务客户端, 包含群聊, 消息网关接口
	cli *client.GatewayRpcImpl
	sub *client.SubscriptionRpcImpl
}

func (c *imServiceRpcClient) PublishChannel(channel string, message *messages.ChatMessage) error {
	err := c.sub.Publish(subscription.ChanID(channel), &subscription_impl.PublishMessage{
		From:    "system",
		Type:    subscription_impl.TypeMessage,
		Message: messages.NewMessage(0, messages.ActionGroupMessage, message),
	})
	return err
}

func (c *imServiceRpcClient) UpdateSubscriber(chanId string, uid string, perm subscription_impl.Permission) error {
	err := c.sub.UpdateSubscriber(subscription.ChanID(chanId), subscription.SubscriberID(uid), &subscription_impl.SubscriberOptions{Perm: perm})
	return err
}

func (c *imServiceRpcClient) CreateChannel(id string) error {
	err := c.sub.CreateChannel(subscription.ChanID(id), &subscription.ChanInfo{
		ID:      subscription.ChanID(id),
		Type:    subscription.ChanTypeUnknown,
		Muted:   false,
		Blocked: false,
		Closed:  false,
	})
	return err
}

func (c *imServiceRpcClient) Close() {
	c.cli.Close()
	c.sub.Close()
}

func (c *imServiceRpcClient) SubscribeChannel(uid string, channels []string) error {
	for _, channel := range channels {
		err := c.sub.Subscribe(subscription.ChanID(channel), subscription.SubscriberID(uid), &subscription_impl.SubscriberOptions{Perm: subscription_impl.PermNone})
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *imServiceRpcClient) UpdateClientSecret(id string, secret string) error {
	return c.cli.UpdateClient(gate.NewID2(id), &gate.ClientSecrets{MessageDeliverSecret: secret})
}

func (c *imServiceRpcClient) SetID(old string, new string) error {
	return c.cli.SetClientID(gate.NewID2(old), gate.NewID2(new))
}

func (c *imServiceRpcClient) Exit(id string, device string) error {
	return c.cli.ExitClient(gate.NewID("", id, device))
}

func (c *imServiceRpcClient) EnqueueMessage(uid string, device string, message *messages.GlideMessage) error {
	return c.cli.EnqueueMessage(gate.NewID("", uid, device), message)
}
