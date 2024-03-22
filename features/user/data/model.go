package data

import (
	comment "Social_Media_Project_BE/features/comment/data"
	post "Social_Media_Project_BE/features/post/data"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname  string
	Username  string `gorm:"type:varchar(15);unique"`
	Email     string
	Password  string
	Handphone string
	Biodata   string
	Posts     []post.Post
	Comments  []comment.Comment
}
