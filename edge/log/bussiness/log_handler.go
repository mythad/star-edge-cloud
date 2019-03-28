package bussiness

import (
	"star-edge-cloud/edge/log/interfaces"
	"star-edge-cloud/edge/models"
	"star-edge-cloud/edge/transport/coding"
)

// LogHandler 接收到数据，调用回调方法
type LogHandler struct {
	decoder coding.IDecoder
	encoder coding.IEncoder
	logger  interfaces.ILogger
}

// SetDecoder -
func (it *LogHandler) SetDecoder(de coding.IDecoder) {
	it.decoder = de
}

// SetEncoder -
func (it *LogHandler) SetEncoder(en coding.IEncoder) {
	it.encoder = en
}

// SetLogger -
func (it *LogHandler) SetLogger(logger interfaces.ILogger) {
	it.logger = logger
}

// OnReceive -
func (it *LogHandler) OnReceive(v interface{}) (response []byte, err error) {
	loginfo := v.(models.LogInfo)

	it.logger.Write(&loginfo)
	response = it.encoder.Encode(&models.Response{Status: "true"})
	return
}
