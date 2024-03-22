package posting

import (
	"Social_Media_Project_BE/features/comment"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type PostController interface {
	Create() echo.HandlerFunc
	Edit() echo.HandlerFunc
	Posts() echo.HandlerFunc
	PostById() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type PostModel interface {
	Create(username string, newPost Post) error
	Edit(username string, postID string, editPost Post) error
	Posts(username string, limit string) ([]Post, error)
	PostById(postID string) (Post, error)
	Delete(username string, postID string) error
}

type PostServices interface {
	Create(token *jwt.Token, newPost Post) error
	Edit(token *jwt.Token, postID string, EditPost Post) error
	Posts(username string, limit string) ([]Post, error)
	PostById(postID string) (Post, error)
	Delete(token *jwt.Token, postID string) error
}

type Post struct {
	Id        uint              `json:"id"`
	CreatedAt time.Time         `json:"created_at"`
	Username  string            `json:"username"`
	Image     string            `json:"image"`
	Caption   string            `json:"caption"`
	Comments  []comment.Comment `json:"comments"`
}

type CreatePost struct {
	Image   string
	Caption string
}

type EditPost struct {
	Image   string
	Caption string
}
