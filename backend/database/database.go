package database

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(mysql.Open(os.Getenv("ConnectionString")), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to the database: ", err)
		return
	}
}
