package interfaces

import "star-edge-cloud/edge/transport/server/interfaces"

// IDriverHandler -
type IDriverHandler interface {
	interfaces.IHandler
	OnError(error)
	OnInitialized()
	OnBegin()
	OnEnd()
}
