package zigbee

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gogo/protobuf/proto"
	"github.com/ninjasphere/go-zigbee/gateway"
)

type ZStackGateway struct {
	*ZStackServer
	pendingResponses         map[uint32]*pendingGatewayResponse
	zoneStateListeners       []zoneStateListener
	attributeReportListeners []attributeReportListener
	boundClustersListeners   []boundClusterListener
}

type zStackGatewayCommand interface {
	proto.Message
	GetCmdId() gateway.GwCmdIdT
}

type pendingGatewayResponse struct {
	response zStackGatewayCommand
	finished chan error
}

type attributeReportListener struct {
	address  uint64
	endpoint uint32
	channel  chan *gateway.GwAttributeReportingInd
}

type zoneStateListener struct {
	address  uint64
	endpoint uint32
	channel  chan *gateway.DevZoneStatusChangeInd
}

func (s *ZStackGateway) OnZoneState(addr uint64, endpoint uint32) chan *gateway.DevZoneStatusChangeInd {
	listener := zoneStateListener{addr, endpoint, make(chan *gateway.DevZoneStatusChangeInd)}

	s.zoneStateListeners = append(s.zoneStateListeners, listener)

	return listener.channel
}

type boundClusterListener struct {
	address  uint64
	endpoint uint32
	cluster  uint32
	channel  chan *gateway.GwZclFrameReceiveInd
}

func (s *ZStackGateway) OnBoundCluster(addr uint64, endpoint uint32, cluster uint32) chan *gateway.GwZclFrameReceiveInd {
	listener := boundClusterListener{addr, endpoint, cluster, make(chan *gateway.GwZclFrameReceiveInd)}

	s.boundClustersListeners = append(s.boundClustersListeners, listener)

	return listener.channel
}

func (s *ZStackGateway) waitForSequenceResponse(sequenceNumber uint32, response zStackGatewayCommand, timeoutDuration time.Duration) error {
	// We accept uint32 as thats what comes back from protobuf
	log.Debugf("Waiting for sequence %d", sequenceNumber)
	_, exists := s.pendingResponses[sequenceNumber]
	if exists {
		s.pendingResponses[sequenceNumber].finished <- fmt.Errorf("Another command with the same sequence id (%d) has been sent.", sequenceNumber)
	}

	pending := &pendingGatewayResponse{
		response: response,
		finished: make(chan error),
	}
	s.pendingResponses[sequenceNumber] = pending

	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(timeoutDuration)
		timeout <- true
	}()

	var err error

	select {
	case error := <-pending.finished:
		err = error
	case <-timeout:
		err = fmt.Errorf("The request timed out after %s", timeoutDuration)
	}

	s.pendingResponses[sequenceNumber] = nil

	return err
}

// SendAsyncCommand sends a command that requires an async response from the device, using ZCL SequenceNumber
func (s *ZStackGateway) SendAsyncCommand(request zStackGatewayCommand, response zStackGatewayCommand, timeout time.Duration) error {
	confirmation := &gateway.GwZigbeeGenericCnf{}

	//	spew.Dump("sending", request)

	err := s.SendCommand(request, confirmation)

	if err != nil {
		return err
	}

	//spew.Dump(confirmation)

	if confirmation.Status.String() != "STATUS_SUCCESS" {
		return fmt.Errorf("Invalid confirmation status: %s", confirmation.Status.String())
	}

	return s.waitForSequenceResponse(*confirmation.SequenceNumber, response, timeout)
}

