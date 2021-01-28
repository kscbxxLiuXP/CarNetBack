package service

import (
	"CarNetBack/core/database"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func SetUp(){
	db = database.ConnetDB()
}