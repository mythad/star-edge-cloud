package bussiness

import (
	"star-edge-cloud/edge/models"
	"star-edge-cloud/edge/transport/coding"
)

// AlarmHandler 接收到数据，调用回调方法
type AlarmHandler struct {
	decoder coding.IDecoder
	encoder coding.IEncoder
}

// SetDecoder -
func (it *AlarmHandler) SetDecoder(de coding.IDecoder) {
	it.decoder = de
}

// SetEncoder -
func (it *AlarmHandler) SetEncoder(en coding.IEncoder) {
	it.encoder = en
}

// OnReceive -
func (it *AlarmHandler) OnReceive(v interface{}) ([]byte, error) {
	result := v.(models.Alarm)
	if err := store.Save(result.ID, it.encoder.Encode(result)); err != nil {
		return it.encoder.Encode(&models.Response{Status: "false"}), err
	}
	return it.encoder.Encode(&models.Response{Status: "true"}), nil
}
