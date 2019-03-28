package main

import (
	"star-edge-cloud/edge/extension/demo"
	"star-edge-cloud/edge/models/constant"
	"star-edge-cloud/edge/transport/server/configuration"
	"star-edge-cloud/edge/utils/daemon"
)

func main() {
	// 创建一个task
	task := &demo.Task{}
	task.LoadConfig(configuration.ChannelConfig{LoaclAddress: ":18020", Protocol: constant.HTTP})
	// task.LoadMetadata(nil)
	task.SetAlgorithm(&demo.Algorithm1{})
	task.Initialize()
	// 启动deamon服务
	daemon.New("ext1", "××服务", task).Work()

}
