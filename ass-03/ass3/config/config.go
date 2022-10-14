package config

import (
	"ass-03/ass3/structs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/auto_reload_ass?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(structs.Reload{})
	return db
}
