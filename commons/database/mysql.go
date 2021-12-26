package database

import (
	"fmt"

	"tirelease/commons/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Mysql handler infomation
type MysqlInfo struct {
	DB *gorm.DB
	// Anything else...
}

var DBConn = &MysqlInfo{}

func Connect() {
	// Params
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
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
