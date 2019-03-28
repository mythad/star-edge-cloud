package interfaces

// IAlgorithm --算法
type IAlgorithm interface {
	Calculate(v interface{}) ([]byte, error)
}
