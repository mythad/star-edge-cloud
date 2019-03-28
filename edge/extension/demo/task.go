package demo

import (
	"container/list"
	"star-edge-cloud/edge/extension/interfaces"
	"star-edge-cloud/edge/transport/coding/json"
	"star-edge-cloud/edge/transport/server"
	"star-edge-cloud/edge/transport/server/configuration"
)

// Task -
type Task struct {
	cmdChConfig *configuration.ChannelConfig
	algorithm   interfaces.IAlgorithm
	server      *server.Container
}

// LoadConfig -
func (it *Task) LoadConfig(v interface{}) {
	conf := v.(configuration.ChannelConfig)
	it.cmdChConfig = &conf

	rlisteners = list.New()
	// 发送到存储服务
	rlisteners.PushBack("http://localhost:18002/api/realtime_data/load")
}

// LoadMetadata -
func (it *Task) LoadMetadata(v interface{}) {

}

// SetAlgorithm -
func (it *Task) SetAlgorithm(algorithm interfaces.IAlgorithm) {
	it.algorithm = algorithm
}

// Initialize -
func (it *Task) Initialize() {
	handler := &RealtimeDataHandler{}
	handler.SetAlgorithm(it.algorithm)
	handler.SetEncoder(&json.JSONEncoder{})
	handler.SetDecoder(&json.JSONDecoder{})
	// 开启接收数据服务
	server := server.New()
	server.RegisteChannel(it.cmdChConfig).
		AddRoute("realtime_data/load", handler).
		// AddRoute("command", &CommandHandler{}).
		Build()
	it.server = server

	// c := client.New().GetChannel(&cc.ChannelConfig{
	// 	Timeout:  10 * time.Second,
	// 	Keep:     true,
	// 	DataType: constant.RealtimeData,
	// 	Protocol: constant.HTTP,
	// 	Cache:    false,
	// })
	// c.SetDecoder(&json.JSONDecoder{})
	// c.SetEncoder(&json.JSONEncoder{})
	// c.Send("http://localhost:18010/command",
	// 	&models.Command{Type: "addlistener", Data: []byte("http://localhost:18020/realtime_data/load")})
}

// Begin -
func (it *Task) Begin() {
	go it.server.Start()
}

// End -
func (it *Task) End() {
	it.server.Stop()
}
