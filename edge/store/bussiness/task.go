package bussiness

import (
	"star-edge-cloud/edge/transport/server"
	"star-edge-cloud/edge/transport/server/configuration"
)

// Task -
type Task struct {
	cmdChConfig *configuration.ChannelConfig
	container   *server.Container
}

// LoadConfig -
func (it *Task) LoadConfig(v interface{}) {
	conf := v.(configuration.ChannelConfig)
	it.cmdChConfig = &conf
}

// LoadMetadata -
func (it *Task) LoadMetadata(v interface{}) {

}

// Initialize -
func (it *Task) Initialize() {
	// 开启接收数据服务
	server := server.New()
	server.RegisteChannel(it.cmdChConfig).
		AddRoute("realtime_data/load", &RealtimeDataHandler{}).
		AddRoute("result/load", &ResultHandler{}).
		AddRoute("alarm/load", &AlarmHandler{}).
		Build()
	it.container = server
}

// Begin -
func (it *Task) Begin() {
	go it.container.Start()
}

// End -
func (it *Task) End() {
	it.container.Stop()
}
