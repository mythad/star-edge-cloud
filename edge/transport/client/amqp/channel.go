package amqp

import (
	"context"
	"fmt"
	"log"
	"star-edge-cloud/edge/transport/client/configuration"
	"star-edge-cloud/edge/transport/coding"
	"time"

	"pack.ag/amqp"
)

// Channel -
type Channel struct {
	config  *configuration.ChannelConfig
	client  *amqp.Client
	encoder coding.IEncoder
	decoder coding.IDecoder
}

// SetConfig -
func (it *Channel) SetConfig(config *configuration.ChannelConfig) {
	it.config = config
}

// SetDecoder -
func (it *Channel) SetDecoder(decoder coding.IDecoder) {
	it.decoder = decoder
}

// SetEncoder -
func (it *Channel) SetEncoder(encoder coding.IEncoder) {
	it.encoder = encoder
}

// Send -
func (it *Channel) Send(addr string, v interface{}) (interface{}, error) {
	// Create client
	if it.client == nil {
		client, err := amqp.Dial("amqps://127.0.0.1",
			amqp.ConnSASLPlain("root", "123456"),
		)
		it.client = client
		if err != nil {
			log.Fatal("Dialing AMQP server:", err)
		}
		defer it.client.Close()
	}
	// Open a session
	session, err := it.client.NewSession()
	if err != nil {
		log.Fatal("Creating AMQP session:", err)
	}
	if it.encoder == nil {
		return nil, fmt.Errorf("没有编码器，无法发送数据。")
	}
	request := it.encoder.Encode(v)
	// Create a sender
	sender, err := session.NewSender(
		amqp.LinkTargetAddress("/queue-name"),
	)
	if err != nil {
		log.Fatal("Creating sender link:", err)
	}
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	// Send message
	err = sender.Send(ctx, amqp.NewMessage(request))
	if err != nil {
		log.Fatal("Sending message:", err)
	}

	sender.Close(ctx)
	cancel()

	// Create a receiver
	receiver, err := session.NewReceiver(
		amqp.LinkSourceAddress("/queue-name"),
		amqp.LinkCredit(10),
	)
	if err != nil {
		log.Fatal("Creating receiver link:", err)
	}
	defer func() {
		ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
		receiver.Close(ctx)
		cancel()
	}()
	msg, err := receiver.Receive(ctx)
	if err != nil {
		log.Fatal("Reading message from AMQP:", err)
	}

	// Accept message
	msg.Accept()

	response := msg.GetData()
	if it.decoder == nil {
		return response, fmt.Errorf("没有解码器。")
	}
	var result interface{}
	it.decoder.Decode(response, result)
	return result, err
}
