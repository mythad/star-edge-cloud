package client

import (
	"star-edge-cloud/edge/transport/client/configuration"
	"star-edge-cloud/edge/transport/coding"
)

// IChannel -
type IChannel interface {
	SetConfig(config *configuration.ChannelConfig)
	SetDecoder(coding.IDecoder)
	SetEncoder(coding.IEncoder)
	Send(string, interface{}, map[string]interface{}) (interface{}, error)
}
