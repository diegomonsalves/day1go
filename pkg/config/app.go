package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //averiguar por qué el _ al inicio
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "diego:123qweasd/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
