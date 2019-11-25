package model

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/restapi_belajar?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DB.SingularTable(true)
}
