package data

import (
	comment "Social_Media_Project_BE/features/comment/data"
	post "Social_Media_Project_BE/features/post/data"
	"time"
)

type User struct {
	Id        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Fullname  string
	Username  string `gorm:"unique"`
	Email     string
	Password  string
	Handphone string
	Biodata   string
	Posts     []post.Post       `gorm:"foreignKey:Username;references:Username"`
	Comments  []comment.Comment `gorm:"foreignKey:Username;references:Username"`
}
