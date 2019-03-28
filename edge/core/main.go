package main

import (
	"log"
	"os"
	"star-edge-cloud/edge/core/config"
	"star-edge-cloud/edge/core/device"
	"star-edge-cloud/edge/core/service"
	"star-edge-cloud/edge/core/share"
	"star-edge-cloud/edge/utils/common"

	"github.com/joho/godotenv"
)

var conf config.Config

const sql1 = `create table device (id string not null primary key, name text,file_name string, describe string, registry_time string, type string, other string, protocol string, conf string, status integer,liseners string, command_server_address string, log_base_url string);`
const sql2 = `create table extension (id string not null primary key, name text,file_name string, describe string, registry_time string, type string, other string, protocol string, conf string, status integer,liseners string, command_server_address string, log_base_url string);`

func main() {
	log.Println("启动元数据服务程序...")
	godotenv.Load("./conf/conf.env")
	// dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	// log.Println(dir)
	// conf := &config.Config{}
	conf.MetadataDBPath = os.Getenv("Core.MetadataDBPath")
	conf.LogDBPath = os.Getenv("Core.LogDBPath")
	conf.Port = os.Getenv("Core.Port")
	conf.RulesAddr = os.Getenv("Rules.Address")
	conf.SchedulerTaskAddr = os.Getenv("SchedulerTask.Addr")

	// 初始化元数据库
	share.InitializeDB(conf.MetadataDBPath)
	share.CreateTable(sql1)
	share.CreateTable(sql2)
	share.WorkingDir = common.GetCurrentDirectory()

	server := &service.CoreServer{}
	server.Conf = &conf
	server.DevManager.Conf = &conf
	server.ExtManager.Conf = &conf

	// loadDevice(&server.DevManager)

	go func() {
		if err := server.Start(); err != nil {
			log.Println("服务器启动失败:" + err.Error())
		}
	}()

	select {}
}

func loadDevice(manager *device.DeviceManager) {
	arr, _ := manager.QueryAllDevice()
	for _, d := range arr {
		if d.Status == 0 {
			continue
		}
		if err := manager.Run(d.ID); err != nil {
			d.Status = 0
			manager.UpdateDevice(&d)
		}
	}
}
