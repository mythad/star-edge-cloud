package demo

import (
	"container/list"
	"star-edge-cloud/edge/models"
	"star-edge-cloud/edge/transport/coding"
)

// CommandHandler 接收到数据，调用回调方法
type CommandHandler struct {
	decoder coding.IDecoder
	encoder coding.IEncoder
}

// SetDecoder -
func (it *CommandHandler) SetDecoder(de coding.IDecoder) {
	it.decoder = de
}

// SetEncoder -
func (it *CommandHandler) SetEncoder(en coding.IEncoder) {
	it.encoder = en
}

// OnReceive -
func (it *CommandHandler) OnReceive(v interface{}) ([]byte, error) {
	response := &models.Response{}
	// if request.Type == "exit" {
	// 	response.Message = []byte("success")
	// 	interfaces.ChStopTask <- 0
	// }
	request := v.(models.Command)
	if request.Type == "addlistener" {
		// 如果没有监听
		if listeners == nil {
			listeners = list.New()
		}
		listeners.PushBack(string(request.Data[:]))
		response.Status = "sucess"
	}

	return it.encoder.Encode(response), nil
}
