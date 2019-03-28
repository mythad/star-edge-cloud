package bussiness

import (
	"star-edge-cloud/edge/models"
	"star-edge-cloud/edge/scheduler/interfaces"
	"star-edge-cloud/edge/transport/coding"
)

// SchedulerTaskHandler 接收到数据，调用回调方法
type SchedulerTaskHandler struct {
	// Client  interfaces.IClient
	manager interfaces.IScheduler
	decoder coding.IDecoder
	encoder coding.IEncoder
}

// SetDecoder -
func (it *SchedulerTaskHandler) SetDecoder(de coding.IDecoder) {
	it.decoder = de
}

// SetEncoder -
func (it *SchedulerTaskHandler) SetEncoder(en coding.IEncoder) {
	it.encoder = en
}

// SetScheduler -
func (it *SchedulerTaskHandler) SetScheduler(scheduler interfaces.IScheduler) {
	it.manager = scheduler
}

// OnReceive -
func (it *SchedulerTaskHandler) OnReceive(v interface{}) ([]byte, error) {
	task := v.(models.SchedulerTask)

	err := it.manager.AddSchedulerTask(&task)
	return nil, err
}
