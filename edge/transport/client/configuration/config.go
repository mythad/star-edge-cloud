package configuration

import (
	"star-edge-cloud/edge/models/constant"
	"time"
)

// ChannelConfig -
type ChannelConfig struct {
	Timeout  time.Duration
	Keep     bool
	DataType constant.DataType
	Protocol constant.Protocol
	Cache    bool
}
