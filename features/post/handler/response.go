package handler

import (
	"Social_Media_Project_BE/features/comment"
	"time"
)

type GetPostWithCommentsResponse struct {
	Id        uint              `json:"id"`
	Username  string            `json:"username"`
	CreatedAt time.Time         `json:"created_at"`
	Image     string            `json:"image"`
	Caption   string            `json:"caption"`
	Comments  []comment.Comment `json:"comments"`
}
