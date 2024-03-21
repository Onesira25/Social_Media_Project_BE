package data

import (
	"time"
)

type Comment struct {
	Id        uint `gorm:"primaryKey"`
	Username  string
	PostId    uint
	CreatedAt time.Time
	Comment   string
}
