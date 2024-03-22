package comment

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type CommentController interface {
	Create() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type CommentModel interface {
	Create(userID uint, username string, postID uint, newComment string) error
	Delete(username string, postID string) error
}

type CommentServices interface {
	Create(token *jwt.Token, postID uint, newComment string) error
	Delete(token *jwt.Token, commentID string) error
}

type Comment struct {
	CreatedAt time.Time
	Username  string
	Comment   string
}

type CreateComment struct {
	PostId  uint
	Comment string
}
