package services

import (
	post "Social_Media_Project_BE/features/post"
	"Social_Media_Project_BE/helper"
	"Social_Media_Project_BE/middlewares"
	"context"
	"errors"
	"log"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type services struct {
	m post.PostModel
	v *validator.Validate
}

func NewTodoService(model post.PostModel) post.PostServices {
	return &services{
		m: model,
		v: validator.New(),
	}
}

func (s *services) Create(token *jwt.Token, image *multipart.FileHeader, caption string) error {
	decodeUsername := middlewares.DecodeTokenUsername(token)
	if decodeUsername == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return errors.New("data tidak valid")
	}

	src, err := image.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	cld, err := cloudinary.NewFromURL("cloudinary://641886961486146:HylqIrzq6ZTtaqzzbScWz5v9-aM@dajfvp3yw")
	if err != nil {
		log.Println("error connecting to cloudinary:", err.Error())
		return err
	}

	resp, err := cld.Upload.Upload(context.Background(), src, uploader.UploadParams{})
	if err != nil {
		log.Println("cloudinary upload error:", err.Error())
		return err
	}

	imageUrl := resp.SecureURL

	err = s.m.Create(decodeUsername, imageUrl, caption)
	if err != nil {
		return errors.New(helper.ServerGeneralError)
	}

	return nil
}

func (s *services) Edit(token *jwt.Token, postID string, image *multipart.FileHeader, caption string) error {
	decodeUsername := middlewares.DecodeTokenUsername(token)
	if decodeUsername == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return errors.New("data tidak valid")
	}

	src, err := image.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	cld, err := cloudinary.NewFromURL("cloudinary://641886961486146:HylqIrzq6ZTtaqzzbScWz5v9-aM@dajfvp3yw")
	if err != nil {
		log.Println("error connecting to cloudinary:", err.Error())
		return err
	}

	resp, err := cld.Upload.Upload(context.Background(), src, uploader.UploadParams{})
	if err != nil {
		log.Println("cloudinary upload error:", err.Error())
		return err
	}

	imageUrl := resp.SecureURL

	err = s.m.Edit(decodeUsername, postID, imageUrl, caption)
	if err != nil {
		return err
	}

	return nil
}

func (s *services) Posts(username string, limit string) ([]post.Post, error) {
	result, err := s.m.Posts(username, limit)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *services) PostById(postID string) (post.Post, error) {
	result, err := s.m.PostById(postID)
	if err != nil {
		return post.Post{}, err
	}

	return result, nil
}

func (s *services) Delete(token *jwt.Token, postID string) error {
	decodeUsername := middlewares.DecodeTokenUsername(token)
	if decodeUsername == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return errors.New("data tidak valid")
	}

	err := s.m.Delete(decodeUsername, postID)
	if err != nil {
		return err
	}

	return nil
}
