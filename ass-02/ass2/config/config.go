package config

import (
	"ass-02/ass2/structs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/ass_godb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(structs.Item{}, structs.Orders{})
	return db
}
