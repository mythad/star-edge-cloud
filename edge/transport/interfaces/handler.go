package interfaces

import "star_cloud/edge/models"

// IRequestHandler 接收到数据，调用回调方法
type IRequestHandler interface {
	Handle(request *models.Request, response *models.Response) error
}

// ICommandHandler -
type ICommandHandler interface {
	Handle(command *models.Command, response *models.Response) error
}

// IRealtimeDataHandler -
type IRealtimeDataHandler interface {
	Handle(data *models.RealtimeData, response *models.Response) error
}

// IAlarmHandler -
type IAlarmHandler interface {
	Handle(alart *models.Alarm, response *models.Response) error
}

// IFaultHandler -
type IFaultHandler interface {
	Handle(fault *models.Fault, response *models.Response) error
}

// IResultHandler -
type IResultHandler interface {
	Handle(result *models.Result, response *models.Response) error
}
