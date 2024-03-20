package user

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Controller interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type Service interface {
	Register(register_data Register) error
	Login(login_data User) (LoginResponse, error)
	Profile(token *jwt.Token) (User, error)
	Update(token *jwt.Token, update_data User) error
	Delete(token *jwt.Token) error
}

type Model interface {
	Register(register_data User) error
	Login(input string) (User, error)
	Profile(id string) (User, error)
	Update(data User) error
	Delete(id string) error
}

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Fullname  string
	Username  string
	Handphone string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Biodata   string
	Password  string
}

type Login struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type LoginResponse struct {
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
	Fullname  string    `json:"fullname" form:"fullname"`
	Username  string    `json:"username" form:"username"`
	Handphone string    `json:"handphone" form:"handphone"`
	Email     string    `json:"email" form:"email"`
	Biodata   string    `json:"biodata" form:"biodata"`
	Token     string    `json:"token" form:"token"`
}

type Register struct {
	Fullname  string `validate:"required,min=5"`
	Username  string `validate:"required,min=5"`
	Handphone string `validate:"required,number,min=11,max=14"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required,min=8"`
}

type Update struct {
	Fullname  string `validate:"required,min=5"`
	Username  string `validate:"required,min=5"`
	Handphone string `validate:"required,number,min=11,max=14"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required,min=8"`
}
