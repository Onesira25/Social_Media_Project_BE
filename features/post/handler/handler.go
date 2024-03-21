package handler

import (
	post "Social_Media_Project_BE/features/post"
	"Social_Media_Project_BE/helper"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type controller struct {
	s post.PostServices
}

func NewHandler(service post.PostServices) post.PostController {
	return &controller{
		s: service,
	}
}

func (ct *controller) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input CreatePostRequest
		err := c.Bind(&input)
		if err != nil {
			log.Println("error bind data:", err.Error())
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.UserInputFormatError, nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		token, ok := c.Get("user").(*jwt.Token)
		defer func() {
			if err := recover(); err != nil {
				log.Println("error jwt creation:", err)

			}
		}()
		if !ok {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		var inputProcess post.Post
		inputProcess.Image = input.Image
		inputProcess.Caption = input.Caption

		err = ct.s.Create(token, inputProcess)
		if err != nil {
			log.Println("error insert db:", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
		}

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "berhasil menambahkan kegiatan", nil))
	}
}

func (ct *controller) Edit() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		defer func() {
			if err := recover(); err != nil {
				log.Println("error jwt creation:", err)

			}
		}()
		if !ok {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		var input EditPostRequest
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.UserInputFormatError, nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		var processInput post.Post
		processInput.Image = input.Image
		processInput.Caption = input.Caption

		postID := c.Param("postID")

		err = ct.s.Edit(token, postID, processInput)
		if err != nil {
			var code = http.StatusInternalServerError
			if strings.Contains(err.Error(), "validation") || strings.Contains(err.Error(), "cek kembali") {
				code = http.StatusBadRequest
			}
			return c.JSON(code,
				helper.ResponseFormat(code, err.Error(), nil))
		}
		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "success update task", nil))
	}
}

func (ct *controller) Posts() echo.HandlerFunc {
	return func(c echo.Context) error {
		var inputUserId = c.QueryParam("username")
		var inputLimit = c.QueryParam("limit")

		result, err := ct.s.Posts(inputUserId, inputLimit)
		if err != nil {
			var code = http.StatusInternalServerError
			if strings.Contains(err.Error(), "validation") || strings.Contains(err.Error(), "cek kembali") {
				code = http.StatusBadRequest
			}
			return c.JSON(code,
				helper.ResponseFormat(code, err.Error(), nil))
		}
		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "successfully get posts", result))
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
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		postID := c.Param("postID")

		err := ct.s.Delete(token, postID)
		if err != nil {
			var code = http.StatusInternalServerError
			if strings.Contains(err.Error(), "validation") || strings.Contains(err.Error(), "cek kembali") {
				code = http.StatusBadRequest
			}
			return c.JSON(code,
				helper.ResponseFormat(code, err.Error(), nil))
		}
		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "success update task", nil))
	}
}

func (ct *controller) PostById() echo.HandlerFunc {
	return func(c echo.Context) error {
		postID := c.Param("postID")

		result, err := ct.s.PostById(postID)
		if err != nil {
			var code = http.StatusInternalServerError
			if strings.Contains(err.Error(), "validation") || strings.Contains(err.Error(), "cek kembali") {
				code = http.StatusBadRequest
			}
			return c.JSON(code,
				helper.ResponseFormat(code, err.Error(), nil))
		}

		var withComment GetPostWithCommentsResponse
		withComment.Id = result.Id
		withComment.CreatedAt = result.CreatedAt
		withComment.Username = result.Username
		withComment.Image = result.Image
		withComment.Caption = result.Caption
		withComment.Comments = result.Comments

		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "success update task", result))
	}
}
