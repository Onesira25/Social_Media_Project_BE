package comment

import "time"

type Comment struct {
	Id        uint
	UserId    uint
	PostId    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Comment   string
}
