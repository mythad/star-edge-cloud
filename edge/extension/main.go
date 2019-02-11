package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"star-edge-cloud/edge/extension/demo"
	"star-edge-cloud/edge/extension/extlog"
	"star-edge-cloud/edge/extension/interfaces"
	"star-edge-cloud/edge/models"
	"star-edge-cloud/edge/transport/http"
	"star-edge-cloud/edge/utils/common"
	"strings"

	"github.com/takama/daemon"
)

func main() {
	var ext models.Extension
	var data []byte
	dir := common.GetCurrentDirectory()
	data, err := ioutil.ReadFile(dir + "/ext.json")
	if err != nil {
		os.Exit(1)
	}
	json.Unmarshal(data, &ext)
	extlog.SetLogServiceURL(ext.LogBaseURL)
	// logger.BaseAddress = ext.LogBaseURL
	// logger.Write(&models.LogInfo{ID: common.GetGUID(), Message: string(data[:]), Level: 0, Time: time.Now()})

	task := &interfaces.TaskImp{}
	a := &demo.Algorithm1{}
	addrs := strings.Split(ext.Listeners, ",")
	for _, item := range addrs {
		a.AddListener(item)
	}
	task.SetAlgorithm(a)
	task.SetTransportServer(&http.RestServer{ServerAddr: ext.ServerAddress})
	// task.SetTransportClient(&http.RestClient{})

	srv, err := daemon.New(ext.ID, ext.Describe)
	if err != nil {
		extlog.WriteLog(err.Error())
		os.Exit(1)
	}
	task.Daemon = srv
	_, err = task.Manage()
	if err != nil {
		extlog.WriteLog(err.Error())
		os.Exit(1)
	}
}