// SendCommand sends a protobuf Message to the Z-Stack server, and waits for the response
func (s *ZStackGateway) SendCommand(request zStackGatewayCommand, response zStackGatewayCommand) error {

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

func (s *ZStackGateway) onIncomingCommand(commandID uint8, bytes *[]byte) {

	//bytes := <-s.Incoming

	log.Debugf("gateway: Got gateway message %s", gateway.GwCmdIdT_name[int32(commandID)])

	//commandID := uint8((*bytes)[1])

	if commandID == uint8(gateway.GwCmdIdT_GW_ATTRIBUTE_REPORTING_IND) {
		log.Debugf("gateway: Parsing as GwAttributeReportingInd")
		report := &gateway.GwAttributeReportingInd{}
		err := proto.Unmarshal(*bytes, report)
		if err != nil {
			log.Errorf("gateway: Could not read attribute report: %s, %v", err, *bytes)
			return
		}

		if log.IsDebugEnabled() {
			spew.Dump("Got attribute report", report)
		}

		if len(s.attributeReportListeners) > 0 {
			for _, listener := range s.attributeReportListeners {
				if listener.address == *report.SrcAddress.IeeeAddr && listener.endpoint == *report.SrcAddress.EndpointId {

					go func(l chan *gateway.GwAttributeReportingInd) {
						l <- report
					}(listener.channel)

				}

			}
		} else {
			log.Debugf("gateway: Received an unhandled attribute report from % X : %v", *report.SrcAddress.IeeeAddr, report)
		}

		return
	}

	if commandID == uint8(gateway.GwCmdIdT_GW_ZCL_FRAME_RECEIVE_IND) {

		log.Debugf("gateway: Parsing as GwZclFrameReceiveInd")

		frame := &gateway.GwZclFrameReceiveInd{}
		err := proto.Unmarshal(*bytes, frame)
		if err != nil {
			log.Errorf("gateway: Could not read GwZclFrameReceiveInd: %s, %v", err, *bytes)
			return
		}

		if log.IsDebugEnabled() {
			spew.Dump("Got zcl frame (bound cluster?)", frame)
		}

		handled := false

		for _, listener := range s.boundClustersListeners {
			if listener.address == *frame.SrcAddress.IeeeAddr && listener.endpoint == *frame.SrcAddress.EndpointId && listener.cluster == *frame.ClusterId {

				go func(l chan *gateway.GwZclFrameReceiveInd) {
					handled = true
					l <- frame
				}(listener.channel)

			}

		}

		if !handled {
			log.Debugf("gateway: Received an unhandled zcl frame from %+v : %+v", *frame.SrcAddress, frame)
		}

		return
	}

	if commandID == uint8(gateway.GwCmdIdT_DEV_ZONE_STATUS_CHANGE_IND) {

		log.Debugf("gateway: Parsing as GwDevZoneStatusChangeInd")

		zoneStatus := &gateway.DevZoneStatusChangeInd{}
		err := proto.Unmarshal(*bytes, zoneStatus)
		if err != nil {
			log.Errorf("gateway: Could not read zone status change: %s, %v", err, *bytes)
			return
		}

		if log.IsDebugEnabled() {
			spew.Dump("Got zone status change", zoneStatus)
		}

		if len(s.zoneStateListeners) > 0 {
			for _, listener := range s.zoneStateListeners {
				if listener.address == *zoneStatus.SrcAddress.IeeeAddr && listener.endpoint == *zoneStatus.SrcAddress.EndpointId {

					go func(l chan *gateway.DevZoneStatusChangeInd) {
						l <- zoneStatus
					}(listener.channel)
				}
			}
		} else {
			log.Debugf("gateway: Received an unhandled zone status change from % X : %v", *zoneStatus.SrcAddress.IeeeAddr, zoneStatus)
		}

		return
	}

	var sequenceNumber uint32

	if commandID == uint8(gateway.GwCmdIdT_ZIGBEE_GENERIC_CNF) {
		log.Debugf("gateway: Parsing as GwZigbeeGenericCnf")
		message := &gateway.GwZigbeeGenericCnf{}
		err := proto.Unmarshal(*bytes, message)
		if err != nil {
			log.Errorf("gateway: Could not read generic confirmation: %s, %v", err, *bytes)
			return
		}

		sequenceNumber = *message.SequenceNumber

	} else {
		log.Debugf("gateway: Parsing as GwZigbeeGenericRspInd")
		message := &gateway.GwZigbeeGenericRspInd{} // Not always this, but it will always give us the sequence number?
		err := proto.Unmarshal(*bytes, message)
		if err != nil {
			log.Errorf("gateway: Could not get sequence number from incoming gateway message : %s, %v", err, *bytes)
			return
		}

		sequenceNumber = *message.SequenceNumber
	}

	log.Debugf("gateway: Got an incoming gateway message, sequence:%d", sequenceNumber)

	if sequenceNumber == 0 {
		log.Debugf("gateway: Failed to get a sequence number from an incoming gateway message: %x", bytes)
	}

	pending := s.pendingResponses[sequenceNumber]

	if pending == nil {
		log.Infof("gateway: Received response to sequence number %d but we aren't listening for it", sequenceNumber)
	} else {

		if uint8(pending.response.GetCmdId()) != commandID {
			pending.finished <- fmt.Errorf("Wrong ZCL response type. Wanted: 0x%X Received: 0x%X", uint8(pending.response.GetCmdId()), commandID)
		} else {
			pending.finished <- proto.Unmarshal(*bytes, pending.response)
		}
	}

}

func ConnectToGatewayServer(hostname string, port int) (*ZStackGateway, error) {
	server, err := connectToServer("Gateway", uint8(gateway.ZStackGwSysIdT_RPC_SYS_PB_GW), hostname, port)
	if err != nil {
		return nil, err
	}

	gateway := &ZStackGateway{
		ZStackServer:             server,
		pendingResponses:         map[uint32]*pendingGatewayResponse{},
		zoneStateListeners:       []zoneStateListener{},
		attributeReportListeners: []attributeReportListener{},
		boundClustersListeners:   []boundClusterListener{},
	}

	server.onIncoming = func(commandID uint8, bytes *[]byte) {
		gateway.onIncomingCommand(commandID, bytes)
	}

	return gateway, nil
}
