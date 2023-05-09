package database

import (
	"fmt"
	"log"
	"os"

	"github.com/PeemXD/expenses-gin/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() {
	fmt.Println("connected! from fmt")
	var err error
	// Db, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	Db, err = gorm.Open(postgres.Open(os.Getenv("POSTGRESQL_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		log.Println("connected!")
	}
	log.Println("connected! from log")
	fmt.Println("connected! from fmt")
	Db.AutoMigrate(&model.Expenses{})

}
