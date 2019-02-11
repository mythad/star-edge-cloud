package amqp

import "testing"

// RestAgreementImpl - rest的功能实现
type RestAgreementImpl struct {
}

// Push - 推送数据
func (r *RestAgreementImpl) Push([]byte) (string, error) {
	return "", nil
}

// Query - 查询数据
func (r *RestAgreementImpl) Query(string) ([]byte, error) {
	return nil, nil
}

// Config - 配置
func (r *RestAgreementImpl) Config(string) error {
	return nil
}

func TestAll(t *testing.T) {

}
