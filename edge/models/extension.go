package models

// Extension - 服务
type Extension struct {
	ID           string
	Name         string
	FileName     string
	Describe     string
	RegistryTime string
	Type         string
	Other        string
	// Protocol -设备通讯协议
	// 目前支持http,mqtt,amqp...
	Protocol string
	//配置信息
	Conf          string
	Status        int
	Listeners     string
	ServerAddress string
	LogBaseURL    string
}
