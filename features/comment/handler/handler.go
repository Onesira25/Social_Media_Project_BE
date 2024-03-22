package handler

import (
	"Social_Media_Project_BE/features/comment"
	"Social_Media_Project_BE/helper"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type controller struct {
	s comment.CommentServices
}

func NewHandler(service comment.CommentServices) comment.CommentController {
	return &controller{
		s: service,
	}
}

func (ct *controller) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input CreateCommentRequest
		err := c.Bind(&input)
		if err != nil {
			log.Println("error bind data:", err.Error())
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.UserInputFormatError, nil))
			}
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		token, ok := c.Get("user").(*jwt.Token)
		defer func() {
			if err := recover(); err != nil {
				log.Println("error jwt creation:", err)

			}
		}()
		if !ok {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		err = ct.s.Create(token, input.PostId, input.Comment)
		if err != nil {
			log.Println("error insert db:", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
		}

		return c.JSON(helper.ResponseFormat(http.StatusCreated, "success create comment", nil))
	}
}

func (ct *controller) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		defer func() {
			if err := recover(); err != nil {
				log.Println("error jwt creation:", err)

			}
		}()
		if !ok {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		commentID := c.Param("commentID")

		err := ct.s.Delete(token, commentID)
		if err != nil {
			var code = http.StatusInternalServerError
			if strings.Contains(err.Error(), "validation") || strings.Contains(err.Error(), "cek kembali") {
				code = http.StatusBadRequest
			}
			return c.JSON(helper.ResponseFormat(code, err.Error(), nil))
		}
		return c.JSON(helper.ResponseFormat(http.StatusOK, "success delete comment", nil))
	}
}
