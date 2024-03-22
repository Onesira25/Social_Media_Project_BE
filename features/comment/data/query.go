package data

import (
	"Social_Media_Project_BE/features/comment"
	"errors"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func New(db *gorm.DB) comment.CommentModel {
	return &model{
		connection: db,
	}
}

func (cm *model) Create(userID uint, username string, postID uint, comment string) error {
	var inputProcess = Comment{
		Username: username,
		UserID:   userID,
		PostID:   postID,
		Comment:  comment,
	}

	qry := cm.connection.Create(&inputProcess)
	if err := qry.Error; err != nil {
		return err
	}

	if qry.RowsAffected < 1 {
		return errors.New("no data affected")
	}
	return nil
}

func (cm *model) Delete(username string, commentID string) error {
	qry := cm.connection.Where("id = ? AND username = ?", commentID, username).Delete(&Comment{})

	if err := qry.Error; err != nil {
		return err
	}

	if qry.RowsAffected < 1 {
		return errors.New("no data affected")
	}

	return nil
}
