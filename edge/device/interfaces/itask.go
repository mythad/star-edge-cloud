package interfaces

import (
	uinterface "star-edge-cloud/edge/utils/interfaces"
)

// IDriverTask -
type IDriverTask interface {
	uinterface.ITask
	// SetDriver -
	SetDriver(IDriver, IDriverHandler)
}
