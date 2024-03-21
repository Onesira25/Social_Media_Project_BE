package data

import (
	"Social_Media_Project_BE/features/comment"
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

func (pm *model) Create(username string, image string, caption string) error {
	var inputProcess = Post{
		Username: username,
		Image:    image,
		Caption:  caption,
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

func (pm *model) Edit(username string, postID string, image string, caption string) error {
	qry := pm.connection.Model(&Post{}).Where("username = ? AND id = ?", username, postID).Update("image", image).Update("caption", caption)
	if err := qry.Error; err != nil {
		return err
	}

	if qry.RowsAffected < 1 {
		return errors.New("no data affected")
	}

	return nil
}

func (pm *model) Posts(username string, page string) ([]post.Post, error) {
	var posts []post.Post

	reqPage, _ := strconv.Atoi(page)

	if reqPage < 1 {
		reqPage = 1
	}

	if username == "" {
		err := pm.connection.Limit(10).Offset(reqPage*10 - 10).Find(&posts).Error
		if err != nil {
			return nil, err
		}
	} else {
		if err := pm.connection.Where("username = ?", username).Limit(10).Offset(reqPage*10 - 10).Find(&posts).Error; err != nil {
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

	if err := pm.connection.Where("post_id = ? AND username = ?", postID, username).Delete(&comment.Comment{}).Error; err != nil {
		return err
	}

	if err := pm.connection.Where("id = ? AND username = ?", postID, username).Delete(&Post{}).Error; err != nil {
		return err
	}
	return nil
}
