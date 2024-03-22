package data

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID   uint
	Username string
	PostID   uint
	Comment  string
}
