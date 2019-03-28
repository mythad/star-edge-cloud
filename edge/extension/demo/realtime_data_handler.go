package demo

import (
	"log"
	"star-edge-cloud/edge/extension/interfaces"
	"star-edge-cloud/edge/models"
	"star-edge-cloud/edge/models/constant"
	"star-edge-cloud/edge/transport/client"
	"star-edge-cloud/edge/transport/client/configuration"
	"star-edge-cloud/edge/transport/coding"
	"star-edge-cloud/edge/transport/coding/json"
	"star-edge-cloud/edge/utils/common"
	"time"
)

// RealtimeDataHandler 接收到数据，调用回调方法
type RealtimeDataHandler struct {
	decoder coding.IDecoder
	encoder coding.IEncoder
	client  *client.Client
	a1      interfaces.IAlgorithm
}

// SetDecoder -
func (it *RealtimeDataHandler) SetDecoder(de coding.IDecoder) {
	it.decoder = de
}

// SetEncoder -
func (it *RealtimeDataHandler) SetEncoder(en coding.IEncoder) {
	it.encoder = en
}

// SetAlgorithm -
func (it *RealtimeDataHandler) SetAlgorithm(algorithm interfaces.IAlgorithm) {
	it.a1 = algorithm
}

// OnReceive -
func (it *RealtimeDataHandler) OnReceive(v interface{}) (response []byte, err error) {
	// 初始化客户端
	it.client = client.New()
	c := it.client.GetChannel(&configuration.ChannelConfig{
		Timeout:  10 * time.Second,
		Keep:     true,
		DataType: constant.RealtimeData,
		Protocol: constant.HTTP,
		Cache:    false,
	})
	c.SetDecoder(&json.JSONDecoder{})
	c.SetEncoder(&json.JSONEncoder{})

	var i int
	it.decoder.Decode(v.([]byte), &i)
	// extlog.WriteLog("收到实时数据:" + request.ID)
	if data, err := it.a1.Calculate(i); err == nil {
		response = it.encoder.Encode(&models.Response{Status: "true"})
		alarm := &models.Alarm{ID: common.GetGUID(), Data: data}
		log.Println(alarm)
		// for e := it.Listeners.Front(); e != nil; e = e.Next() {
		// 	c.Send(e.Value.(string)+"/api/transport/alarm", alarm)
		// 	// extlog.WriteLog("发送报警数据:" + e.Value.(string) + "/api/transport/alarm" + alarm.ID)
		// }
		return nil, nil
	}

	response = it.encoder.Encode(&models.Response{Status: "false"})
	return nil, nil
}

type rdecoder struct {
}

func (it *rdecoder) Decode(data []byte, v interface{}) {
	common.StringToInt(string(data[:]))
}

type rencoder struct{}

func (it *rencoder) Encode(v interface{}) []byte {
	return nil
}
