package server

import (
	"container/list"
	"star-edge-cloud/edge/models/constant"
	"star-edge-cloud/edge/transport/server/configuration"
	"star-edge-cloud/edge/transport/server/http"
	"star-edge-cloud/edge/transport/server/interfaces"
)

// Container -
type Container struct {
	channelList *list.List
}

// New -
func New() (c *Container) {
	c = &Container{}
	if c.channelList == nil {
		c.channelList = list.New()
	}
	return
}

// RegisteChannel -
func (it *Container) RegisteChannel(cc *configuration.ChannelConfig) *Container {
	switch cc.Protocol {
	case constant.HTTP:
		channel := &http.Channel{}
		channel.SetConfig(cc)
		it.channelList.PushBack(channel)
	default:
		channel := &http.Channel{}
		channel.SetConfig(cc)
		it.channelList.PushBack(channel)
	}

	return it
}

// AddRoute -
func (it *Container) AddRoute(rule string, handler interfaces.IHandler) *Container {
	channel := it.channelList.Back().Value.(interfaces.IChannel)
	channel.RegisteHandler(rule, handler)
	return it
}

// Build -
func (it *Container) Build() *Container {
	for e := it.channelList.Front(); e != nil; e = e.Next() {
		channel := e.Value.(interfaces.IChannel)
		channel.MakeRouter()
	}
	return it
}

// Start -
func (it *Container) Start() {
	for e := it.channelList.Front(); e != nil; e = e.Next() {
		go e.Value.(interfaces.IChannel).Open()
	}
}

// Stop -
func (it *Container) Stop() {
	for e := it.channelList.Front(); e != nil; e = e.Next() {
		e.Value.(interfaces.IChannel).Close()
	}
}
