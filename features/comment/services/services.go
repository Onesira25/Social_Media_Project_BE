package services

import (
	"Social_Media_Project_BE/features/comment"
	"Social_Media_Project_BE/helper"
	"Social_Media_Project_BE/middlewares"
	"errors"
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type services struct {
	m comment.CommentModel
	v *validator.Validate
}

func CommentService(model comment.CommentModel) comment.CommentServices {
	return &services{
		m: model,
		v: validator.New(),
	}
}

func (s *services) Create(token *jwt.Token, postID uint, comment string) error {
	decodeUsername := middlewares.DecodeTokenUsername(token)
	decodeUserID := middlewares.DecodeToken(token)
	if decodeUsername == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return errors.New("data tidak valid")
	}

	userID, _ := strconv.ParseUint(decodeUserID, 10, 32)

	err := s.m.Create(uint(userID), decodeUsername, postID, comment)
	if err != nil {
		return errors.New(helper.ServerGeneralError)
	}

	return nil
}

func (s *services) Delete(token *jwt.Token, commentID string) error {
	decodeUsername := middlewares.DecodeTokenUsername(token)
	if decodeUsername == "" {
		log.Println("error decode token:", "token tidak ditemukan")
		return errors.New("data tidak valid")
	}

	err := s.m.Delete(decodeUsername, commentID)
	if err != nil {
		return err
	}

	return nil
}
