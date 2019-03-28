package main

import (
	"star-edge-cloud/edge/models/constant"
	"star-edge-cloud/edge/scheduler/bussiness"
	"star-edge-cloud/edge/transport/server/configuration"
	"star-edge-cloud/edge/utils/daemon"
)

func main() {
	// 创建一个task
	task := &bussiness.Task{}
	task.LoadConfig(configuration.ChannelConfig{LoaclAddress: ":18004", Protocol: constant.HTTP})
	// task.LoadMetadata(nil)
	task.Initialize()
	// 启动deamon服务
	daemon.New("scheduler", "调度服务", task).Work()
	// dir := common.GetCurrentDirectory()
	// godotenv.Load(dir + "/conf/scheduler_conf.env")
	// // serverAddr := os.Getenv("Scheduler.ServerAddr")
	// dbpath := os.Getenv("Scheduler.DBPath")

	// task1 := &bussiness.TaskImp{DBPath: dbpath}
	// // server := &http.RestServer{ServerAddr: serverAddr}
	// // client := &http.RestClient{}
	// // task1.SetTransportServer(server)
	// // config := &configuration.ServerConfig{Channels: []configuration.ServerChannelConfig{
	// // 	configuration.ServerChannelConfig{
	// // 		LoaclAddress: serverAddr,
	// // 		Protocol:     constant.HTTP,
	// // 	},
	// // }}
	// // server := transport.NewServer(config)
	// // task1.SetTransportServer(server)
	// // task1.SetTransportClient(client)
	// task1.Manager = bussiness.NewManager()
	// task1.Manager.MetadataDBPath = dbpath
	// task1.Manager.InitDB(dbpath)
	// task1.Manager.MetadataDBPath = fmt.Sprintf("file:%s?cache=shared", dbpath)
	// task1.Manager.Load()

	// srv, err := daemon.New("scheduler", "规则引擎")
	// if err != nil {
	// 	// extlog.WriteLog(err.Error())
	// 	os.Exit(1)
	// }
	// task1.Daemon = srv
	// _, err = task1.Manage()
	// if err != nil {
	// 	// extlog.WriteLog(err.Error())
	// 	os.Exit(1)
	// }
}
