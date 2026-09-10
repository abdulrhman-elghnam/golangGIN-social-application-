package model

import "time"

type Post struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string
	Content   string
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time

	User     User
	Comments []Comment
}