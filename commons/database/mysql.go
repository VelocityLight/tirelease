package database

import (
	"fmt"

	"tirelease/commons/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connection returns infomation
type ConnectionInfo struct {
	DB *gorm.DB
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
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	DBConn.DB = db

	// Close(Delayed)
	// defer db.Close()
}
