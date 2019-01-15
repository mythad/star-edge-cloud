package amqp

import (
	"context"
	"encoding/json"
	"log"
	"star_cloud/edge/models"
	"strings"
	"time"

	"pack.ag/amqp"
)

// AMQPClient -
type AMQPClient struct{}

// PostRequest -
func (ac *AMQPClient) PostRequest(addr string, request *models.Request) (rsp *models.Response, err error) {
	// Create client
	// client, err := amqp.Dial("amqps://my-namespace.servicebus.windows.net",
	// 	amqp.ConnSASLPlain("access-key-name", "access-key"),
	// )
	path := strings.Split(addr, "/")
	client, err := amqp.Dial(path[0],
		amqp.ConnSASLPlain("access-key-name", "access-key"),
	)
	if err != nil {
		log.Fatal("Dialing AMQP server:", err)
	}
	defer client.Close()

	// Open a session
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Creating AMQP session:", err)
	}

	ctx := context.Background()
	// Create a sender
	sender, err := session.NewSender(
		amqp.LinkTargetAddress(path[1]),
	)
	if err != nil {
		log.Fatal("Creating sender link:", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	// Send message
	data, err := json.Marshal(request)
	err = sender.Send(ctx, amqp.NewMessage(data))
	if err != nil {
		log.Fatal("Sending message:", err)
	}

	sender.Close(ctx)
	cancel()

	return
}

// PostCommand -
func (ac *AMQPClient) PostCommand(addr string, command *models.Command) (rsp *models.Response, err error) {
	// data, err := json.Marshal(command)
	// if err != nil {
	// 	return nil, err
	// }
	// r, err := PostData(addr, data)
	// if err != nil {
	// 	return nil, err
	// }
	// if err = json.Unmarshal(r, &rsp); err != nil {
	// 	return nil, err
	// }

	return
}

// PostRealtimeData -
func (ac *AMQPClient) PostRealtimeData(addr string, data *models.RealtimeData) (rsp *models.Response, err error) {
	return
}

// PostAlarm -
func (ac *AMQPClient) PostAlarm(addr string, alarm *models.Alarm) (rsp *models.Response, err error) {
	return
}

// PostFault -
func (ac *AMQPClient) PostFault(addr string, fault *models.Fault) (rsp *models.Response, err error) {
	return
}

// PostResult -
func (ac *AMQPClient) PostResult(addr string, result *models.Result) (rsp *models.Response, err error) {
	return
}
