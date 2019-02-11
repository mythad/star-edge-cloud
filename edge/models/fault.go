package models

// Fault - 异常数据
type Fault struct {
	ID      string
	Type    string
	Message string
	Reason  string
	Time    string
	Data    []byte
}
