package amqp

// BaseAgreement - 扩展接口协议
type BaseAgreement interface {
	// 推送数据
	Push([]byte) (string, error)
	// 查询数据
	Query(string) ([]byte, error)
	// 配置
	Config(string) error
}

// RestAgreement - 扩展接口协议
type RestAgreement interface {
	BaseAgreement
}

// AMQPAgreement - 扩展接口协议
type AMQPAgreement interface {
	BaseAgreement
	// 建立连接
	Listen() error
	// 断开连接
	Close() error
}

// AMQPConnectionContext - 连接上下文
type AMQPConnectionContext struct {
	// IP地址
	IPAddr string
	// 端口
	Port int
}
