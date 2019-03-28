package models

// Request -
type Request struct {
	ID         string
	Status     string
	CreateTime int64
	Content    []byte
	// Type -
	// realtime_data--实时数据
	// fault --错误
	// command -- 命令
	// result -- 数据处理结果
	// scheduler_task -- 调度任务
	// rules -- 规则
	Type string
}
