package database

import (
	"golang/source/database/model"
	"golang/source/database/repository"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var UserRepository *repository.BaseRepository[model.User]
var PostRepository *repository.BaseRepository[model.Post]
var CommentRepository *repository.BaseRepository[model.Comment]

func Connection() (*gorm.DB, error) {
	db, err := gorm.Open(
		sqlite.Open(os.Getenv("DATABASE_URI")),
		&gorm.Config{},
	)

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Comment{},
	)

	if err != nil {
		return nil, err
	}

	UserRepository = repository.NewBaseRepository[model.User](db)
	PostRepository = repository.NewBaseRepository[model.Post](db)
	CommentRepository = repository.NewBaseRepository[model.Comment](db)

	return db, nil
}