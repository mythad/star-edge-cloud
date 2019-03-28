package main

import (
	"star-edge-cloud/edge/models/constant"
	"star-edge-cloud/edge/rules_engine/bussiness"
	"star-edge-cloud/edge/transport/server/configuration"
	"star-edge-cloud/edge/utils/common"
	"star-edge-cloud/edge/utils/daemon"
)

func main() {
	// 创建一个task
	task := &bussiness.Task{}
	task.LoadConfig(configuration.ChannelConfig{LoaclAddress: ":18003", Protocol: constant.HTTP})
	task.SetRulesEngine(&bussiness.YQLRulesEngine{})
	task.LoadMetadata(common.ReadString(common.GetCurrentDirectory() + "/rules.xml"))
	task.Initialize()
	// 启动deamon服务
	daemon.New("rules_engine", "存储服务", task).Work()
	// dir := common.GetCurrentDirectory()
	// // 加载规则配置
	// // data := common.ReadString(dir + "/rules.xml")
	// // err := xml.Unmarshal([]byte(data), &models.RulesCollection)
	// // if err == nil {
	// // 	log.Println(err.Error())
	// // 	os.Exit(1)
	// // }

	// godotenv.Load(dir + "/conf/rules_engine_conf.env")
	// // serverAddr := os.Getenv("RulesEngine.ServerAddr")
	// // dbpath := os.Getenv("RulesEngine.DBPath")

	// // var conf models.Config
	// // var data []byte
	// // data, err := ioutil.ReadFile(dir + "/rule.json")
	// // if err != nil {
	// // 	os.Exit(1)
	// // }
	// // json.Unmarshal(data, &ext)
	// // extlog.SetLogServiceURL(ext.LogBaseURL)
	// // logger.BaseAddress = ext.LogBaseURL
	// // logger.Write(&models.LogInfo{ID: common.GetGUID(), Message: string(data[:]), Level: 0, Time: time.Now()})

	// task1 := &bussiness.TaskImp{}
	// // config := &configuration.ServerConfig{Channels: []configuration.ServerChannelConfig{
	// // 	configuration.ServerChannelConfig{
	// // 		LoaclAddress: serverAddr,
	// // 		Protocol:     constant.HTTP,
	// // 	},
	// // }}
	// // server := transport.NewServer(config)
	// // task1.SetTransportServer(server)
	// // task.SetTransportClient(&http.RestClient{})
	// // task := &interfaces.TaskImp{}
	// // a := &demo.Algorithm1{}
	// // addrs := strings.Split(ext.Listeners, ",")
	// // for _, item := range addrs {
	// // 	a.AddListener(item)
	// // }
	// // task.SetAlgorithm(a)
	// // task.SetTransportServer(&http.RestServer{ServerAddr: ext.ServerAddress})
	// srv, err := daemon.New("rules engine", "规则引擎")
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
