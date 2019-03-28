package interfaces

import (
	"star-edge-cloud/edge/models"
	"time"
)

// ILogger -
type ILogger interface {
	// Write - 写日志
	Write(info *models.LogInfo)

	// QueryLevel - 查询日志
	QueryLevel(start time.Time, end time.Time, level int) (infos []models.LogInfo, err error)

	// QueryMinLevel - 查询日志
	QueryMinLevel(start time.Time, end time.Time, level int) (infos []models.LogInfo, err error)

	// QueryTop - 查询日志
	QueryTop(top int) (infos []models.LogInfo, err error)

	// Delete -
	Delete(id string) error

	// Initialize - 初始化日志数据库
	Initialize()
}
