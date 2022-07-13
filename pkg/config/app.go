package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	A := "akhil:Axlesharma@12@/simplerest?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", A)
	if err != nil {
		log.Fatal(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
