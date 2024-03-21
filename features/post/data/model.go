package data

import (
	"Social_Media_Project_BE/features/comment/data"
	"time"
)

type Post struct {
	Id        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Username  string
	Image     string
	Caption   string
	Comments  []data.Comment `gorm:"foreignKey:PostId;references:Id"`
}
