package data

import (
	comment "Social_Media_Project_BE/features/comment/data"
	post "Social_Media_Project_BE/features/post/data"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname  string
	Username  string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Handphone string `gorm:"unique"`
	Password  string
	Biodata   string
	Posts     []post.Post       `gorm:"foreignKey:Username;references:Username"`
	Comments  []comment.Comment `gorm:"foreignKey:Username;references:Username"`
}
