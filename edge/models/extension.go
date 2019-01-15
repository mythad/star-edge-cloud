package models

// Extension - 服务
type Extension struct {
	ID            string
	Name          string
	FileName      string
	Describe      string
	RegistryTime  string
	Type          string
	Other         string
	Protocol      string //设备通讯协议
	Conf          string //配置信息
	Status        int
	Listeners     string
	ServerAddress string
	LogBaseURL    string
}
