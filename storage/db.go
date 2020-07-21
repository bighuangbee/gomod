package storage

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"github.com/bighuangbee/gomod/config"
	"github.com/bighuangbee/gomod/loger"
)

var DB *gorm.DB

func init(){
	open()
}

func open() {
	var err error

	dbType 		:= config.ConfigData.DbType
	user 		:= config.ConfigData.DbUser
	password 	:= config.ConfigData.DbPassword
	//password 	:= "arMonitor2020."
	host 		:= config.ConfigData.DbHost
	dbName 		:= config.ConfigData.DbName


	if dbType == "mysql" {
		DB, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user,
			password,
			host,
			dbName))
	}else if dbType == "postgres"{

		DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s dbname=%s user=%s sslmode=disable password=%s",
			host,
			dbName,
			user,
			password))
	}


	if err != nil {
		panic("DataBase Connected Failed!" + err.Error())
	}

	DB.LogMode(config.ConfigData.Debug)
	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	DB.SetLogger(loger.Loger)

	loger.Info("MySql Client SetUp Success...", config.ConfigData.DbHost)
}

func CloseDB() {
	defer DB.Close()
}

