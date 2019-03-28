package constant

// Protocol -系统支持的协议
type Protocol int

// 支持的协议类型
const (
	HTTP Protocol = iota
	RPCX
	AMQP
	MQTT
	MODBUS
)

// DataType -系统支持的数据类型
type DataType int

const (
	// RealtimeData -实时数据，传感器上传或设备采集来的数据
	RealtimeData DataType = iota
	// Command -上下行命令
	Command
	// State - 状态数据
	State
	// Alarm - 报警数据
	Alarm
	// SchedulerTask -
	SchedulerTask
	// Result  -结果数据
	Result
	// LogInfo - log
	LogInfo
	// Unkown -未知
	Unkown
)

// TaskFrequency -调度任务频率
type TaskFrequency int

const (
	// Once -
	Once TaskFrequency = iota
	// Second -上下行命令
	Second
	// Minute - 状态数据
	Minute
	// Hour -
	Hour
	// Day - 报警数据
	Day
	// Week -未知
	Week
	// Month -
	Month
	// // Year -
	// Year
)

// String -
func (it DataType) String() string {
	switch it {
	case RealtimeData:
		return "realtime_data"
	case Command:
		return "command"
	case State:
		return "state"
	case Alarm:
		return "alarm"
	default:
		return "unknow"
	}
}

// Convert -
func Convert(str string) DataType {
	switch str {
	case "realtime_data":
		return RealtimeData
	case "command":
		return Command
	case "state":
		return State
	case "alarm":
		return Alarm
	default:
		return Unkown
	}
}
