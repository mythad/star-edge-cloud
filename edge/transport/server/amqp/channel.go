package amqp

import (
	"context"
	"log"
	"star-edge-cloud/edge/models"
	"star-edge-cloud/edge/transport/server/configuration"
	"star-edge-cloud/edge/transport/server/interfaces"
	"time"

	"pack.ag/amqp"
)

type temp struct {
	handler  interfaces.IHandler
	receiver *amqp.Receiver
}

// Channel -
type Channel struct {
	handlers map[string]*temp
	conf     *configuration.ChannelConfig
	// httpServer *http.Server
	client  *amqp.Client
	session *amqp.Session
	ctx     context.Context
}

// SetConfig -
func (it *Channel) SetConfig(config *configuration.ChannelConfig) {
	it.conf = config
}

// RegisteHandler - channel和handler是一对一关系
func (it *Channel) RegisteHandler(rule string, handler interfaces.IHandler) {
	if it.handlers == nil {
		it.handlers = make(map[string]*temp)
	}
	it.handlers[rule] = &temp{handler: handler}
}

// MakeRouter - 统一一个接口，但是请求由不同的携程处理
func (it *Channel) MakeRouter() {
	// Create client
	client, err := amqp.Dial("amqps://127.0.0.1",
		amqp.ConnSASLPlain("root", "123456"),
	)
	it.client = client
	if err != nil {
		log.Fatal("Dialing AMQP server:", err)
	}
	// defer client.Close()

	// Open a session
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Creating AMQP session:", err)
	}
	it.ctx = context.Background()
	it.session = session

	for _, v := range it.handlers {
		// Create a receiver
		receiver, err := session.NewReceiver(
			amqp.LinkSourceAddress("/queue-name"),
			amqp.LinkCredit(10),
		)
		if err != nil {
			log.Fatal("Creating receiver link:", err)
		}
		v.receiver = receiver
	}
}

// Open -
func (it *Channel) Open() error {
	for _, v := range it.handlers {
		go func(receiver *amqp.Receiver, handler interfaces.IHandler) {
			reply := &models.Response{}
			for {
				// Receive next message
				msg, err := receiver.Receive(it.ctx)
				if err != nil {
					log.Fatal("Reading message from AMQP:", err)
				}
				// Accept message
				msg.Accept()
				defer func() {
					ctx, cancel := context.WithTimeout(it.ctx, 1*time.Second)
					receiver.Close(ctx)
					cancel()
				}()
				// 没有数据处理器，发送错误
				if handler == nil {
					reply.Status = "404"
					return
				}

				// 将数据发送给用户定义的处理器，如果处理出错，返回错误
				if request, err := handler.OnReceive(msg.GetData()); err == nil {
					reply.Status = "success"
					reply.Message = request
				} else {
					reply.Status = "404"
				}
			}
		}(v.receiver, v.handler)
	}

	return nil
}

// Close -
func (it *Channel) Close() error {
	return nil
}
