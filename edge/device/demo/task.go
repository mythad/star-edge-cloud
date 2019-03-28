package demo

import (
	"container/list"
	"star-edge-cloud/edge/device/interfaces"
	"star-edge-cloud/edge/transport/coding/json"
	"star-edge-cloud/edge/transport/server"
	"star-edge-cloud/edge/transport/server/configuration"
)

// Task -
type Task struct {
	cmdChConfig *configuration.ChannelConfig
	driver      interfaces.IDriver
	server      *server.Container
}

// LoadConfig -
func (it *Task) LoadConfig(v interface{}) {
	conf := v.(configuration.ChannelConfig)
	it.cmdChConfig = &conf

	// 可以读取配置文件或从数据库中获取，也可以通过command，注册监听
	listeners = list.New()
	// 发送到算法服务
	listeners.PushBack("http://localhost:18020/api/realtime_data/load")
	// 发送到存储服务
	listeners.PushBack("http://localhost:18002/api/realtime_data/load")
}

// LoadMetadata -
func (it *Task) LoadMetadata(v interface{}) {
	it.driver.LoadMetadata(v)
}

// SetDriver -
func (it *Task) SetDriver(driver interfaces.IDriver, handler interfaces.IDriverHandler) {
	it.driver = driver
	if handler != nil {
		driver.SetHandler(handler)
	}
}

// Initialize -
func (it *Task) Initialize() {
	handler := &CommandHandler{}
	handler.SetEncoder(&json.JSONEncoder{})
	handler.SetDecoder(&json.JSONDecoder{})
	// 开启接收数据服务
	server := server.New()
	server.RegisteChannel(it.cmdChConfig).
		// AddRoute("realtime_data/load", &RealtimeDataHandler{}).
		AddRoute("command", handler).
		Build()
	it.server = server

	// 初始化驱动
	if it.driver != nil {
		it.driver.Initialize()
	}
}

// Begin -
func (it *Task) Begin() {
	go it.server.Start()
	go it.driver.Work()
}

// End -
func (it *Task) End() {
	it.server.Stop()
	it.driver.Stop()
}
