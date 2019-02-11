package extension

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
	"star-edge-cloud/edge/models"
	"star-edge-cloud/edge/utils/common"
	"strconv"
	"strings"

	// 导入驱动
	_ "github.com/mattn/go-sqlite3"
)

// ExtentionManager - 设备管理，包括设备注册，注销,设备信息查询等
type ExtentionManager struct {
	Conf *config.Config
}

// AddExtension - 注册一个设备
func (dm *ExtentionManager) AddExtension(file multipart.File, ext *models.Extension) error {
	// 创建设备运行文件
	os.Mkdir("./plugins/extension/"+ext.ID, os.ModePerm)
	path := fmt.Sprintf("./plugins/extension/%[1]s/%[2]s", ext.ID, ext.FileName)
	fw, _ := os.Create(path)
	io.Copy(fw, file)
	fw.Close()

	// json配置
	jsonpath := fmt.Sprintf("./plugins/extension/%[1]s/ext.json", ext.ID)
	content, _ := json.Marshal(ext)
	ioutil.WriteFile(jsonpath, content, 0644)
	// 安装daemon
	cmdpath := fmt.Sprintf("./plugins/extension/%[1]s/", ext.ID)
	str := common.ExecDeamonCommand(cmdpath, ext.FileName, "install")
	status := dm.GetStatus(ext)
	if status == -1 {
		return fmt.Errorf("服务没有被安装:%[1]s,原因：%[2]s", ext.ID, str)
	}

	sqlStmt1 := `
	INSERT INTO extension (id, name, file_name,describe,registry_time,type,other,protocol,conf,status,liseners,command_server_address,log_base_url)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`
	db, err := sql.Open("sqlite3", dm.Conf.MetadataDBPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if _, err = db.Exec(sqlStmt1,
		ext.ID,
		ext.Name,
		ext.FileName,
		ext.Describe,
		ext.RegistryTime,
		ext.Type, ext.Other,
		ext.Protocol,
		ext.Conf,
		ext.Status,
		ext.Listeners,
		ext.ServerAddress,
		ext.LogBaseURL); err != nil {
		log.Printf("%q: %s\n", err, sqlStmt1)
	}
	return nil
}

// RemoveExtension - 删除一个算法服务
func (dm *ExtentionManager) RemoveExtension(ID string) error {
	ext, err := dm.GetExtension(ID)
	if ext == nil || err != nil {
		return fmt.Errorf("未获取到设备：%s", ID)
	}

	// remove service
	cmdpath := fmt.Sprintf(`./plugins/extension/%s/`, ID)
	str := common.ExecDeamonCommand(cmdpath, ext.FileName, "remove")
	status := dm.GetStatus(ext)
	if status == -1 {
		return fmt.Errorf("服务没有被删除:%[1]s,状态：%[2]s,原因：%[3]s", ext.ID, strconv.Itoa(status), str)
	}
	// remove directory
	if err := os.RemoveAll(fmt.Sprintf(`./plugins/extension/%s`, ID)); err != nil {
		return err
	}

	sqlStmt1 := `
	DELETE FROM extension WHERE id=?;
	`
	db, err := sql.Open("sqlite3", dm.Conf.MetadataDBPath)
	if err != nil {
		return err
	}
	defer db.Close()
	if _, err := db.Exec(sqlStmt1, ID); err != nil {
		return err
	}
	return nil
}

// UpdateExtension -
func (dm *ExtentionManager) UpdateExtension(ext *models.Extension) error {
	sqlStmt1 := `
	UPDATE extension
	SET name = ?, file_name = ?,describe=?,registry_time=?,type=?,other=?,protocol=?,conf=?,status=?,liseners=?,command_server_address=?,log_base_url=?
	WHERE id=?;
	`
	db, err := sql.Open("sqlite3", dm.Conf.MetadataDBPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if _, err = db.Exec(sqlStmt1,
		ext.Name,
		ext.FileName,
		ext.Describe,
		ext.RegistryTime,
		ext.Type,
		ext.Other,
		ext.Protocol,
		ext.Conf,
		ext.Status,
		ext.Listeners,
		ext.ServerAddress,
		ext.LogBaseURL,
		ext.ID); err != nil {
		log.Printf("%q: %s\n", err, sqlStmt1)
	}
	return nil
}

// GetExtension - 删除一个算法服务
func (dm *ExtentionManager) GetExtension(ID string) (*models.Extension, error) {
	sqlStmt1 := `
	SELECT id, name, file_name,describe,registry_time,type,other,protocol,conf,status,liseners,command_server_address,log_base_url
	FROM extension
	WHERE id=?
	`
	db, err := sql.Open("sqlite3", dm.Conf.MetadataDBPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	ext := &models.Extension{}
	if rows, err := db.Query(sqlStmt1, ID); err == nil {
		if rows.Next() {
			rows.Scan(&ext.ID,
				&ext.Name,
				&ext.FileName,
				&ext.Describe,
				&ext.RegistryTime,
				&ext.Type,
				&ext.Other,
				&ext.Protocol,
				&ext.Conf,
				&ext.Status,
				&ext.Listeners,
				&ext.ServerAddress,
				&ext.LogBaseURL)
		}
		rows.Close()
	} else {
		return nil, err
	}

	return ext, nil
}

// QueryAllExtension - 查询所有算法服务
func (dm *ExtentionManager) QueryAllExtension(etype string) (exts []models.Extension, err error) {
	sqlStmt1 := `
	SELECT id, name, file_name,describe,registry_time,type,other,protocol,conf,status,liseners,command_server_address,log_base_url
	FROM extension where type=?
	`
	db, err := sql.Open("sqlite3", dm.Conf.MetadataDBPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var rows *sql.Rows
	// defer rows.Close()
	if rows, err = db.Query(sqlStmt1, etype); err == nil {
		for rows.Next() {
			ext := models.Extension{}
			rows.Scan(&ext.ID,
				&ext.Name,
				&ext.FileName,
				&ext.Describe,
				&ext.RegistryTime,
				&ext.Type,
				&ext.Other,
				&ext.Protocol,
				&ext.Conf,
				&ext.Status,
				&ext.Listeners,
				&ext.ServerAddress,
				&ext.LogBaseURL)
			exts = append(exts, ext)
		}
		rows.Close()
	}

	return exts, err
}

// GetStatus -
func (dm *ExtentionManager) GetStatus(ext *models.Extension) int {
	path := fmt.Sprintf(`./plugins/extension/%s`, ext.ID)
	result := common.ExecCheckStatus(path, ext.FileName, "status")
	// log.Println(path)
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
func (dm *ExtentionManager) Run(id string) error {
	ext, _ := dm.GetExtension(id)

	path := fmt.Sprintf(`./plugins/extension/%s`, ext.ID)
	str := common.ExecDeamonCommand(path, ext.FileName, "start")
	status := dm.GetStatus(ext)
	if status != 2 {
		return fmt.Errorf("服务没有被运行:%[1]s,原因：%[2]s", ext.ID, str)
	}

	return nil
}

// Stop -
func (dm *ExtentionManager) Stop(id string) error {
	ext, _ := dm.GetExtension(id)
	path := fmt.Sprintf(`./plugins/extension/%s`, ext.ID)
	str := common.ExecDeamonCommand(path, ext.FileName, "stop")
	status := dm.GetStatus(ext)
	if status != 1 {
		return fmt.Errorf("服务没有被运行:%[1]s,原因：%[2]s", ext.ID, str)
	}
	ext.Status = 2
	dm.UpdateExtension(ext)

	return nil
}
