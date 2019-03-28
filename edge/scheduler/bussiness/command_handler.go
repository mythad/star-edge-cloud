package bussiness

import (
	"encoding/json"
	"star-edge-cloud/edge/models"
	"star-edge-cloud/edge/transport/client"
	"star-edge-cloud/edge/transport/coding"
)

// CommandHandler 接收到数据，调用回调方法
type CommandHandler struct {
	client  *client.Client
	Manager *SchedulerManager
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
	command := v.(models.Command)

	switch command.Type {
	case "remove":
		id := string(command.Data)
		if err := it.Manager.RemoveSchedulerTask(id); err == nil {
			response.Status = "success"
		} else {
			response.Status = err.Error()
		}
	case "list":
		if collection, err := it.Manager.QueryAllSchedulerTask(); err == nil {
			response.Status = "success"
			response.Message, _ = json.Marshal(collection)
		} else {
			response.Status = err.Error()
		}
	}

	return nil, nil
}
