package daemon

import (
	"fmt"
	"os"
	"star-edge-cloud/edge/utils/interfaces"

	"github.com/takama/daemon"
)

// Service -
type Service struct {
	daemon.Daemon
	task interfaces.ITask
}

// 停止信号
var stop = make(chan int)

// New - 新建服务
func New(name, describe string, task interfaces.ITask) *Service {
	service := &Service{}
	das, err := daemon.New(name, describe)
	if err != nil {
		// devlog.WriteLog(err.Error())
		os.Exit(1)
	}
	service.Daemon = das
	service.task = task
	return service
}

// Work -
func (it *Service) Work() {
	_, err := it.manage()
	if err != nil {
		// devlog.WriteLog(err.Error())
		os.Exit(1)
	}
}

// manage by daemon commands or run the daemon
func (it *Service) manage() (string, error) {
	usage := "Usage: myservice install | remove | start | stop | status"
	// if received any kind of command, do it
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return it.Install()
		case "remove":
			return it.Remove()
		case "start":
			return it.Start()
		case "stop":
			return it.Stop()
		case "status":
			str, err := it.Status()
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(str)
			return str, err
		default:
			return usage, nil
		}
	}
	it.execute()
	return "", nil
}

func (it *Service) execute() {
	// 运行设备，比如开始采集数据
	if it.task != nil {
		go it.task.Begin()
	}
	stop <- 1
}

func (it *Service) exit() {
	it.task.End()
	// <-stop
}
