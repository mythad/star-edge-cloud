package xml

import "encoding/xml"

// XMLDecoder -
type XMLDecoder struct{}

// Decode -
func (it *XMLDecoder) Decode(data []byte, result interface{}) {
	xml.Unmarshal(data, result)
}
