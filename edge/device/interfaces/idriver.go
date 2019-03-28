package interfaces

// Callback 采集数据回调
type Callback func([]byte)

// IDriver -- 设备服务接口
type IDriver interface {
	// 加载配置
	LoadMetadata(interface{})
	// SetHandler -
	SetHandler(IDriverHandler)
	// 初始化
	Initialize()
	// 开始工作
	Work() error
	// 停止
	Stop() error
}
