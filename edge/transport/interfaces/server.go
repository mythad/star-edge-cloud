package interfaces

// IServer -- edge和cloud之间通信的服务接口
type IServer interface {
	SetConfig(string)
	Start() error
	Stop() error
	SetRequestHandler(IRequestHandler)
	SetCommandHandler(ICommandHandler)
	SetRealtimeDataHandler(IRealtimeDataHandler)
	SetAlarmHandler(IAlarmHandler)
	SetFaultHandler(IFaultHandler)
	SetResultHandler(IResultHandler)
}
