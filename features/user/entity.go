package user

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Controller interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type Service interface {
	Register(register_data User) error
	Login(login_data User) (string, error)
	Profile(token *jwt.Token) (User, error)
	Update(token *jwt.Token, update_data User) error
	Delete(token *jwt.Token) error
}

type Model interface {
	Register(register_data User) error
	Login(email string) (User, error)
	Profile(id string) (User, error)
	Update(data User, pass bool) error
	Delete(id string) error
}

type User struct {
	gorm.Model
	Name     string
	Hp       string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
}

type Login struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,alphanumunicode"`
}

type Register struct {
	Name     string `validate:"required,alpha,min=4"`
	Hp       string `validate:"required,number,min=11"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,alphanumunicode"`
}

type Update struct {
	Name  string `validate:"alpha,min=4"`
	Hp    string `validate:"number,min=11"`
	Email string `validate:"email"`
}

type UpdatePassword struct {
	Password string `validate:"alphanumunicode,min=8"`
}
