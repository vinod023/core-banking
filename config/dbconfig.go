package config

import (
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func ConfigDB() *gorm.DB {

	// custom GORM
	DB, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/corebanking?parseTime=true")
	if err != nil {
		panic(err)
	}

	return DB
}
