package device

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"star-edge-cloud/edge/core/config"
	"star-edge-cloud/edge/core/share"
	"star-edge-cloud/edge/models"
	"star-edge-cloud/edge/utils/common"
	"strconv"
	"strings"

	// 导入驱动
	_ "github.com/mattn/go-sqlite3"
)

// DeviceManager - 设备管理，包括设备注册，注销,设备信息查询等
type DeviceManager struct {
	Conf *config.Config
}

// AddDevice - 注册一个设备
func (dm *DeviceManager) AddDevice(file multipart.File, dev *models.Device) error {
	// 创建设备运行文件
	os.Mkdir("./plugins/device/"+dev.ID, os.ModePerm)
	// path := "./plugins/device/" + id + "/" + head.Filename
	path := fmt.Sprintf("./plugins/device/%[1]s/%[2]s", dev.ID, dev.FileName)
	fw, _ := os.Create(path)
	io.Copy(fw, file)
	fw.Close()

	// json配置
	jsonpath := fmt.Sprintf("./plugins/device/%[1]s/dev.json", dev.ID)
	content, _ := json.Marshal(dev)
	ioutil.WriteFile(jsonpath, content, 0644)
	// 安装daemon
	cmdpath := fmt.Sprintf("./plugins/device/%[1]s/", dev.ID)
	common.ExecDeamonCommand(share.WorkingDir, cmdpath, dev.FileName, "install")
	// str := common.ExecDeamonCommand(cmdpath, dev.FileName, "install")
	status := dm.GetStatus(dev)
	if status == -1 {
		// return fmt.Errorf("服务没有被安装:%[1]s,原因：%[2]s", dev.ID, "请使用root账号安装。")
	}

	// 信息入库
	sqlStmt1 := `
	INSERT INTO device (id, name, file_name,describe,registry_time,type,other,protocol,conf,status,liseners,command_server_address,log_base_url)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`
	if _, err := share.SqliteDB.Exec(sqlStmt1,
		dev.ID,
		dev.Name,
		dev.FileName,
		dev.Describe,
		dev.RegistryTime,
		dev.Type, dev.Other,
		dev.Protocol,
		dev.Conf,
		dev.Status,
		dev.Listeners,
		dev.ServerAddress,
		dev.LogBaseURL); err != nil {
		log.Printf("%q: %s\n", err, sqlStmt1)
		return err
	}
	return nil
}

// RemoveDevice - 删除一个设备
func (dm *DeviceManager) RemoveDevice(deviceID string) error {
	dev, err := dm.GetDevice(deviceID)
	if dev == nil || err != nil {
		return fmt.Errorf("未获取到设备：%s", deviceID)
	}

	// remove service
	cmdpath := fmt.Sprintf(`./plugins/device/%s/`, deviceID)
	str := common.ExecDeamonCommand(share.WorkingDir, cmdpath, dev.FileName, "remove")
	// log.Println(cmdpath + dev.FileName)
	status := dm.GetStatus(dev)
	if status == -1 {
		return fmt.Errorf("服务没有被删除:%[1]s,状态：%[2]s,原因：%[3]s", dev.ID, strconv.Itoa(status), str)
	}
	// remove directory
	if err := os.RemoveAll(fmt.Sprintf(`./plugins/device/%s`, deviceID)); err != nil {
		return err
	}

	// remove db row
	sqlStmt1 := `
	DELETE FROM device WHERE id=?;
	`
	if _, err = share.SqliteDB.Exec(sqlStmt1, deviceID); err != nil {
		return err
	}
	return nil
}

// UpdateDevice -
func (dm *DeviceManager) UpdateDevice(device *models.Device) error {
	sqlStmt1 := `
	UPDATE device
	SET name = ?, file_name = ?,describe=?,registry_time=?,type=?,other=?,protocol=?,conf=?,status=?,liseners=?, command_server_address=?, log_base_url=?
	WHERE id=?;
	`
	if _, err := share.SqliteDB.Exec(sqlStmt1,
		device.Name,
		device.FileName,
		device.Describe,
		device.RegistryTime,
		device.Type,
		device.Other,
		device.Protocol,
		device.Conf,
		device.Status,
		device.Listeners,
		device.LogBaseURL,
		device.ServerAddress,

		device.ID); err != nil {
		log.Printf("%q: %s\n", err, sqlStmt1)
	}
	return nil
}

// GetDevice - 删除一个设备
func (dm *DeviceManager) GetDevice(ID string) (*models.Device, error) {
	sqlStmt1 := `
	SELECT id, name, file_name,describe,registry_time,type,other,protocol,conf,status,liseners, command_server_address, log_base_url
	FROM device
	WHERE id=?
	`
	dev := &models.Device{}
	if rows, err := share.SqliteDB.Query(sqlStmt1, ID); err == nil {
		if rows.Next() {
			rows.Scan(&dev.ID,
				&dev.Name,
				&dev.FileName,
				&dev.Describe,
				&dev.RegistryTime,
				&dev.Type,
				&dev.Other,
				&dev.Protocol,
				&dev.Conf,
				&dev.Status,
				&dev.Listeners,
				&dev.ServerAddress,
				&dev.LogBaseURL)
		}
		rows.Close()
	} else {
		return nil, err
	}

	return dev, nil
}

// QueryAllDevice - 查询所有设备
func (dm *DeviceManager) QueryAllDevice() (devices []models.Device, err error) {
	sqlStmt1 := `
	SELECT *
	FROM device
	`
	var rows *sql.Rows
	// defer rows.Close()
	if rows, err = share.SqliteDB.Query(sqlStmt1); err == nil {
		for rows.Next() {
			dev := models.Device{}
			rows.Scan(&dev.ID,
				&dev.Name,
				&dev.FileName,
				&dev.Describe,
				&dev.RegistryTime,
				&dev.Type,
				&dev.Other,
				&dev.Protocol,
				&dev.Conf,
				&dev.Status,
				&dev.Listeners,
				&dev.ServerAddress,
				&dev.LogBaseURL)
			devices = append(devices, dev)
		}
		rows.Close()
	}
	return devices, err
}

// GetStatus -
func (dm *DeviceManager) GetStatus(dev *models.Device) int {
	path := fmt.Sprintf(`./plugins/device/%s`, dev.ID)
	result := common.ExecCheckStatus(share.WorkingDir, path, dev.FileName, "status")

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
func (dm *DeviceManager) Run(id string) error {
	dev, _ := dm.GetDevice(id)

	path := fmt.Sprintf(`./plugins/device/%s`, dev.ID)
	str := common.ExecDeamonCommand(share.WorkingDir, path, dev.FileName, "start")
	status := dm.GetStatus(dev)
	if status != 2 {
		return fmt.Errorf("服务没有被运行:%[1]s,原因：%[2]s", dev.ID, str)
	}
	dev.Status = 2
	dm.UpdateDevice(dev)

	return nil
}

// Stop -
func (dm *DeviceManager) Stop(id string) error {
	dev, _ := dm.GetDevice(id)
	path := fmt.Sprintf(`./plugins/device/%s`, dev.ID)
	str := common.ExecDeamonCommand(share.WorkingDir, path, dev.FileName, "stop")
	status := dm.GetStatus(dev)
	if status != 1 {
		return fmt.Errorf("服务没有被运行:%[1]s,原因：%[2]s", dev.ID, str)
	}

	return nil
}
