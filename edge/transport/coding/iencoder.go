package coding

// IEncoder - 用户自定义协议解码器
type IEncoder interface {
	Encode(interface{}) []byte
}
