package models

// Device - 设备服务实体
type Device struct {
	ID           string
	Name         string
	FileName     string
	Describe     string
	RegistryTime string
	Type         string
	Other        string
	//设备通讯协议
	Protocol string
	//配置信息
	Conf          string
	Status        int
	Listeners     string
	ServerAddress string
	LogBaseURL    string
}
