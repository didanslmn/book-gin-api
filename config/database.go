package config

import (
	"fmt"
	"log"

	"github.com/didanslmn/gin-restful-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=godgame357 dbname=book_pustaka port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("db connection error", err)
	}
	err = db.AutoMigrate(&model.Book{})
	if err != nil {
		log.Fatal("failed to migrate database", err)
	}
	fmt.Println("database conecteed and migrate succes")
	DB = db
}
