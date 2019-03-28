package rtmp

import (
	zrtmp "github.com/zhangpeihao/gortmp"
)

func test() {
	// To connect FMS server
	obConn, _ := zrtmp.Dial("url", &outboundConnHandler{}, 100)

	// To connect
	obConn.Connect()

	// // When new stream created, handler event OnStreamCreated() would been called
	// func (handler *TestOutboundConnHandler) OnStreamCreated(stream rtmp.OutboundStream) {
	// 	// To play
	// 	err = stream.Play(*streamName, nil, nil, nil)
	// 	// Or publish
	// 	err = stream.Publish(*streamName, "live")
	// }

	// // To publish data
	// stream.PublishAudioData(data, deltaTimestamp)
	// // or
	// stream.PublishVideoData(data, deltaTimestamp)
	// // or
	// stream.PublishData(tagHeader.TagType, data, deltaTimestamp)

	// // You can close stream by
	// stream.Close()

	// // You can close connection by
	// obConn.Close()
}

type outboundConnHandler struct{}

// OnStatus -When connection status changed
func (it *outboundConnHandler) OnStatus(obConn zrtmp.OutboundConn) {

}

// OnStreamCreated -
func (it *outboundConnHandler) OnStreamCreated(obConn zrtmp.OutboundConn, stream zrtmp.OutboundStream) {
	// To play
	stream.Play("", nil, nil, nil)
	// Or publish
	// stream.Publish("", "live")
}

// Received message
func (it *outboundConnHandler) OnReceived(conn zrtmp.Conn, message *zrtmp.Message) {

}

// Received command
func (it *outboundConnHandler) OnReceivedRtmpCommand(conn zrtmp.Conn, command *zrtmp.Command) {

}

// Connection closed
func (it *outboundConnHandler) OnClosed(conn zrtmp.Conn) {

}
