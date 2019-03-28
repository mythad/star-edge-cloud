package interfaces

// IRulesEngine -
type IRulesEngine interface {
	LoadRules([]byte)
	Handle(interface{})
	// GetRule(string) models.Rule
	// Match(*models.Rule, map[string]interface{}) (bool, error)
}
