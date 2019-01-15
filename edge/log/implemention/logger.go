package implemention

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"star_cloud/edge/models"
	"time"

	// 导入驱动，golang要求必须注释一下
	_ "github.com/mattn/go-sqlite3"
)

// Logger -
type Logger struct {
	LogDBPath string
}

// Write - 写日志
func (helper *Logger) Write(info *models.LogInfo) {
	sqlStmt1 := `
	INSERT INTO log (id, message, time,level)
	VALUES (?, ?, ?, ?);
	`

	path := fmt.Sprintf("file:%s?cache=shared", helper.LogDBPath)
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	if _, err = db.Exec(sqlStmt1, info.ID, info.Message, info.Time, info.Level); err != nil {
		log.Printf("%q: %s\n", err, sqlStmt1)
	}
}

// QueryLevel - 查询日志
func (helper *Logger) QueryLevel(start time.Time, end time.Time, level int) (infos []models.LogInfo, err error) {
	sqlStmt1 := `
	SELECT id, message, time,level
	FROM log
	WHERE ? <= time and ? >= time and level = ?
	`

	path := fmt.Sprintf("file:%s?cache=shared", helper.LogDBPath)
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if rows, err := db.Query(sqlStmt1, start, end, level); err == nil {
		for rows.Next() {
			info := models.LogInfo{}
			rows.Scan(&info.ID, &info.Message, info.Time, info.Level)
			infos = append(infos, info)
		}
	}
	return infos, nil
}

// QueryMinLevel - 查询日志
func (helper *Logger) QueryMinLevel(start time.Time, end time.Time, level int) (infos []models.LogInfo, err error) {
	sqlStmt1 := `
	SELECT id, message, time,level
	FROM log
	WHERE ? <= time and ? >= time and level = ?
	`

	path := fmt.Sprintf("file:%s?cache=shared", helper.LogDBPath)
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if rows, err := db.Query(sqlStmt1, start, end, level); err == nil {
		for rows.Next() {
			info := models.LogInfo{}
			rows.Scan(&info.ID, &info.Message, info.Time, info.Level)
			infos = append(infos, info)
		}
	}
	return infos, nil
}

// QueryTop - 查询日志
func (helper *Logger) QueryTop(top int) (infos []models.LogInfo, err error) {
	sqlStmt1 := `
	select * from log order by time limit 0,?;
	`

	path := fmt.Sprintf("file:%s?cache=shared", helper.LogDBPath)
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if rows, err := db.Query(sqlStmt1, top); err == nil {
		for rows.Next() {
			info := models.LogInfo{}
			rows.Scan(&info.ID, &info.Message, info.Time, info.Level)
			infos = append(infos, info)
		}
	}
	return infos, nil
}

// Delete -
func (helper *Logger) Delete(id string) error {
	sqlStmt1 := `
	DELETE FROM log WHERE id=?;
	`

	path := fmt.Sprintf("file:%s?cache=shared", helper.LogDBPath)
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return err
	}
	defer db.Close()
	if _, err = db.Exec(sqlStmt1, id); err != nil {
		return err
	}
	return nil
}

// InitLogDB - 初始化日志数据库
func (helper *Logger) InitLogDB() {
	if _, err := os.Stat(helper.LogDBPath); err == nil {
		log.Println("日志文件：" + helper.LogDBPath)
		return
	}

	sqlStmt2 := `
	create table log (id string not null primary key, message text,time string, level integer);
	delete from log;
	`

	db1, err := sql.Open("sqlite3", helper.LogDBPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db1.Close()

	if _, err = db1.Exec(sqlStmt2); err != nil {
		log.Printf("%q: %s\n", err, sqlStmt2)
	}

}
