package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

// DB name:	 		golang
// DB username: 	root
// DB password: 	root22YCA:
// DB port: 		localhost:3306

var urlDSN = "root:root22YCA:@tcp(localhost:3306)/golang?parseTime=true"
var err error

//database migrate.
func DataMigration() {
	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("connection failed :(")
	}
	//Tables
	Database.AutoMigrate(&Employee{})
	Database.AutoMigrate(&Product{})
}
