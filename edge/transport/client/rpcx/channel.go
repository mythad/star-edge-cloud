package rpcx

import (
	"context"
	"fmt"
	"star-edge-cloud/edge/models"
	"star-edge-cloud/edge/transport/client/configuration"
	"star-edge-cloud/edge/transport/coding"
	"strings"

	"github.com/smallnest/rpcx/client"
)

// Channel -
type Channel struct {
	config  *configuration.ChannelConfig
	client  *client.Client
	encoder coding.IEncoder
	decoder coding.IDecoder
}

// SetConfig -
func (it *Channel) SetConfig(config *configuration.ChannelConfig) {
	it.config = config
}

// SetDecoder -
func (it *Channel) SetDecoder(decoder coding.IDecoder) {
	it.decoder = decoder
}

// SetEncoder -
func (it *Channel) SetEncoder(encoder coding.IEncoder) {
	it.encoder = encoder
}

// Send -
func (it *Channel) Send(addr string, v interface{}) (interface{}, error) {
	if it.client == nil {
		d := client.NewPeer2PeerDiscovery(addr, "")
		path := strings.Split(addr, "/")
		it.client = client.NewXClient(path[1], client.Failtry, client.RandomSelect, d, client.DefaultOption)
		defer it.client.Close()
	}
	if it.encoder == nil {
		return nil, fmt.Errorf("没有编码器，无法发送数据。")
	}
	request := it.encoder.Encode(v)
	rsp := &models.Response{}
	it.client.Call(context.Background(), path[1], request, rsp)

	if it.decoder == nil {
		return rsp, fmt.Errorf("没有解码器。")
	}
	var result interface{}
	it.decoder.Decode(rsp, result)

	return result, err
}
