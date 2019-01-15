package interfaces

import (
	"encoding/json"
	"fmt"
	"os"
	"star-edge-cloud/edge/models"
	"star-edge-cloud/edge/transport/interfaces"

	"github.com/takama/daemon"
)

// ChStopTask - 停止信号
var ChStopTask chan int

// TaskImp -
type TaskImp struct {
	daemon.Daemon
	// client interfaces.IClient
	server interfaces.IServer
	store  IStore
}

// Execute 带参启动一个设备服务，至少需要配置设备Id信息
func (task TaskImp) Execute() error {
	task.server.SetRealtimeDataHandler(&StoreHandler{Store: task.store})
	task.server.SetResultHandler(&ResultHandler{Store: task.store})
	task.server.SetAlarmHandler(&AlarmHandler{})
	go task.server.Start()
	task.Exit()
	return nil
}

// Exit - 回收资源，停止设备
func (task *TaskImp) Exit() error {
	<-ChStopTask
	return nil
}

// SetTransportServer - 添加接收数据服务端
func (task *TaskImp) SetTransportServer(server interfaces.IServer) {
	task.server = server
}

// SetTransportClient - 添加数据传送客户端
func (task *TaskImp) SetTransportClient(client interfaces.IClient) {

}

// SetStore - 存储
func (task *TaskImp) SetStore(store IStore) {
	task.store = store
}

// Manage by daemon commands or run the daemon
func (task *TaskImp) Manage() (string, error) {

	usage := "Usage: myservice install | remove | start | stop | status"

	// if received any kind of command, do it
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return task.Install()
		case "remove":
			return task.Remove()
		case "start":
			return task.Start()
		case "stop":
			return task.Stop()
		case "status":
			str, err := task.Status()
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(str)
			return str, err
		default:
			return usage, nil
		}
	}
	task.Execute()
	return "", nil
}

// StoreHandler -
type StoreHandler struct {
	Store IStore
}

// Handle -
func (sh *StoreHandler) Handle(request *models.RealtimeData, response *models.Response) error {
	data, err := json.Marshal(request)
	if err != nil {
		response = &models.Response{Status: "false"}
		return err
	}
	if err := sh.Store.Save(request.ID, data); err != nil {
		response = &models.Response{Status: "false"}
		return err
	}

	response = &models.Response{Status: "true"}
	return nil

}

// ResultHandler -
type ResultHandler struct {
	Store IStore
}

// Handle -
func (sh *ResultHandler) Handle(request *models.Result, response *models.Response) error {
	data, err := json.Marshal(request)
	if err != nil {
		response = &models.Response{Status: "false"}
		return err
	}

	if err := sh.Store.Save(request.ID, data); err != nil {
		response = &models.Response{Status: "false"}
		return err
	}

	response = &models.Response{Status: "true"}
	return nil
}

// AlarmHandler -
type AlarmHandler struct {
}

// Handle -
func (sh *AlarmHandler) Handle(request *models.Alarm, response *models.Response) error {
	response.Status = "true"
	return nil

}
