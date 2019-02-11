package zigbee

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"

	"github.com/gogo/protobuf/proto"
)

// ZStackServer holds the connection to one of the Z-Stack servers (nwkmgr, gateway and otasrvr)
type ZStackServer struct {
	name       string
	subsystem  uint8
	conn       net.Conn
	outgoing   chan *zStackPendingCommand
	received   chan *zStackReceivedPacket
	onIncoming func(uint8, *[]byte)
}

// ZStackPendingCommand is a thing
type zStackPendingCommand struct {
	request  *zStackCommand
	response *zStackCommand
	complete chan error
}

type zStackReceivedPacket struct {
	commandID uint8
	packet    []byte
}

// ZStackCommand contains a protobuf message and a command id
type zStackCommand struct {
	message   proto.Message
	commandID uint8
}

func (s *ZStackServer) sendCommand(request *zStackCommand, response *zStackCommand) error {

	if s == nil {
		log.Fatalf("receiver was nil!")
	}

	if response == nil {
		log.Fatalf("illegal argument: response was nil!")
	}

	pending := &zStackPendingCommand{
		request:  request,
		response: response,
		complete: make(chan error),
	}

	s.outgoing <- pending
	err := <-pending.complete

	return err
}

func (s *ZStackServer) transmitCommand(command *zStackCommand) error {

	proto.SetDefaults(command.message)

	packet, err := proto.Marshal(command.message)
	if err != nil {
		log.Fatalf("%s: Outgoing marshaling error: %s", s.name, err)
	}

	log.Debugf("Protobuf packet %x", packet)

	buffer := new(bytes.Buffer)

	// Add the Z-Stack 4-byte header
	err = binary.Write(buffer, binary.LittleEndian, uint16(len(packet))) // Packet length
	err = binary.Write(buffer, binary.LittleEndian, s.subsystem)         // Subsystem
	err = binary.Write(buffer, binary.LittleEndian, command.commandID)   // Command Id

	_, err = buffer.Write(packet)

	log.Debugf("%s: Sending packet: % X", s.name, buffer.Bytes())

	// Send it to the Z-Stack server
	_, err = s.conn.Write(buffer.Bytes())
	return err
}

func (s *ZStackServer) incomingLoop() {
	for {
		buf := make([]byte, 1024)
		n, err := s.conn.Read(buf)

		log.Debugf("Read %d from %s", n, s.name)
		if err != nil {
			log.Fatalf("%s: Error reading socket %s", s.name, err)
		}
		pos := 0

		for {
			var length uint16
			var incomingSubsystem uint8
			reader := bytes.NewReader(buf[pos:])
			err := binary.Read(reader, binary.LittleEndian, &length)
			if err != nil {
				log.Fatalf("%s: Failed to read packet length %s", s.name, err)
			}

			err = binary.Read(reader, binary.LittleEndian, &incomingSubsystem)
			if err != nil {
				log.Fatalf("%s: Failed to read packet subsystem %s", s.name, err)
			}

			log.Debugf("%s: Incoming subsystem %d (wanted: %d)", s.name, incomingSubsystem, s.subsystem)

			log.Debugf("%s: Found packet of size : %d", s.name, length)

			commandID := uint8(buf[pos+3])

			packet := buf[pos+4 : pos+4+int(length)]

			log.Debugf("%s: Command ID:0x%X Packet: % X", s.name, commandID, packet)

			s.received <- &zStackReceivedPacket{
				commandID: commandID,
				packet:    packet,
			}

			pos += int(length) + 4

			if pos >= n {
				break
			}
		}
	}
}

// This loop runs for ever. It processes one outgoing packet at a time (looking for matching responses) and any number
// of non-matching inbound packets.
func (s *ZStackServer) matchingLoop() {
NoOutgoing:
	for {
		var pending *zStackPendingCommand = nil

		select {
		case pending = <-s.outgoing:
			err := s.transmitCommand(pending.request)
			if err != nil {
				pending.complete <- err
				continue
			}
		case incoming := <-s.received:
			if s.onIncoming != nil {
				go s.onIncoming(incoming.commandID, &incoming.packet)
			} else {
				log.Errorf("%s: ERR: Unhandled incoming command, packet: %d, %X", s.name, incoming.commandID, incoming.packet)
			}
			continue
		}

		timeout := time.NewTimer(time.Duration(5) * time.Second)

		if pending == nil {
			log.Fatalf("logic error: pending == nil")
		} else if pending.response == nil {
			log.Fatalf("logic error: pending.response == nil")
		}

		for {
			select {
			case incoming := <-s.received:
				if incoming.commandID == pending.response.commandID {
					pending.complete <- proto.Unmarshal(incoming.packet, pending.response.message)
					continue NoOutgoing
				} else {
					if s.onIncoming != nil {
						go s.onIncoming(incoming.commandID, &incoming.packet)
					} else {
						log.Errorf("%s: ERR: Unhandled incoming command, packet: %d, %X", s.name, incoming.commandID, incoming.packet)
					}
					continue
				}
			case _ = <-timeout.C:
				pending.complete <- fmt.Errorf("timed out while waiting for response to commandID: %d", pending.response.commandID)
				continue NoOutgoing
			}
		}

	}

}

func connectToServer(name string, subsystem uint8, hostname string, port int) (*ZStackServer, error) {

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", hostname, port))

	if err != nil {
		return nil, err
	}

	server := &ZStackServer{
		name:      name,
		subsystem: subsystem,
		conn:      conn,
		outgoing:  make(chan *zStackPendingCommand),
		received:  make(chan *zStackReceivedPacket),
	}

	go server.incomingLoop()
	go server.matchingLoop()

	return server, nil
}

// proposed re-write
// one go routine extracts packets from the wire as fast as they come
// another go routine accepts outbound packets and matches inbound packets, times out waits for and matches inbound packets
