package devlog

import (
	"star-edge-cloud/edge/log/http"
	"star-edge-cloud/edge/models"
	"star-edge-cloud/edge/utils/common"
	"time"
)

var loggerClient http.HttpLogClient

// SetLogServiceURL - 日志服务地址
func SetLogServiceURL(url string) {
	loggerClient.BaseAddress = url
}

// WriteLog - 写日志
func WriteLog(message string) {
	info := &models.LogInfo{ID: common.GetGUID(), Message: "ext:" + message, Level: 0, Time: time.Now()}
	loggerClient.Write(info)
}
