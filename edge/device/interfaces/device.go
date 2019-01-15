package interfaces

import "star_cloud/edge/transport/interfaces"

// IDevice -- 设备服务接口
type IDevice interface {
	// 开始工作
	Work() error
	// 停止
	Stop() error
	// 添加接收数据服务端
	SetTransportServer(trans interfaces.IServer)
	// 添加数据传送客户端
	SetTransportClient(trans interfaces.IClient)
	// 添加数据监听
	AddListener(addr string)
}
