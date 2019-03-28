package interfaces

import "star-edge-cloud/edge/transport/server/configuration"

// IChannel -
type IChannel interface {
	SetConfig(config *configuration.ChannelConfig)
	RegisteHandler(rule string, handler IHandler)
	MakeRouter()
	Open() error
	Close() error
}
