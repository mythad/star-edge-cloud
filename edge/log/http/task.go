package http

import (
	"fmt"
	"os"

	"github.com/takama/daemon"
)

// ChStopTask - 停止信号
var ChStopTask chan int

// TaskImp -
type TaskImp struct {
	daemon.Daemon
	Server *LogServer
}

// Execute 带参启动一个设备服务，至少需要配置设备Id信息
func (task TaskImp) Execute() error {
	go task.Server.Start()
	task.Exit()
	return nil
}

// Exit - 回收资源，停止设备
func (task *TaskImp) Exit() error {
	<-ChStopTask
	return nil
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
