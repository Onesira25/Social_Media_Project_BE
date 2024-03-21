package data

import (
	post "Social_Media_Project_BE/features/post"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func New(db *gorm.DB) post.PostModel {
	return &model{
		connection: db,
	}
}

func (pm *model) Create(username string, post post.Post) error {
	var inputProcess = Post{
		Username: username,
		Image:    post.Image,
		Caption:  post.Caption,
	}

	qry := pm.connection.Create(&inputProcess)
	if err := qry.Error; err != nil {
		return err
	}

	if qry.RowsAffected < 1 {
		return errors.New("no data affected")
	}
	return nil
}

func (pm *model) Edit(username string, postID string, editPost post.Post) error {
	qry := pm.connection.Where("username = ? AND id = ?", username, postID).Updates(&editPost)
	if err := qry.Error; err != nil {
		return err
	}

	if qry.RowsAffected < 1 {
		return errors.New("no data affected")
	}

	return nil
}

func (pm *model) Posts(username string, limit string) ([]post.Post, error) {
	var posts []post.Post

	if username == "" || limit == "" {
		err := pm.connection.Find(&posts).Error
		if err != nil {
			return nil, err
		}
	} else {
		reqLimit, _ := strconv.Atoi(limit)
		if err := pm.connection.Where("username = ?", username).Find(&posts).Limit(reqLimit).Error; err != nil {
			return nil, err
		}
	}

	return posts, nil
}

func (pm *model) PostById(postID string) (post.Post, error) {
	var result post.Post
	if err := pm.connection.Model(&Post{}).Where("id = ?", postID).Preload("Comments").First(&result).Error; err != nil {
		return post.Post{}, err
	}

	return result, nil
}

func (pm *model) Delete(username string, postID string) error {
	if err := pm.connection.Where("id = ?", postID).Delete(&username).Error; err != nil {
		return err
	}
	return nil
}
