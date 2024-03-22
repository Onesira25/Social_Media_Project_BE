package post

import (
	"Social_Media_Project_BE/features/comment"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PostController interface {
	Create() echo.HandlerFunc
	Edit() echo.HandlerFunc
	Posts() echo.HandlerFunc
	PostById() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type PostModel interface {
	Create(newPost Post) error
	Edit(userID string, postID string, editPost Post) error
	Posts(username string, limit string) ([]Post, error)
	PostById(postID string) (Post, error)
	Delete(username string, postID string) error
}

type PostServices interface {
	Create(token *jwt.Token, newPost Post) error
	Edit(token *jwt.Token, postID string, EditPost Post) error
	Posts(username string, page string) ([]Post, error)
	PostById(postID string) (Post, error)
	Delete(token *jwt.Token, postID string) error
}

type Post struct {
	gorm.Model
	Username string
	Image    string
	Caption  string
	Comments []comment.Comment
}

type CreatePost struct {
	Image   string
	Caption string
}

type EditPost struct {
	Image   string
	Caption string
}
