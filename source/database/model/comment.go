package model

import "time"

type Comment struct {
	ID        uint      `gorm:"primaryKey"`
	Content   string
	UserID    uint
	PostID    uint
	CreatedAt time.Time
	UpdatedAt time.Time

	User User
	Post Post
}