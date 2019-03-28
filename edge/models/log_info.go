package models

// LogInfo - 消息
type LogInfo struct {
	ID      string
	Message string
	Level   int
	Time    int64
}

// &model.LogInfo{ID: common.GetGUID(), Message: "", Level: 0, Time: time.Now()}
