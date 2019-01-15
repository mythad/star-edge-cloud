package nbiot

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

type OutputStream struct {
	ws *websocket.Conn
}

type OutputDataMessage struct {
	Device   Device `json:"device"`
	Payload  []byte `json:"payload"`
	Received int64  `json:"received"`
}

func (c *Client) CollectionOutputStream(collectionID string) (*OutputStream, error) {
	return c.outputStream(fmt.Sprintf("/collections/%s", collectionID))
}

func (c *Client) DeviceOutputStream(collectionID, deviceID string) (*OutputStream, error) {
	return c.outputStream(fmt.Sprintf("/collections/%s/devices/%s", collectionID, deviceID))
}

func (c *Client) outputStream(path string) (*OutputStream, error) {
	url, err := url.Parse(c.addr)
	if err != nil {
		return nil, err
	}

	scheme := "wss"
	if url.Scheme == "http" {
		scheme = "ws"
	}

	urlStr := fmt.Sprintf("%s://%s%s/from", scheme, url.Host, path)

	header := http.Header{}
	header.Add("X-API-Token", c.token)

	dialer := websocket.Dialer{}
	ws, _, err := dialer.Dial(urlStr, header)
	if err != nil {
		return nil, err
	}

	return &OutputStream{ws}, nil
}

func (s *OutputStream) Recv() (OutputDataMessage, error) {
	for {
		var msg struct {
			Type string `json:"type"`
			OutputDataMessage
		}
		err := s.ws.ReadJSON(&msg)
		if err != nil {
			return OutputDataMessage{}, err
		}

		if msg.Type == "data" {
			return msg.OutputDataMessage, nil
		}
	}
}

func (s *OutputStream) Close() {
	s.ws.Close()
}
