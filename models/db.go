package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = Connection()

	db.AutoMigrate(
		new(User),
		new(Drug),
		new(Transaction),
	)
}

// Connection
// check connection for database server
func Connection() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_apotek?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database server.")
	return db
}
