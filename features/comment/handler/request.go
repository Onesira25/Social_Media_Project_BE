package handler

type CreateCommentRequest struct {
	PostId  uint   `json:"post_id"`
	Comment string `json:"comment"`
}
