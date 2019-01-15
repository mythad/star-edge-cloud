package extension

import (
	"fmt"
	"star_cloud/edge/utils/common"
	"strings"
)

// LogManager - 存储服务管理
type LogManager struct {
}

// GetStatus -
func (dm *LogManager) GetStatus() int {
	path := `./`
	result := common.ExecCheckStatus(path, "log", "status")

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
func (dm *LogManager) Run() error {
	path := `./`
	if dm.GetStatus() <= 0 {
		dm.Install()
	}
	str := common.ExecDeamonCommand(path, "log", "start")
	status := dm.GetStatus()
	if status != 2 {
		return fmt.Errorf("服务没有被运行:log,原因：%[1]s", str)
	}

	return nil
}

// Stop -
func (dm *LogManager) Stop() error {
	path := `./`
	str := common.ExecDeamonCommand(path, "log", "stop")
	status := dm.GetStatus()
	if status != 1 {
		return fmt.Errorf("服务没有被运行:log,原因：%[1]s", str)
	}

	return nil
}

// Install -
func (dm *LogManager) Install() error {
	path := `./`
	str := common.ExecDeamonCommand(path, "log", "install")
	status := dm.GetStatus()
	if status != 2 {
		return fmt.Errorf("服务没有被运行:log,原因：%[1]s", str)
	}

	return nil
}
