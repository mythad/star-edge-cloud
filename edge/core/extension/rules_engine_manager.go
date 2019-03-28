package extension

import (
	"fmt"
	"star-edge-cloud/edge/core/share"
	"star-edge-cloud/edge/utils/common"
	"strings"
)

// RulesEngineManager - 存储服务管理
type RulesEngineManager struct {
}

// GetStatus -
func (dm *RulesEngineManager) GetStatus() int {
	path := `./`
	result := common.ExecCheckStatus(share.WorkingDir, path, "rules_engine", "status")

	// Service (pid  ******) is running...
	if strings.Contains(result, "running") {
		return 2
	}

	// Service is stopped
	if strings.Contains(result, "stopped") {
		return 1
	}

	// Service is not installed
	if strings.Contains(result, "installed") {
		return 0
	}

	return -1
}

// Run -
func (dm *RulesEngineManager) Run() error {
	path := `./`
	if dm.GetStatus() <= 0 {
		dm.Install()
	}
	str := common.ExecDeamonCommand(share.WorkingDir, path, "rules_engine", "start")
	status := dm.GetStatus()
	if status != 2 {
		return fmt.Errorf("服务没有被运行:rules_engine,原因：%[1]s", str)
	}

	return nil
}

// Stop -
func (dm *RulesEngineManager) Stop() error {
	path := `./`
	str := common.ExecDeamonCommand(share.WorkingDir, path, "rules_engine", "stop")
	status := dm.GetStatus()
	if status != 1 {
		return fmt.Errorf("服务没有被运行:rules_engine,原因：%[1]s", str)
	}

	return nil
}

// Install -
func (dm *RulesEngineManager) Install() error {
	path := `./`
	str := common.ExecDeamonCommand(share.WorkingDir, path, "rules_engine", "install")
	status := dm.GetStatus()
	if status != 2 {
		return fmt.Errorf("服务没有被运行:rules_engine,原因：%[1]s", str)
	}

	return nil
}
