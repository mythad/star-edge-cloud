package main

import (
	"star-edge-cloud/edge/device/demo"
	"star-edge-cloud/edge/models/constant"
	"star-edge-cloud/edge/transport/server/configuration"
	"star-edge-cloud/edge/utils/daemon"
)

func main() {
	// 创建一个task
	task := &demo.Task{}
	task.LoadConfig(configuration.ChannelConfig{LoaclAddress: ":18010", Protocol: constant.HTTP})
	// task.LoadMetadata(nil)
	task.SetDriver(&demo.Driver1{}, &demo.DriverHandler{})
	task.Initialize()
	// 启动deamon服务
	daemon.New("dev3", "××服务", task).Work()
}
