package interfaces

import "star_cloud/edge/models"

// IClient -- edge和cloud之间通信的客户端接口
type IClient interface {
	PostRequest(addr string, request *models.Request) (*models.Response, error)
	PostCommand(addr string, command *models.Command) (*models.Response, error)
	PostRealtimeData(addr string, RealtimeData *models.RealtimeData) (*models.Response, error)
	PostAlarm(addr string, alarm *models.Alarm) (*models.Response, error)
	PostFault(addr string, fault *models.Fault) (*models.Response, error)
	PostResult(addr string, result *models.Result) (*models.Response, error)
}
