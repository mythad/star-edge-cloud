package main

import (
	"os"
	hg "star_cloud/edge/log/http"
	"star_cloud/edge/models"
	"star_cloud/edge/store/implemetion/kv"
	"star_cloud/edge/store/interfaces"
	"star_cloud/edge/transport/http"
	"star_cloud/edge/utils/common"
	"time"

	"github.com/takama/daemon"
)

var logger hg.HttpLogClient

func main() {
	dir := common.GetCurrentDirectory()
	task := &interfaces.TaskImp{}
	task.SetTransportServer(&http.RestServer{ServerAddr: ":18000"})
	task.SetStore(&kv.KVStore{Dir: dir + "/data/badger/"})
	srv, err := daemon.New("store", "存储服务")
	logger.BaseAddress = "http://localhost:17000"
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
