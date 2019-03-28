package bussiness

import (
	"star-edge-cloud/edge/rules_engine/interfaces"
	"star-edge-cloud/edge/transport/server"
	"star-edge-cloud/edge/transport/server/configuration"
	"star-edge-cloud/edge/utils/common"
)

// Task -
type Task struct {
	cmdChConfig *configuration.ChannelConfig
	container   *server.Container
	engine      interfaces.IRulesEngine
}

// LoadConfig -
func (it *Task) LoadConfig(v interface{}) {
	conf := v.(configuration.ChannelConfig)
	it.cmdChConfig = &conf
}

// LoadMetadata -
func (it *Task) LoadMetadata(v interface{}) {
	data := common.ReadString(v.(string))
	it.engine.LoadRules([]byte(data))
}

// SetRulesEngine -
func (it *Task) SetRulesEngine(engine interfaces.IRulesEngine) {
	it.engine = engine
}

// Initialize -
func (it *Task) Initialize() {
	// handler := &LogHandler{}
	// handler.SetLogger(&Logger{})
	// handler.SetDecoder(&json.JSONDecoder{})
	// handler.SetEncoder(&json.JSONEncoder{})
	// 开启接收数据服务
	server := server.New()
	server.RegisteChannel(it.cmdChConfig).
		AddRoute("rules/realtime_data", &RealtimeDataHandler{}).
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
