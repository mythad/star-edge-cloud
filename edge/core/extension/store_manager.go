package extension

import (
	"fmt"
	"star-edge-cloud/edge/core/share"
	"star-edge-cloud/edge/utils/common"
	"strings"
)

// StoreManager - 存储服务管理
type StoreManager struct {
}

// GetStatus -
func (dm *StoreManager) GetStatus() int {
	path := `./`
	result := common.ExecCheckStatus(share.WorkingDir, path, "store", "status")

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
func (dm *StoreManager) Run() error {
	path := `./`
	if dm.GetStatus() <= 0 {
		dm.Install()
	}
	str := common.ExecDeamonCommand(share.WorkingDir, path, "store", "start")
	status := dm.GetStatus()
	if status != 2 {
		return fmt.Errorf("服务没有被运行:store,原因：%[1]s", str)
	}

	return nil
}

// Stop -
func (dm *StoreManager) Stop() error {
	path := `./`
	str := common.ExecDeamonCommand(share.WorkingDir, path, "store", "stop")
	status := dm.GetStatus()
	if status != 1 {
		return fmt.Errorf("服务没有被运行:store,原因：%[1]s", str)
	}

	return nil
}

// Install -
func (dm *StoreManager) Install() error {
	path := `./`
	str := common.ExecDeamonCommand(share.WorkingDir, path, "store", "install")
	status := dm.GetStatus()
	if status != 2 {
		return fmt.Errorf("服务没有被运行:store,原因：%[1]s", str)
	}

	return nil
}
