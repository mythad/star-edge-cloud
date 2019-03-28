package cache

// IPersistence -
type IPersistence interface {
	Push(interface{})
	Pop() (interface{}, bool)
	Clear()
}
