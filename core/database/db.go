package database

import (
	"CarNetBack/config"
	"CarNetBack/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var address model.Address

//连接数据库
func ConnetDB() *gorm.DB {
	var db *gorm.DB
	DATABASE_URL := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.APP.DB_Username,
		config.APP.DB_Password,
		config.APP.DB_IP,
		config.APP.DB_Port,
		config.APP.DB_Name)
	db, err := gorm.Open("mysql", DATABASE_URL)
	if err != nil {
		panic(err.Error())
	}

	db.SingularTable(true)
	return db
}
