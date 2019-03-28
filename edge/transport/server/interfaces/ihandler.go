package interfaces

import "star-edge-cloud/edge/transport/coding"

// IHandler -
type IHandler interface {
	SetDecoder(coding.IDecoder)
	SetEncoder(coding.IEncoder)
	OnReceive(interface{}) ([]byte, error)
}
