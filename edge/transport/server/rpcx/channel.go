package rpcx

import (
	"context"
	"log"
	"star-edge-cloud/edge/models"
	"star-edge-cloud/edge/transport/server/configuration"
	"star-edge-cloud/edge/transport/server/interfaces"

	"github.com/smallnest/rpcx/server"
)

// Channel - rpcx服务，接收配置、命令等
type Channel struct {
	handlers   map[string]interfaces.IHandler
	conf       *configuration.ChannelConfig
	rpcxServer *server.Server
}

// SetConfig -
func (it *Channel) SetConfig(config *configuration.ChannelConfig) {
	it.conf = config
}

// RegisteHandler -
func (it *Channel) RegisteHandler(rule string, handler IHandler) {
	if it.handlers == nil {
		it.handlers = make(map[string]interfaces.IHandler)
	}
	it.handlers[rule] = handler 
} 

// MakeRouter -
func (it *Channel) MakeRouter() {
	it.rpcxServer = server.NewServer()
	for k, e := range it.handlers {
		service := &RpcxService{Rule: k}
		it.rpcxServer.RegisterName(k, service, "")
	}
}

// Open -
func (it *Channel) Open() error {
	log.Println("RPCX服务器启动，开始监听地址：" + it.Address)
	s.Serve("tcp", it.Address)
	return nil
}

// Close -
func (it *Channel) Close() error {
	return nil
}

// RpcxService -
type RpcxService struct {
	Rule    string
	handler interfaces.IHandler
}

func (it *RpcxService) handleRequest(ctx context.Context, data []byte, reply *models.Response) {
	reply = &models.Response{}
	go func() {
		if it.handler == nil {
			reply.Status ="404"
			return
		}

		// 将数据发送给用户定义的处理器，如果处理出错，返回错误
		if request, err := it.handler.OnReceive(data); err == nil {
			reply.Status ="success"
			reply.Message = request
		} else {
			reply.Status = err.Error()
		}
	}()
}
