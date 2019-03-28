package interfaces

// ITask -
type ITask interface {
	// LoadConfig -
	LoadConfig(v interface{})
	// LoadMetadata -
	LoadMetadata(v interface{})
	// Initialize -
	Initialize()
	// BeginCollect -
	Begin()
	// EndCollect -
	End()
}
