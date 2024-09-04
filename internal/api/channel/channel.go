package channel

import (
	route "github.com/glide-im/api/internal/api/router"
)

type ChannelApi struct {
}

func NewChannelApi() *ChannelApi {
	return &ChannelApi{}
}

func (c *ChannelApi) CreateChannel(ctx *route.Context) error {
	return nil
}

func (c *ChannelApi) GetChannelList(ctx *route.Context) error {

	return nil
}
