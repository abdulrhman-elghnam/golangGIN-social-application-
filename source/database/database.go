package database

import (
	"golang/source/database/model"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connection() {
	
	db, err := gorm.Open(sqlite.Open(os.Getenv("DATABASE_URI")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.Comment{})
}
