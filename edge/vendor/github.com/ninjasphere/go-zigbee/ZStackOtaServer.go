package zigbee

import (
	"github.com/gogo/protobuf/proto"
	"github.com/ninjasphere/go-zigbee/otasrvr"
)

type ZStackOta struct {
	*ZStackServer
}

type zStackOtaCommand interface {
	proto.Message
	GetCmdId() otasrvr.OtaMgrCmdIdT
}

// SendCommand sends a protobuf Message to the Z-Stack OTA server, and waits for the response
func (s *ZStackOta) SendCommand(request zStackOtaCommand, response zStackOtaCommand) error {

	return s.sendCommand(
		&zStackCommand{
			message:   request,
			commandID: uint8(request.GetCmdId()),
		},
		&zStackCommand{
			message:   response,
			commandID: uint8(response.GetCmdId()),
		},
	)

}

func ConnectToOtaServer(hostname string, port int) (*ZStackOta, error) {
	server, err := connectToServer("OTA", uint8(otasrvr.ZStackOTASysIDs_RPC_SYS_PB_OTA_MGR), hostname, port)
	if err != nil {
		return nil, err
	}

	return &ZStackOta{server}, nil
}
