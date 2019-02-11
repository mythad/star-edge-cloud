package http

import (
	"encoding/json"
	"star-edge-cloud/edge/models"
	"star-edge-cloud/edge/transport/http"
	"time"
)

// HttpLogClient -
type HttpLogClient struct {
	BaseAddress string
}

// Write - 写日志
func (hlc *HttpLogClient) Write(info *models.LogInfo) {
	data, _ := json.Marshal(&info)
	http.PostData(hlc.BaseAddress+"/api/log/write", data)
}

// QueryLevel - 查询日志
func (hlc *HttpLogClient) QueryLevel(start time.Time, end time.Time, level int) (infos []models.LogInfo, err error) {
	return nil, nil
}

// QueryMinLevel - 查询日志
func (hlc *HttpLogClient) QueryMinLevel(start time.Time, end time.Time, level int) (infos []models.LogInfo, err error) {
	return nil, nil
}

// QueryTop - 查询日志
func (hlc *HttpLogClient) QueryTop(top int) (infos []models.LogInfo, err error) {
	return nil, nil
}

// Delete -
func (hlc *HttpLogClient) Delete(id string) error {
	return nil
}
