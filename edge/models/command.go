package models

// Command -
type Command struct {
	ID   string
	Type string
	// unix time
	Time int64
	Data []byte
}
