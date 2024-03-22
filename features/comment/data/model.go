package data

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Username  string `gorm:"type:varchar(15)"`
	PostId    uint
	CreatedAt time.Time
	Comment   string
}
