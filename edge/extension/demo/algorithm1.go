package demo

import (
	"errors"
	"star-edge-cloud/edge/utils/common"
)

// Algorithm1 -
type Algorithm1 struct{}

// Calculate -
func (it *Algorithm1) Calculate(v interface{}) ([]byte, error) {
	i := v.(int)
	if i%10 == 0 {
		return common.Int2Bytes(i), nil
	}

	return nil, errors.New("结果错误")
}
