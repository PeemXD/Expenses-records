package database

import (
	"os"

	"github.com/PeemXD/expenses-gin/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() {

	var err error
	Db, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	Db.AutoMigrate(&model.Expenses{})

}
