package demo

import (
	"container/list"
	"errors"
	"log"
	"star-edge-cloud/edge/models/constant"
	"star-edge-cloud/edge/transport/client"
	cc "star-edge-cloud/edge/transport/client/configuration"
	"star-edge-cloud/edge/transport/coding"
	"star-edge-cloud/edge/transport/coding/json"
	"time"
)

// DriverHandler -
type DriverHandler struct {
	client *client.Client

	encoder coding.IEncoder
	decoder coding.IDecoder
}

// OnError -
func (it *DriverHandler) OnError(err error) {

}

// SetEncoder -
func (it *DriverHandler) SetEncoder(encoder coding.IEncoder) {
	it.encoder = encoder
}

// SetDecoder -
func (it *DriverHandler) SetDecoder(decoder coding.IDecoder) {
	it.decoder = decoder
}

// OnReceive -
func (it *DriverHandler) OnReceive(v interface{}) ([]byte, error) {
	log.Println(v)
	// 初始化客户端
	if it.client == nil {
		it.client = client.New()
	}
	c := it.client.GetChannel(&cc.ChannelConfig{
		Timeout:  10 * time.Second,
		Keep:     true,
		DataType: constant.RealtimeData,
		Protocol: constant.HTTP,
		Cache:    false,
	})
	c.SetDecoder(&json.JSONDecoder{})
	c.SetEncoder(&json.JSONEncoder{})

	// 如果没有监听
	if listeners == nil {
		listeners = list.New()
		return nil, errors.New("未添加结果监听")
	}

	// 向监听者发送采集到的数据
	for e := listeners.Front(); e != nil; e = e.Next() {
		// result, err := c.Send(e.Value.(string)+"/api/realtime_data/load", v)
		go c.Send(e.Value.(string), v, map[string]interface{}{"headerType": "text/plain"})
	}

	return nil, nil
}

// OnInitialized  -
func (it *DriverHandler) OnInitialized() {

}

// OnBegin -
func (it *DriverHandler) OnBegin() {

}

// OnEnd -
func (it *DriverHandler) OnEnd() {

}

// AddListener -
func (it *DriverHandler) AddListener(url string) {

}
