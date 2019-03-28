package interfaces

// IStore --协议
type IStore interface {
	// 保存数据
	Save(string, []byte) error
	// 查询数据
	Get(string) []byte
	// 查询数据
	Query(string) []byte
	// 删除数据
	Delete(string) int
	// 初始化
	Inialize()
	// 释放资源
	Dispose()
}
