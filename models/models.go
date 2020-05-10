package models

import (
	"awesomeProject/package/settings"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

var database *gorm.DB

func init()  {
	var (
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)

	db, err := settings.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Failed to get db config 'database': %v", err)
	}

	dbType = db.Key("TYPE").String()
	dbName = db.Key("NAME").String()
	user = db.Key("USER").String()
	password = db.Key("PASSWORD").String()
	host = db.Key("HOST").String()
	tablePrefix = db.Key("TABLE_PREFIX").String()

	database, err := gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", user, password, host, dbName))
	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	database.SingularTable(true)
	database.DB().SetMaxIdleConns(10)
	database.DB().SetMaxOpenConns(100)
}

func CloseDB()  {
	defer database.Close()
}

