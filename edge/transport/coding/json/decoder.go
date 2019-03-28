package json

import "encoding/json"

// JSONDecoder -
type JSONDecoder struct{}

// Decode -
func (it *JSONDecoder) Decode(data []byte, result interface{}) {
	json.Unmarshal(data, result)
}
