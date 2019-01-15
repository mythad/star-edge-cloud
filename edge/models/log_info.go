package models

import "time"

// LogInfo - 消息
type LogInfo struct {
	ID      string
	Message string
	Level   int
	Time    time.Time
}

// &model.LogInfo{ID: common.GetGUID(), Message: "", Level: 0, Time: time.Now()}
