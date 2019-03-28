package share

import (
	"database/sql"
	"fmt"
	"log"
	// 驱动
	_ "github.com/mattn/go-sqlite3"
)

// SqliteDB -
var SqliteDB *sql.DB
var connectionString string

// Sqlite3Repostiory -
// var Sqlite3Repostiory Repostiory

// // Repostiory -
// type Repostiory struct {
// 	connectionString string
// 	db               *sql.DB
// }

// Initialize -
func InitializeDB(filename string) {
	connectionString = fmt.Sprintf("file:%s?cache=shared&mode=rwc&_locking=NORMAL", filename)
	db, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	SqliteDB = db
}

// CreateTable -
func CreateTable(sql string) {
	if _, err := SqliteDB.Exec(sql); err != nil {
		log.Printf("%q: %s\n", err, sql)
	}
}

// // Get -
// func (it *Repostiory) Get(sql string, id string, v ...interface{}) error {
// 	if row := it.db.QueryRow(sql, id); row == nil {
// 		row.Scan(v)
// 	}

// 	return nil
// }

// // Add -
// func (it *Repostiory) Add(sql string, v ...interface{}) error {
// 	if _, err := it.db.Exec(sql, v); err != nil {
// 		log.Println(err)
// 		return err
// 	}

// 	return nil
// }

// // Update -
// func (it *Repostiory) Update(sql string, v ...interface{}) error {
// 	if _, err := it.db.Exec(sql, v); err != nil {
// 		log.Printf("%q: %s\n", err, sql)
// 		return err
// 	}
// 	return nil
// }

// // Remove -
// func (it *Repostiory) Remove(sql string, id string) error {
// 	if _, err := it.db.Exec(sql, id); err != nil {
// 		return err
// 	}
// 	return nil
// }

// // Query -
// func (it *Repostiory) Query(sql string, args ...interface{}) (rows *sql.Rows, err error) {
// 	return it.db.Query(sql, args)
// }

// // Dispose -
// func (it *Repostiory) Dispose() {
// 	if it.db != nil {
// 		it.db.Close()
// 	}
// }
