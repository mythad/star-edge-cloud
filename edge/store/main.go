package main

import (
	"star-edge-cloud/edge/models/constant"
	"star-edge-cloud/edge/store/bussiness"
	"star-edge-cloud/edge/transport/server/configuration"

	"star-edge-cloud/edge/utils/daemon"
)

func main() {
	// 创建一个task
	task := &bussiness.Task{}
	task.LoadConfig(configuration.ChannelConfig{LoaclAddress: ":18002", Protocol: constant.HTTP})
	// task.LoadMetadata(nil)
	task.Initialize()
	// 启动deamon服务
	daemon.New("store", "存储服务", task).Work()

	// dir := common.GetCurrentDirectory()
	// task := &bussiness.TaskImp{}
	// // config := &configuration.ServerConfig{Channels: []configuration.ServerChannelConfig{
	// // 	configuration.ServerChannelConfig{
	// // 		LoaclAddress: ":8888",
	// // 		Protocol:     constant.HTTP,
	// // 	},
	// // }}
	// // server := transport.NewServer(config)
	// // task.SetTransportServer(server)
	// task.SetStore(&kv.KVStore{Dir: dir + "/data/badger/"})
	// srv, err := daemon.New("store", "存储服务")
	// logger.BaseAddress = "http://localhost:17000"
	// if err != nil {
	// 	logger.Write(&models.LogInfo{ID: common.GetGUID(), Message: err.Error(), Level: 0, Time: time.Now().Unix()})
	// 	os.Exit(1)
	// }
	// task.Daemon = srv
	// _, err = task.Manage()
	// if err != nil {
	// 	logger.Write(&models.LogInfo{ID: common.GetGUID(), Message: err.Error(), Level: 0, Time: time.Now().Unix()})
	// 	os.Exit(1)
	// }
}
