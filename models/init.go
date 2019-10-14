package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	dbuser := "root"
	dbpassword := "root"
	dbhost := "127.0.0.1:3306"
	dbname := "topstory"
	charset := "utf8mb4"
	parseTime := "True"
	loc := "UTC"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=%s",
		dbuser,
		dbpassword,
		dbhost,
		dbname,
		charset,
		parseTime,
		loc)
	var err error
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true) // 禁用表名复数
}
