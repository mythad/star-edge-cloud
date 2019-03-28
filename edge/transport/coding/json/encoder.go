package json

import "encoding/json"

// JSONEncoder -
type JSONEncoder struct {
}

// Encode -
func (it *JSONEncoder) Encode(v interface{}) []byte {
	result, _ := json.Marshal(v)
	return result
}
