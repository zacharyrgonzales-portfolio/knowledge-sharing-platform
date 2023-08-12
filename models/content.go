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
	Tags      []string
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
	ID        uint
	Username  string
	Email     string
	Password  string
	Articles  []Article
	Comments  []Comment
	CreatedAt time.Time
	UpdatedAt time.Time
}