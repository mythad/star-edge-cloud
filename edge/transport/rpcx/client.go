package rpcx

import (
	"context"
	"log"
	"star-edge-cloud/edge/models"
	"strings"

	"github.com/smallnest/rpcx/client"
)

// RpcxClient -
type RpcxClient struct{}

// // Send -
// func Send(addr string, path string, data []byte) {
// 	d := client.NewPeer2PeerDiscovery(addr, "")
// 	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
// 	defer xclient.Close()

// 	reply := models.Reply{}
// 	err := xclient.Call(context.Background(), "Receive", data, reply)
// 	if err != nil {
// 		log.Fatalf("调用失败: %v", err)
// 	}

// }

// PostRequest -
func (rc *RpcxClient) PostRequest(addr string, request *models.Request) (rsp *models.Response, err error) {
	d := client.NewPeer2PeerDiscovery(addr, "")
	path := strings.Split(addr, "/")
	xclient := client.NewXClient("transport", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	err = xclient.Call(context.Background(), path[1], request, rsp)
	if err != nil {
		log.Fatalf("调用失败: %v", err)
	}

	return
}

// PostCommand -
func (rc *RpcxClient) PostCommand(addr string, command *models.Command) (rsp *models.Response, err error) {
	d := client.NewPeer2PeerDiscovery(addr, "")
	path := strings.Split(addr, "/")
	xclient := client.NewXClient("transport", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	err = xclient.Call(context.Background(), path[1], command, rsp)
	if err != nil {
		log.Fatalf("调用失败: %v", err)
	}

	return
}

// PostRealtimeData -
func (rc *RpcxClient) PostRealtimeData(addr string, data *models.RealtimeData) (rsp *models.Response, err error) {
	d := client.NewPeer2PeerDiscovery(addr, "")
	path := strings.Split(addr, "/")
	xclient := client.NewXClient("transport", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	err = xclient.Call(context.Background(), path[1], data, rsp)
	if err != nil {
		log.Fatalf("调用失败: %v", err)
	}

	return
}

// PostAlarm -
func (rc *RpcxClient) PostAlarm(addr string, alarm *models.Alarm) (rsp *models.Response, err error) {
	d := client.NewPeer2PeerDiscovery(addr, "")
	path := strings.Split(addr, "/")
	xclient := client.NewXClient("transport", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	err = xclient.Call(context.Background(), path[1], alarm, rsp)
	if err != nil {
		log.Fatalf("调用失败: %v", err)
	}

	return
}

// PostFault -
func (rc *RpcxClient) PostFault(addr string, fault *models.Fault) (rsp *models.Response, err error) {
	d := client.NewPeer2PeerDiscovery(addr, "")
	path := strings.Split(addr, "/")
	xclient := client.NewXClient("transport", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	err = xclient.Call(context.Background(), path[1], fault, rsp)
	if err != nil {
		log.Fatalf("调用失败: %v", err)
	}

	return
}

// PostResult -
func (rc *RpcxClient) PostResult(addr string, result *models.Result) (rsp *models.Response, err error) {
	d := client.NewPeer2PeerDiscovery(addr, "")
	path := strings.Split(addr, "/")
	xclient := client.NewXClient("transport", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	err = xclient.Call(context.Background(), path[1], result, rsp)
	if err != nil {
		log.Fatalf("调用失败: %v", err)
	}

	return
}
