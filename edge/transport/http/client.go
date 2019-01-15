package http

import (
	"encoding/json"
	"star_cloud/edge/models"
)

// RestClient -
type RestClient struct{}

// PostRequest -
func (rc *RestClient) PostRequest(addr string, request *models.Request) (rsp *models.Response, err error) {
	data, err := json.Marshal(request)
	if err != nil {
		rsp.Status = "error"
		return
	}

	r, err := PostData(addr, data)
	if err != nil {
		rsp.Status = "error"
		return
	}

	if err = json.Unmarshal(r, &rsp); err != nil {
		rsp.Status = "error"
		return
	}

	return
}

// PostCommand -
func (rc *RestClient) PostCommand(addr string, command *models.Command) (rsp *models.Response, err error) {
	data, err := json.Marshal(command)
	if err != nil {
		return nil, err
	}
	r, err := PostData(addr, data)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(r, &rsp); err != nil {
		return nil, err
	}

	return
}

// PostRealtimeData -
func (rc *RestClient) PostRealtimeData(addr string, data *models.RealtimeData) (rsp *models.Response, err error) {
	d, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	r, err := PostData(addr, d)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(r, &rsp); err != nil {
		return nil, err
	}

	return
}

// PostAlarm -
func (rc *RestClient) PostAlarm(addr string, alarm *models.Alarm) (rsp *models.Response, err error) {
	data, err := json.Marshal(alarm)
	if err != nil {
		return nil, err
	}
	r, err := PostData(addr, data)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(r, &rsp); err != nil {
		return nil, err
	}

	return
}

// PostFault -
func (rc *RestClient) PostFault(addr string, fault *models.Fault) (rsp *models.Response, err error) {
	data, err := json.Marshal(fault)
	if err != nil {
		return nil, err
	}
	r, err := PostData(addr, data)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(r, &rsp); err != nil {
		return nil, err
	}

	return
}

// PostResult -
func (rc *RestClient) PostResult(addr string, result *models.Result) (rsp *models.Response, err error) {
	data, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	r, err := PostData(addr, data)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(r, &rsp); err != nil {
		return nil, err
	}

	return
}
