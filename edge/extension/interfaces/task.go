package interfaces

import "star_cloud/edge/transport/interfaces"

// ITask - 工作任务
type ITask interface {
	// 启动一个工作任务
	Execute() error
	// Stop - 回收资源，停止设备
	Exit() error
	// 工作任务对应的算法，当前是一对一
	SetAlgorithm(algorithm IAlgorithm)
	// 添加接收数据服务端
	SetTransportServer(trans interfaces.IServer)
	// 添加数据传送客户端
	SetTransportClient(trans interfaces.IClient)
	// Manage -
	Manage() (string, error)
}
