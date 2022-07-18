package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var localDBDNS string = "127.0.0.1"
var localDBUsername string = "root"
var localDBPassword string = "lavi6310"

func ConnStageLocalDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/assistant", localDBUsername, localDBPassword, localDBDNS))
	return db, err
}