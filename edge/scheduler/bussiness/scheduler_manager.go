package bussiness

import (
	"database/sql"
	"log"
	"os"
	"star-edge-cloud/edge/models"
	"star-edge-cloud/edge/models/constant"
	"star-edge-cloud/edge/scheduler/interfaces"
	"star-edge-cloud/edge/transport/client"
	"star-edge-cloud/edge/transport/client/configuration"
	"star-edge-cloud/edge/transport/coding/json"
	"star-edge-cloud/edge/utils/common"
	"time"
	// 导入驱动
	_ "github.com/mattn/go-sqlite3"
)

// SchedulerManager - 调度管理
type SchedulerManager struct {
	HighPriorityQueue *Queue //存放10分钟内需要执行的任务
	LowPriorityQueue  *Queue //存放超过10分钟内需要执行的任务
	MetadataDBPath    string
	Client            *client.Client
}

// New -
func New() interfaces.IScheduler {
	manager := SchedulerManager{}
	manager.HighPriorityQueue = NewQueue()
	manager.LowPriorityQueue = NewQueue()

	return &manager
}

// AddSchedulerTask - 注册一个调度任务
func (it *SchedulerManager) AddSchedulerTask(task *models.SchedulerTask) error {
	it.AddTaskToQueue(task)
	sqlStmt1 := `
	INSERT INTO task (id, name, address, topic, frequency, offset, user_id, user_data, is_available)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);
	`
	db, err := sql.Open("sqlite3", it.MetadataDBPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if _, err = db.Exec(sqlStmt1,
		task.ID,
		task.Name,
		task.Address,
		task.Topic,
		task.ExecutingFrequency,
		task.Offset,
		task.UserID,
		task.UserData,
		task.IsAvailable); err != nil {
		log.Printf("%q: %s\n", err, sqlStmt1)
	}
	return nil
}

// RemoveSchedulerTask - 删除一个调度任务
func (it *SchedulerManager) RemoveSchedulerTask(ID string) error {
	//todo:update task status
	sqlStmt1 := `
	DELETE FROM task WHERE id=?;
	`
	db, err := sql.Open("sqlite3", it.MetadataDBPath)
	if err != nil {
		return err
	}
	defer db.Close()
	if _, err := db.Exec(sqlStmt1, ID); err != nil {
		return err
	}
	return nil
}

// UpdateSchedulerTask -
func (it *SchedulerManager) UpdateSchedulerTask(task *models.SchedulerTask) error {
	sqlStmt1 := `
	UPDATE task
	SET name=?, address=?, topic=?, frequency=?, offset=?, user_id=?, user_data=?, is_available=?
	WHERE id=?;
	`
	db, err := sql.Open("sqlite3", it.MetadataDBPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if _, err = db.Exec(sqlStmt1,
		task.Name,
		task.Address,
		task.Topic,
		task.ExecutingFrequency,
		task.Offset,
		task.UserID,
		task.UserData,
		task.IsAvailable,
		task.ID); err != nil {
		log.Printf("%q: %s\n", err, sqlStmt1)
	}
	return nil
}

// QueryAllSchedulerTask - 查询所有调度任务
func (it *SchedulerManager) QueryAllSchedulerTask() (tasks []models.SchedulerTask, err error) {
	sqlStmt1 := `
	SELECT id, name, address, topic, frequency, offset, user_id, user_data, is_available
	FROM task where is_available=true
	`
	db, err := sql.Open("sqlite3", it.MetadataDBPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var rows *sql.Rows
	// defer rows.Close()
	if rows, err = db.Query(sqlStmt1); err == nil {
		for rows.Next() {
			task := models.SchedulerTask{}
			rows.Scan(&task.ID,
				&task.Name,
				&task.Address,
				&task.Topic,
				&task.ExecutingFrequency,
				&task.Offset,
				&task.UserID,
				&task.UserData,
				&task.IsAvailable)
			tasks = append(tasks, task)
		}
		rows.Close()
	}

	return tasks, err
}

// Load - 加载任务
func (it *SchedulerManager) Load() {
	tasks, err := it.QueryAllSchedulerTask()
	if err == nil {
		return
	}

	for _, item := range tasks {
		it.AddTaskToQueue(&item)
	}
}

// AddTaskToQueue - 加入队列
func (it *SchedulerManager) AddTaskToQueue(task *models.SchedulerTask) {
	d := it.GetNextTime(task).Sub(time.Now())
	if d > 0*time.Second && d <= 10*time.Minute {
		it.HighPriorityQueue.LastScanTime = time.Now()
		it.HighPriorityQueue.Push(task)
	} else {
		it.LowPriorityQueue.Push(task)
		it.LowPriorityQueue.LastScanTime = time.Now()
	}
}

// Execute - 执行调度任务
func (it *SchedulerManager) Execute(task *models.SchedulerTask) {
	request := &models.Request{
		ID:      common.GetGUID(),
		Content: []byte(task.UserData),
	}
	// 初始化客户端
	it.Client = client.New()
	c := it.Client.GetChannel(&configuration.ChannelConfig{
		Timeout:  10 * time.Second,
		Keep:     true,
		DataType: constant.RealtimeData,
		Protocol: constant.HTTP,
		Cache:    false,
	})
	c.SetDecoder(&json.JSONDecoder{})
	c.SetEncoder(&json.JSONEncoder{})

	c.Send(task.Address, request, map[string]interface{}{"headerType": "text/plain"})
	// log.Println(fmt.Printf("任务：%[1]s被执行", it.Name))
}

// GetNextTime - 执行调度任务
func (it *SchedulerManager) GetNextTime(task *models.SchedulerTask) time.Time {
	now := time.Now()
	switch task.ExecutingFrequency {
	case constant.Once:
		t := task.Offset * int64(time.Millisecond)
		return now.Add(time.Duration(t))
	case constant.Minute:
		t := task.Offset * int64(time.Millisecond)
		periodTime := now.Truncate(time.Minute * 24).Add(time.Duration(t))
		if periodTime.Sub(now) <= 0 {
			return periodTime.Add(time.Minute)
		}
		return periodTime
	case constant.Hour:
		t := task.Offset * int64(time.Millisecond)
		periodTime := now.Truncate(time.Hour * 24).Add(time.Duration(t))
		if periodTime.Sub(now) <= 0 {
			return periodTime.Add(time.Hour)
		}
		return periodTime
	case constant.Day:
		year, month, day := now.Date()
		t := task.Offset * int64(time.Millisecond)
		periodTime := time.Date(year, month, day, 0, 0, 0, 0, time.Local).Add(time.Duration(t))
		if periodTime.Sub(now) <= 0 {
			return periodTime.Add(24 * time.Hour)
		}
		return periodTime
	case constant.Week:
		weekday := now.Weekday()
		t := task.Offset * int64(time.Nanosecond)
		t1 := time.Duration(weekday)
		periodTime := now.Truncate(time.Hour * 24).Add(-t1*24*time.Hour + time.Duration(t))
		if periodTime.Sub(now) <= 0 {
			return periodTime.Add(7 * 24 * time.Hour)
		}
		return periodTime
	case constant.Month:
		year, month, _ := now.Date()
		t := task.Offset * int64(time.Nanosecond)
		periodTime := time.Date(year, month, 0, 0, 0, 0, 0, time.Local).Add(time.Duration(t))
		if periodTime.Sub(now) <= 0 {
			return time.Date(year, month+1, 0, 0, 0, 0, 0, time.Local).Add(time.Duration(t))
		}
		return periodTime
	default:
		return time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)
	}
}

// InitDB - 初始化调度服务数据库
func (it *SchedulerManager) InitDB(path string) {
	if _, err := os.Stat(path); err == nil {
		log.Println("调度数据库文件：" + path)
		return
	}

	sqlStmt2 := `
	create table task (id string not null primary key, name text, address string, topic string, frequency string, offset BIGINT, user_id string, user_data string, is_available BOOLEAN);
	`

	db1, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}
	defer db1.Close()

	if _, err = db1.Exec(sqlStmt2); err != nil {
		log.Printf("%q: %s\n", err, sqlStmt2)
	}

}
