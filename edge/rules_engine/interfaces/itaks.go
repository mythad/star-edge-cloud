package interfaces

import "star-edge-cloud/edge/utils/interfaces"

// IRulesEngineTask -
type IRulesEngineTask interface {
	interfaces.ITask
	SetRuleEngine(IRulesEngine)
}
