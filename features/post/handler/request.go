package handler

type CreatePostRequest struct {
	Image   string `json:"image" form:"image"`
	Caption string `json:"caption" form:"caption"`
}

type EditPostRequest struct {
	Image   string `json:"image" form:"image"`
	Caption string `json:"caption" form:"caption"`
}
