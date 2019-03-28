package interfaces

import (
	uinterface "star-edge-cloud/edge/utils/interfaces"
)

// IDriverTask -
type IDriverTask interface {
	uinterface.ITask
	// SetAlgorithm -
	SetAlgorithm(IAlgorithm)
}
