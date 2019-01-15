package demo

import (
	"container/list"
	"star_cloud/edge/device/interfaces"
	"star_cloud/edge/models"
	tin "star_cloud/edge/transport/interfaces"
	"star_cloud/edge/utils/common"
	"time"
)

// Device1 - 虚拟设备
type Device1 struct {
	client tin.IClient
	server tin.IServer
	// todo:持久化监听列表？
	listeners *list.List
}

// Work - 带参服务
func (d *Device1) Work() error {
	d.server.SetRequestHandler(&RequestHandler{})
	d.server.SetCommandHandler(&CommandHandler{device: d})
	if d.listeners == nil {
		d.listeners = list.New()
	}
	for i := 0; i < 100000; i++ {
		for e := d.listeners.Front(); e != nil; e = e.Next() {
			rd := &models.RealtimeData{}
			rd.ID = common.GetGUID()
			rd.Type = "realtimedata"
			rd.Data = []byte(common.IntToString(i))
			go d.client.PostRealtimeData(e.Value.(string)+"/api/transport/realtimedata", rd)
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

// Stop - 停止服务
func (d *Device1) Stop() error {
	return nil
}

// SetTransportServer - 添加接收数据服务端
func (d *Device1) SetTransportServer(server tin.IServer) {
	d.server = server
}

// SetTransportClient - 添加数据传送客户端
func (d *Device1) SetTransportClient(client tin.IClient) {
	d.client = client
}

// AddListener -
func (d *Device1) AddListener(addr string) {
	if d.listeners == nil {
		d.listeners = list.New()
	}
	d.listeners.PushBack(addr)
}

// RequestHandler 接收到数据，调用回调方法
type RequestHandler struct {
}

// Handle -
func (rh *RequestHandler) Handle(request *models.Request, response *models.Response) error {
	return nil
}

// CommandHandler 接收到数据，调用回调方法
type CommandHandler struct {
	device *Device1
}

// Handle -
func (ch *CommandHandler) Handle(request *models.Command, response *models.Response) error {
	if request.Type == "exit" {
		response.Message = []byte("success")
		interfaces.ChStopTask <- 0
	}
	if request.Type == "addlistener" {
		ch.device.AddListener(string(request.Data[:]))
		response.Message = []byte("success")
	}

	return nil
}

// // ICommandHandler -
// type ICommandHandler struct {
// 	Handle(command *models.Command, response *models.Response) error
// }

// // IRealtimeDataHandler -
// type IRealtimeDataHandler struct {
// 	Handle(data *models.RealtimeData, response *models.Response) error
// }

// // IAlarmHandler -
// type IAlarmHandler struct {
// 	Handle(alart *models.Alarm, response *models.Response) error
// }

// // IFaultHandler -
// type IFaultHandler struct {
// 	Handle(fault *models.Fault, response *models.Response) error
// }

// // IResultHandler -
// type IResultHandler struct {
// 	Handle(result *models.Result, response *models.Response) error
// }
