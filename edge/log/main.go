package main

import (
	"star-edge-cloud/edge/log/bussiness"
	"star-edge-cloud/edge/models/constant"
	"star-edge-cloud/edge/transport/server/configuration"
	"star-edge-cloud/edge/utils/daemon"
)

func main() {
	// 创建一个task
	task := &bussiness.Task{}
	task.LoadConfig(configuration.ChannelConfig{LoaclAddress: ":18005", Protocol: constant.HTTP})
	// task.LoadMetadata(nil)
	task.Initialize()
	// 启动deamon服务
	daemon.New("log", "存储服务", task).Work()

	// dir := common.GetCurrentDirectory()
	// godotenv.Load(dir + "/conf/log_conf.env")
	// // serverAddr := os.Getenv("Log.ServerAddr")
	// dbpath := os.Getenv("Log.DBPath")
	// // log.Println(dbpath)
	// // addr := os.Getenv("Log.ServerAddr")
	// // dbPath := os.Getenv("Log.DBPath")
	// // logger := &implemention.Logger{LogDBPath: "data/sqlite3/log.db"}
	// // if _, err := os.Stat(dbPath); err != nil {
	// // 	logger.InitLogDB()
	// // }
	// // server := &http.LogServer{ServerAddr: ":17000", Logger: logger}
	// // go server.Start()
	// // select {}
	// logger := &bussiness.Logger{LogDBPath: dir + "/" + dbpath}
	// logger.InitLogDB()
	// // logger.LogDBPath = fmt.Sprintf("file:%s?cache=shared", dbpath)
	// // server := &http.LogServer{ServerAddr: serverAddr, Logger: logger}
	// task := &bussiness.TaskImp{}
	// // config := &configuration.ServerConfig{Channels: []configuration.ServerChannelConfig{
	// // 	configuration.ServerChannelConfig{
	// // 		LoaclAddress: serverAddr,
	// // 		Protocol:     constant.HTTP,
	// // 	},
	// // }}
	// // server := transport.NewServer(config)
	// // task.SetTransportServer(server)

	// srv, err := daemon.New("log", "日志服务")
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
