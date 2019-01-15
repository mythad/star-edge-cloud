package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"star_cloud/edge/device/demo"
	"star_cloud/edge/device/devlog"
	"star_cloud/edge/device/interfaces"
	"star_cloud/edge/models"
	"star_cloud/edge/transport/http"
	"star_cloud/edge/utils/common"
	"strings"

	"github.com/takama/daemon"
)

// var logger hg.HttpLogClient

// var configFile = flag.String("conf", "", "Use -file <filesource>")

func main() {
	// 从dev.json加载配置
	var device models.Device
	var data []byte
	dir := common.GetCurrentDirectory()
	data, err := ioutil.ReadFile(dir + "/dev.json")
	if err != nil {
		os.Exit(1)
	}
	json.Unmarshal(data, &device)
	// 设置日志服务的URL
	devlog.SetLogServiceURL(device.LogBaseURL)
	// logger.BaseAddress = device.LogBaseURL
	// logger.Write(&models.LogInfo{ID: common.GetGUID(), Message: string(data[:]), Level: 0, Time: time.Now()})

	// 创建Task
	task := &interfaces.TaskImp{}
	d := &demo.Device1{}
	addrs := strings.Split(device.Listeners, ",")
	for _, item := range addrs {
		d.AddListener(item)
	}
	task.SetDevice(d)
	// 开启接收数据服务
	task.SetTransportServer(&http.RestServer{ServerAddr: device.ServerAddress})
	task.SetTransportClient(&http.RestClient{})

	// daemon守护进程
	srv, err := daemon.New(device.ID, device.Describe)
	if err != nil {
		devlog.WriteLog(err.Error())
		os.Exit(1)
	}
	task.Daemon = srv
	_, err = task.Manage()
	if err != nil {
		devlog.WriteLog(err.Error())
		os.Exit(1)
	}
}
