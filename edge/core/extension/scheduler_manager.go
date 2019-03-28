package extension

import (
	"fmt"
	"star-edge-cloud/edge/core/share"
	"star-edge-cloud/edge/utils/common"
	"strings"
)

// SchedulerManager - 存储服务管理
type SchedulerManager struct {
}

// GetStatus -
func (dm *SchedulerManager) GetStatus() int {
	path := `./`
	result := common.ExecCheckStatus(share.WorkingDir, path, "scheduler", "status")

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
func (dm *SchedulerManager) Run() error {
	path := `./`
	if dm.GetStatus() <= 0 {
		dm.Install()
	}
	str := common.ExecDeamonCommand(share.WorkingDir, path, "scheduler", "start")
	status := dm.GetStatus()
	if status != 2 {
		return fmt.Errorf("服务没有被运行:log,原因：%[1]s", str)
	}

	return nil
}

// Stop -
func (dm *SchedulerManager) Stop() error {
	path := `./`
	str := common.ExecDeamonCommand(share.WorkingDir, path, "scheduler", "stop")
	status := dm.GetStatus()
	if status != 1 {
		return fmt.Errorf("服务没有被运行:log,原因：%[1]s", str)
	}

	return nil
}

// Install -
func (dm *SchedulerManager) Install() error {
	path := `./`
	str := common.ExecDeamonCommand(share.WorkingDir, path, "scheduler", "install")
	status := dm.GetStatus()
	if status != 2 {
		return fmt.Errorf("服务没有被运行:log,原因：%[1]s", str)
	}

	return nil
}
