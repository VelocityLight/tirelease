// Tool Url: https://github.com/go-gorm/gorm
// Tool Guide: https://gorm.io/docs/

package database

import (
	"fmt"
	"time"

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

func Connect(config *configs.ConfigYaml) {
	// Params
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		config.Mysql.UserName,
		config.Mysql.PassWord,
		config.Mysql.Host,
		config.Mysql.Port,
		config.Mysql.DataBase,
		config.Mysql.CharSet,
	)

	// Connect
	conn, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	sqlDB, err := conn.DB()
	if err != nil {
		panic(err.Error())
	}
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 600)

	DBConn.DB = conn

	// Close(Delayed)
	// defer db.Close()
}
