package connectionpool

import (
	"net/http"
	"time"

	"github.com/smallnest/rpcx/client"
	"pack.ag/amqp"
)

// Pool -
type Pool struct {
	InitCap     int
	MaxCap      int
	IdleTimeout time.Duration
	WaitTimeout time.Duration
	clients     map[string]interface{}
}

// GetClient -
func (p *Pool) GetClient(key string, clientType string) interface{} {
	if p.clients == nil {
		p.clients = make(map[string]interface{})
	}

	switch clientType {
	case "http":
		return &http.Client{Timeout: 10 * time.Second}
	case "rpcx":
		d := client.NewPeer2PeerDiscovery("", "")
		return client.NewXClient("transport", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	case "amqp":
		client, _ := amqp.Dial("amqps://my-namespace.servicebus.windows.net",
			amqp.ConnSASLPlain("access-key-name", "access-key"),
		)
		return client
	case "mqtt":
		return &http.Client{Timeout: 10 * time.Second}
	default:
		return &http.Client{Timeout: 10 * time.Second}
	}

}
