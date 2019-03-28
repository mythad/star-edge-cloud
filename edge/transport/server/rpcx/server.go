package rpcx

import (
	"context"
	"log"
	"star-edge-cloud/edge/models"
	"star-edge-cloud/edge/transport/server/interfaces"

	"github.com/smallnest/rpcx/server"
)

// RpcxServer - rpcx服务，接收配置、命令等
type RpcxServer struct {
	service RpcxService
	Address string
}

// SetConfig -
func (rs *RpcxServer) SetConfig(conf string) {

}

// Start -
func (rs *RpcxServer) Start() error {
	s := server.NewServer()
	s.RegisterName("transport", rs.service, "")
	// s.AuthFunc = ls.auth
	// fmt.Println(*Config.Address)
	// s.Serve("reuseport", *Config.Address)
	log.Println("RPCX服务器启动，正监听地址：" + rs.Address)
	s.Serve("tcp", rs.Address)
	return nil
}

// Stop -
func (rs *RpcxServer) Stop() error {
	return nil
}

// SetRequestHandler -
func (rs *RpcxServer) SetRequestHandler(handler interfaces.IRequestHandler) {
	rs.service.requestHandler = handler
}

// SetCommandHandler -
func (rs *RpcxServer) SetCommandHandler(handler interfaces.ICommandHandler) {
	rs.service.commandHandler = handler
}

// SetRealtimeDataHandler -
func (rs *RpcxServer) SetRealtimeDataHandler(handler interfaces.IRealtimeDataHandler) {
	rs.service.realtimeDataHandler = handler
}

// SetAlarmHandler -
func (rs *RpcxServer) SetAlarmHandler(handler interfaces.IAlarmHandler) {
	rs.service.alarmHandler = handler
}

// SetFaultHandler -
func (rs *RpcxServer) SetFaultHandler(handler interfaces.IFaultHandler) {
	rs.service.faulttHandler = handler
}

// SetResultHandler -
func (rs *RpcxServer) SetResultHandler(handler interfaces.IResultHandler) {
	rs.service.resultHandler = handler
}

// RpcxService -
type RpcxService struct {
	requestHandler      interfaces.IRequestHandler
	alarmHandler        interfaces.IAlarmHandler
	commandHandler      interfaces.ICommandHandler
	faulttHandler       interfaces.IFaultHandler
	realtimeDataHandler interfaces.IRealtimeDataHandler
	resultHandler       interfaces.IResultHandler
}

func (rse *RpcxService) handleRequest(ctx context.Context, request models.Request, reply *models.Response) {
	var response *models.Response
	rse.requestHandler.Handle(&request, response)
}

func (rse *RpcxService) handleCommand(ctx context.Context, request models.Command, reply *models.Response) {
	var response *models.Response
	rse.commandHandler.Handle(&request, response)
}

func (rse *RpcxService) handleRealtimeData(ctx context.Context, request models.RealtimeData, reply *models.Response) {
	var response *models.Response
	rse.realtimeDataHandler.Handle(&request, response)
}

func (rse *RpcxService) handleAlarm(ctx context.Context, request models.Result, reply *models.Response) {
	var response *models.Response
	rse.resultHandler.Handle(&request, response)
}

func (rse *RpcxService) handleFault(ctx context.Context, request models.Alarm, reply *models.Response) {
	var response *models.Response
	rse.alarmHandler.Handle(&request, response)
}

func (rse *RpcxService) handleResult(ctx context.Context, request models.Fault, reply *models.Response) {
	var response *models.Response
	rse.faulttHandler.Handle(&request, response)
}
