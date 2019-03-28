package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"star-edge-cloud/edge/transport/client/configuration"
	"star-edge-cloud/edge/transport/coding"
)

// Channel -
type Channel struct {
	config  *configuration.ChannelConfig
	client  *http.Client
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
func (it *Channel) Send(addr string, v interface{}, properties map[string]interface{}) (interface{}, error) {
	if it.client == nil {
		it.client = &http.Client{Timeout: it.config.Timeout}
	}
	if it.encoder == nil {
		return nil, fmt.Errorf("没有编码器，无法发送数据。")
	}
	request := it.encoder.Encode(v)

	headerType := "text/plain"
	if properties != nil && properties["headerType"] != nil {
		headerType = properties["headerType"].(string)
	}
	response, err := it.post(addr, request, headerType)

	if it.decoder == nil {
		return response, fmt.Errorf("没有解码器。")
	}
	var result interface{}
	it.decoder.Decode(response, result)
	return result, err
}

// Close -
// func (it *Channel) Close() {
// }

// application/octet-stream, application/json, text/xml , text/plain
func (it *Channel) post(url string, data []byte, headerType string) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", headerType)

	resp, err := it.client.Do(req)
	if err != nil {
		// panic(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			// handle error
			return nil, err
		}
		// fmt.Println(string(body))
		return body, nil
	}

	return nil, fmt.Errorf(string(resp.StatusCode))
}
