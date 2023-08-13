package models

import (
	// "gorm.io/gorm"
	"time"
)

type Article struct {
	ID        uint
	Title     string
	Content   string
	AuthorID  uint
	Category  string
	Tags      string
	Comments  []Comment
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Comment struct {
	ID        uint
	Content   string
	AuthorID  uint
	ArticleID uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID        uint 	 `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	PasswordHash  string `gorm:"not null"`
	Articles  []Article // Relationship with articles struct
	Comments  []Comment // Relationship with comments struct
	CreatedAt time.Time
	UpdatedAt time.Time
}