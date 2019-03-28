package client

import (
	"star-edge-cloud/edge/models/constant"
	"star-edge-cloud/edge/transport/client/configuration"
	"star-edge-cloud/edge/transport/client/http"
)

// Client -
type Client struct {
	IsAlive bool
	IsAdle  bool
}

// New -
func New() *Client {
	return &Client{}
}

// GetChannel -
func (it *Client) GetChannel(cc *configuration.ChannelConfig) IChannel {
	switch cc.Protocol {
	case constant.HTTP:
		c := &http.Channel{}
		c.SetConfig(cc)
		return c
	default:
		return &http.Channel{}
	}
}
