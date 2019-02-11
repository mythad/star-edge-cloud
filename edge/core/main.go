package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"star-edge-cloud/edge/core/config"
	"star-edge-cloud/edge/core/device"
	"star-edge-cloud/edge/core/service"

	"github.com/joho/godotenv"
)

var conf config.Config

func main() {
	log.Println("启动元数据服务程序...")
	godotenv.Load("./conf/conf.env")
	// dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	// log.Println(dir)
	// conf := &config.Config{}
	conf.MetadataDBPath = os.Getenv("Core.MetadataDBPath")
	conf.LogDBPath = os.Getenv("Core.LogDBPath")
	conf.Port = os.Getenv("Core.Port")

	// 初始化元数据库
	if _, err := os.Stat(conf.MetadataDBPath); err != nil {
		initMetadataDB()
	}
	conf.MetadataDBPath = fmt.Sprintf("file:%s?cache=shared", conf.MetadataDBPath)

	server := &service.CoreServer{}
	server.Conf = &conf
	server.DevManager.Conf = &conf
	server.ExtManager.Conf = &conf

	loadDevice(&server.DevManager)

	go func() {
		if err := server.Start(); err != nil {
			log.Println("服务器启动失败:" + err.Error())
		}
	}()

	select {}
}

// initMetadataDB - 初始化元数据库
func initMetadataDB() {
	sqlStmt1 := `
	create table device (id string not null primary key, name text,file_name string, describe string, registry_time string, type string, other string, protocol string, conf string, status integer,liseners string, command_server_address string, log_base_url string);
	`

	sqlStmt3 := `
	create table extension (id string not null primary key, name text,file_name string, describe string, registry_time string, type string, other string, protocol string, conf string, status integer,liseners string, command_server_address string, log_base_url string);
	delete from extension;
	`

	db1, err := sql.Open("sqlite3", conf.MetadataDBPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db1.Close()

	if _, err = db1.Exec(sqlStmt1); err != nil {
		log.Printf("%q: %s\n", err, sqlStmt1)
	}

	if _, err = db1.Exec(sqlStmt3); err != nil {
		log.Printf("%q: %s\n", err, sqlStmt3)
	}
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
