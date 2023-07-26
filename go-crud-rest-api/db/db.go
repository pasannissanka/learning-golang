package db

import (
	todoentity "github.com/pasannissanka/learning-golang/go-crud-rest-api/modules/todo/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&todoentity.Todo{})
	if err != nil {
		return
	}

	DB = db
}
