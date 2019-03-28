package xml

import "encoding/xml"

// XMLEncoder -
type XMLEncoder struct {
}

// Encode -
func (it *XMLEncoder) Encode(v interface{}) []byte {
	result, _ := xml.Marshal(v)
	return result
}
