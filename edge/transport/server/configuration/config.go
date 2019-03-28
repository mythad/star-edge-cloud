package configuration

import (
	"star-edge-cloud/edge/models/constant"
)

// ChannelConfig -
type ChannelConfig struct {
	Name string
	// 超时时间
	// Timeout int64
	// 是否保持连接
	// Keep bool
	// 数据类型
	// json, xml,
	// DataType string
	// http,rpcx,mqtt,amqp,modbus...
	Protocol constant.Protocol
	// 是否缓存数据
	// Cache bool
	// 服务端端口设置，绑定IP等，比如：
	// http -- :8080
	// tcp -- 127.0.0.1:8001
	LoaclAddress string
}
