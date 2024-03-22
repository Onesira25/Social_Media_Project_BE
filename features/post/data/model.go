package data

import (
	"Social_Media_Project_BE/features/comment/data"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID   uint
	Username string
	Image    string
	Caption  string
	Comments []data.Comment
}
