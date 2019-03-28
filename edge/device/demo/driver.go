package demo

import (
	"star-edge-cloud/edge/device/interfaces"
	"time"
)

// Driver1 -示例
type Driver1 struct {
	handler interfaces.IDriverHandler
}

// LoadMetadata 加载配置
func (it *Driver1) LoadMetadata(v interface{}) {

}

// SetHandler -
func (it *Driver1) SetHandler(handler interfaces.IDriverHandler) {
	it.handler = handler
}

// Initialize 初始化
func (it *Driver1) Initialize() {

}

// Work 开始工作
func (it *Driver1) Work() error {
	// 模拟设备采集到一些数据，循环发送
	for i := 0; i < 100000000; i++ {
		it.handler.OnReceive(i)
		time.Sleep(1 * time.Microsecond)
	}
	return nil
}

// Stop 停止
func (it *Driver1) Stop() error {
	return nil
}
