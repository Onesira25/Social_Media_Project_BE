package services

import (
	post "Social_Media_Project_BE/features/post"
	"Social_Media_Project_BE/helper"
	"Social_Media_Project_BE/middlewares"
	"errors"
	"log"

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

func (s *services) Create(token *jwt.Token, newPost post.Post) error {
	decodeUsername := middlewares.DecodeToken(token)
	if decodeUsername == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return errors.New("data tidak valid")
	}

	err := s.m.Create(decodeUsername, newPost)
	if err != nil {
		return errors.New(helper.ServerGeneralError)
	}

	return nil
}

func (s *services) Edit(token *jwt.Token, postID string, editPost post.Post) error {
	decodeUsername := middlewares.DecodeToken(token)
	if decodeUsername == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return errors.New("data tidak valid")
	}

	err := s.m.Edit(decodeUsername, postID, editPost)
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
	decodeUsername := middlewares.DecodeToken(token)
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
