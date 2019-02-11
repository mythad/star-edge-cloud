package amqp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"star-edge-cloud/edge/models"
	"star-edge-cloud/edge/transport/interfaces"
	"strings"

	"pack.ag/amqp"
)

// RestServer -
type RestServer struct {
	requestHandler      interfaces.IRequestHandler
	alarmHandler        interfaces.IAlarmHandler
	commandHandler      interfaces.ICommandHandler
	faulttHandler       interfaces.IFaultHandler
	realtimeDataHandler interfaces.IRealtimeDataHandler
	resultHandler       interfaces.IResultHandler
	conf                string
	ctx                 context.Context
	receiver            *amqp.Receiver
}

// Start - 启动服务
func (cs *RestServer) Start() error {
	// Create client
	path := strings.Split(cs.conf, ",")
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

	cs.ctx = context.Background()
	// Create a receiver
	cs.receiver, err = session.NewReceiver(
		amqp.LinkSourceAddress(path[1]),
		amqp.LinkCredit(10),
	)
	if err != nil {
		log.Fatal("Creating receiver link:", err)
	}
	return nil
}

// SetConfig -
func (cs *RestServer) SetConfig(conf string) {
	cs.conf = conf
}

// Stop - 停止服务
func (cs *RestServer) Stop() error {
	return nil
}

// SetRequestHandler -
func (cs *RestServer) SetRequestHandler(handler interfaces.IRequestHandler) {
	cs.requestHandler = handler
}

// SetCommandHandler -
func (cs *RestServer) SetCommandHandler(handler interfaces.ICommandHandler) {
	cs.commandHandler = handler
}

// SetRealtimeDataHandler -
func (cs *RestServer) SetRealtimeDataHandler(handler interfaces.IRealtimeDataHandler) {
	cs.realtimeDataHandler = handler
}

// SetAlarmHandler -
func (cs *RestServer) SetAlarmHandler(handler interfaces.IAlarmHandler) {
	cs.alarmHandler = handler
}

// SetFaultHandler -
func (cs *RestServer) SetFaultHandler(handler interfaces.IFaultHandler) {
	cs.faulttHandler = handler
}

// SetResultHandler -
func (cs *RestServer) SetResultHandler(handler interfaces.IResultHandler) {
	cs.resultHandler = handler
}

func (cs *RestServer) handleRequest(w http.ResponseWriter, r *http.Request) {
	// defer func() {
	// 	ctx, cancel := context.WithTimeout(cs.ctx, 1*time.Second)
	// 	cs.receiver.Close(ctx)
	// 	cancel()
	// }()

	for {
		// Receive next message
		msg, err := cs.receiver.Receive(cs.ctx)
		if err != nil {
			log.Fatal("Reading message from AMQP:", err)
		}

		// Accept message
		msg.Accept()

		fmt.Printf("Message received: %s\n", msg.GetData())
		if cs.requestHandler == nil {
			io.WriteString(w, "error")
		}

		var request *models.Request
		var response *models.Response
		if err := json.Unmarshal(msg.GetData(), &request); err == nil {
			cs.requestHandler.Handle(request, response)
			data, _ := json.Marshal(response)
			io.WriteString(w, string(data[:]))
		} else {
			io.WriteString(w, err.Error())
		}
	}
}

func (cs *RestServer) handleCommand(w http.ResponseWriter, r *http.Request) {
	if cs.commandHandler == nil {
		io.WriteString(w, "error")
	}

	body, _ := ioutil.ReadAll(r.Body)
	var request *models.Request
	var response *models.Response
	if err := json.Unmarshal(body, &request); err == nil {
		cs.requestHandler.Handle(request, response)
		data, _ := json.Marshal(response)
		io.WriteString(w, string(data[:]))
	} else {
		io.WriteString(w, err.Error())
	}
}

func (cs *RestServer) handleRealtimeData(w http.ResponseWriter, r *http.Request) {
	if cs.realtimeDataHandler == nil {
		io.WriteString(w, "error")
	}

	body, _ := ioutil.ReadAll(r.Body)
	var request *models.Request
	var response *models.Response
	if err := json.Unmarshal(body, &request); err == nil {
		cs.requestHandler.Handle(request, response)
		data, _ := json.Marshal(response)
		io.WriteString(w, string(data[:]))
	} else {
		io.WriteString(w, err.Error())
	}
}

func (cs *RestServer) handleAlarm(w http.ResponseWriter, r *http.Request) {
	if cs.alarmHandler == nil {
		io.WriteString(w, "error")
	}

	body, _ := ioutil.ReadAll(r.Body)
	var request *models.Request
	var response *models.Response
	if err := json.Unmarshal(body, &request); err == nil {
		cs.requestHandler.Handle(request, response)
		data, _ := json.Marshal(response)
		io.WriteString(w, string(data[:]))
	} else {
		io.WriteString(w, err.Error())
	}
}

func (cs *RestServer) handleFault(w http.ResponseWriter, r *http.Request) {
	if cs.faulttHandler == nil {
		io.WriteString(w, "error")
	}

	body, _ := ioutil.ReadAll(r.Body)
	var request *models.Request
	var response *models.Response
	if err := json.Unmarshal(body, &request); err == nil {
		cs.requestHandler.Handle(request, response)
		data, _ := json.Marshal(response)
		io.WriteString(w, string(data[:]))
	} else {
		io.WriteString(w, err.Error())
	}
}

func (cs *RestServer) handleResult(w http.ResponseWriter, r *http.Request) {
	if cs.resultHandler == nil {
		io.WriteString(w, "error")
	}

	body, _ := ioutil.ReadAll(r.Body)
	var request *models.Request
	var response *models.Response
	if err := json.Unmarshal(body, &request); err == nil {
		cs.requestHandler.Handle(request, response)
		data, _ := json.Marshal(response)
		io.WriteString(w, string(data[:]))
	} else {
		io.WriteString(w, err.Error())
	}
}
