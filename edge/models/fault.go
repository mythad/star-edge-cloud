package models

// Fault - 异常数据
type Fault struct {
	ID      string
	Type    string
	Message string
	Reason  string
	Time    int64
	Data    []byte
}
