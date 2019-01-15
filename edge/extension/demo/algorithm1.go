package demo

import (
	"container/list"
	"os"
	"star_cloud/edge/extension/extlog"
	"star_cloud/edge/extension/interfaces"
	"star_cloud/edge/models"
	"star_cloud/edge/transport/http"
	tin "star_cloud/edge/transport/interfaces"
	"star_cloud/edge/utils/common"
)

// var realtimeDataList *list.List

// Algorithm1 - 虚拟设备
type Algorithm1 struct {
	client    tin.IClient
	server    tin.IServer
	Listeners *list.List
}

// Work - 服务
func (d *Algorithm1) Work() error {
	if d.Listeners == nil {
		d.Listeners = list.New()
	}
	d.server.SetCommandHandler(&CommandHandler{})
	d.server.SetRealtimeDataHandler(&RealtimeDataHandler{a1: d, client: &http.RestClient{}})

	return nil
}

// Stop - 停止服务
func (d *Algorithm1) Stop() error {
	return nil
}

// SetTransportServer - 添加接收数据服务端
func (d *Algorithm1) SetTransportServer(server tin.IServer) {
	d.server = server
}

// SetTransportClient - 添加数据传送客户端
func (d *Algorithm1) SetTransportClient(client tin.IClient) {
	d.client = client
}

// AddListener -
func (d *Algorithm1) AddListener(addr string) {
	if d.Listeners == nil {
		d.Listeners = list.New()
	}
	d.Listeners.PushBack(addr)
}

// CommandHandler 接收到数据，调用回调方法
type CommandHandler struct {
	al *Algorithm1
}

// Handle - 处理命令
func (ch *CommandHandler) Handle(request *models.Command, response *models.Response) error {
	if request.Type == "exit" {
		interfaces.ChStopTask <- 0
		response = &models.Response{Status: "true"}
		return nil
	}
	if request.Type == "addlistener" {
		str := ""
		for e := ch.al.Listeners.Front(); e != nil; e = e.Next() {
			str = str + e.Value.(string) + ","
		}
		os.Setenv("Listeners", str)
		response = &models.Response{Status: "true"}
		return nil
	}
	response = &models.Response{Status: "false"}
	return nil
}

// RealtimeDataHandler 接收到数据，调用回调方法
type RealtimeDataHandler struct {
	a1     *Algorithm1
	client tin.IClient
}

// Handle - 如果是10的倍数数据，虚拟为报警数据
// TIME_WAIT - 可能原因是传输太快，数据存储或发送未完成，留待观察
// 所以此处尽量减少连接，数据压缩后再传输?
// 修改此问题，也可以通过数据缓存？
func (rh *RealtimeDataHandler) Handle(request *models.RealtimeData, response *models.Response) error {
	i, _ := common.StringToInt(string(request.Data[:]))
	// extlog.WriteLog("收到实时数据:" + request.ID)
	if i%10 == 0 {
		response = &models.Response{Status: "true"}
		alarm := &models.Alarm{}
		alarm.ID = common.GetGUID()
		alarm.Data = request.Data
		for e := rh.a1.Listeners.Front(); e != nil; e = e.Next() {
			rh.client.PostAlarm(e.Value.(string)+"/api/transport/alarm", alarm)
			extlog.WriteLog("发送报警数据:" + e.Value.(string) + "/api/transport/alarm" + alarm.ID)
		}
		return nil
	}

	response = &models.Response{Status: "false"}
	return nil
}
