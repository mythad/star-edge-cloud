package interfaces

import (
	"fmt"
	"os"
	"star-edge-cloud/edge/transport/interfaces"

	"github.com/takama/daemon"
)

// ChStopTask - 停止信号
var ChStopTask chan int

// TaskImp -
type TaskImp struct {
	daemon.Daemon
	algorithm IAlgorithm
	client    interfaces.IClient
	server    interfaces.IServer
}

// Execute 带参启动一个服务，至少需要配置服务Id信息
func (task TaskImp) Execute() error {
	go task.algorithm.Work()
	go task.server.Start()
	task.Exit()
	return nil
}

// Exit - 回收资源，停止服务
func (task *TaskImp) Exit() error {
	<-ChStopTask
	task.server.Stop()
	task.algorithm.Stop()
	return nil
}

// SetTransportServer - 添加接收数据服务端
func (task *TaskImp) SetTransportServer(server interfaces.IServer) {
	task.server = server
	task.algorithm.SetTransportServer(server)
}

// SetTransportClient - 添加数据传送客户端
func (task *TaskImp) SetTransportClient(client interfaces.IClient) {
	task.client = client
	task.algorithm.SetTransportClient(client)
}

// SetAlgorithm - 工作任务对应的算法，当前是一对一
func (task *TaskImp) SetAlgorithm(algorithm IAlgorithm) {
	task.algorithm = algorithm
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
