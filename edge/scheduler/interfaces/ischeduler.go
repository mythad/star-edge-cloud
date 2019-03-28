package interfaces

import (
	"star-edge-cloud/edge/models"
	"time"
)

// IScheduler -
type IScheduler interface {
	// AddSchedulerTask - 注册一个调度任务
	AddSchedulerTask(task *models.SchedulerTask) error
	// RemoveSchedulerTask - 删除一个调度任务
	RemoveSchedulerTask(ID string) error
	// UpdateSchedulerTask -
	UpdateSchedulerTask(task *models.SchedulerTask) error
	// QueryAllSchedulerTask - 查询所有调度任务
	QueryAllSchedulerTask() (tasks []models.SchedulerTask, err error)
	// Load - 加载任务
	Load()
	// AddTaskToQueue - 加入队列
	AddTaskToQueue(task *models.SchedulerTask)
	// Execute - 执行调度任务
	Execute(task *models.SchedulerTask)
	// GetNextTime - 执行调度任务
	GetNextTime(task *models.SchedulerTask) time.Time
	// InitDB - 初始化调度服务数据库
	InitDB(path string)
}
