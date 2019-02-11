package interfaces

import "star-edge-cloud/edge/transport/interfaces"

// IAlgorithm --算法
type IAlgorithm interface {
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
