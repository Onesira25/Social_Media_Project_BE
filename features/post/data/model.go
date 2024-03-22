package data

import (
	"Social_Media_Project_BE/features/comment/data"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Username string `gorm:"type:varchar(15)"`
	Image    string
	Caption  string
	Comments []data.Comment `gorm:"foreignKey:PostId;references:ID"`
}
