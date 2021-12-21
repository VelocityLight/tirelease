package database

import (
	"fmt"

	"database/sql"
	"tirelease/configs"

	_ "github.com/go-sql-driver/mysql"
)

// Connection returns infomation
type ConnectionInfo struct {
	DB *sql.DB
	// Anything else...
}

var DBConn = &ConnectionInfo{}

func Connect() {
	// Params
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		configs.Config.Mysql.UserName,
		configs.Config.Mysql.PassWord,
		configs.Config.Mysql.Host,
		configs.Config.Mysql.Port,
		configs.Config.Mysql.DataBase,
		configs.Config.Mysql.CharSet)

	// Connect
	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err.Error())
	}
	DBConn.DB = db

	// Close(Delayed)
	defer db.Close()
}
