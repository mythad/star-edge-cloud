package models

// Alarm - 报警信息
type Alarm struct {
	ID          string
	Type        string
	TriggerTime string
	SendingTime string
	Level       string
	DeviceID    string
	Data        []byte
}
