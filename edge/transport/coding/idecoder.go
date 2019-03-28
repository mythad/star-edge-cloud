package coding

// IDecoder - 用户自定义协议编码器
type IDecoder interface {
	Decode([]byte, interface{})
}
