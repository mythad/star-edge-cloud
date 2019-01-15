package main

import (
	"os"
	"star_cloud/edge/log/http"
	"star_cloud/edge/log/implemention"
	"star_cloud/edge/models"
	"star_cloud/edge/utils/common"
	"time"

	"github.com/joho/godotenv"
	"github.com/takama/daemon"
)

func main() {
	dir := common.GetCurrentDirectory()
	godotenv.Load(dir + "/conf/log_conf.env")
	serverAddr := os.Getenv("Log.ServerAddr")
	dbpath := os.Getenv("Log.DBPath")
	// log.Println(dbpath)
	// addr := os.Getenv("Log.ServerAddr")
	// dbPath := os.Getenv("Log.DBPath")
	// logger := &implemention.Logger{LogDBPath: "data/sqlite3/log.db"}
	// if _, err := os.Stat(dbPath); err != nil {
	// 	logger.InitLogDB()
	// }
	// server := &http.LogServer{ServerAddr: ":17000", Logger: logger}
	// go server.Start()
	// select {}
	logger := &implemention.Logger{LogDBPath: dir + "/" + dbpath}
	logger.InitLogDB()
	// logger.LogDBPath = fmt.Sprintf("file:%s?cache=shared", dbpath)
	server := &http.LogServer{ServerAddr: serverAddr, Logger: logger}
	task := &http.TaskImp{Server: server}
	srv, err := daemon.New("log", "日志服务")
	if err != nil {
		logger.Write(&models.LogInfo{ID: common.GetGUID(), Message: err.Error(), Level: 0, Time: time.Now()})
		os.Exit(1)
	}
	task.Daemon = srv
	_, err = task.Manage()
	if err != nil {
		logger.Write(&models.LogInfo{ID: common.GetGUID(), Message: err.Error(), Level: 0, Time: time.Now()})
		os.Exit(1)
	}
}
