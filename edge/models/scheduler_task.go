package models

import "star-edge-cloud/edge/models/constant"

// SchedulerTask - 调度任务
type SchedulerTask struct {
	ID      string
	Name    string
	Address string
	Topic   string
	// once = 1, minute = 2, hour = 4, day = 8, week = 16, month = 32, year = 64
	ExecutingFrequency constant.TaskFrequency
	// once, minute, hour, day, week, month + offset,单位微妙（10~-9秒）
	Offset      int64
	UserID      string
	UserData    string //用户数据
	IsAvailable bool
}
