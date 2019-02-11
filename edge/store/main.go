package main

import (
	"os"
	hg "star-edge-cloud/edge/log/http"
	"star-edge-cloud/edge/models"
	"star-edge-cloud/edge/store/implemetion/kv"
	"star-edge-cloud/edge/store/interfaces"
	"star-edge-cloud/edge/transport/http"
	"star-edge-cloud/edge/utils/common"
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
