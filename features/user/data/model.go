package data

import "time"

type User struct {
	ID        uint      `json:"-" form:"-" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
	Fullname  string    `json:"fullname" form:"fullname"`
	Username  string    `json:"usename" form:"usename"`
	Handphone string    `json:"handphone" form:"handphone" gorm:"unique"`
	Email     string    `json:"email" form:"email" gorm:"unique"`
	Password  string    `json:"-" form:"-"`
}
