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
		caption := c.FormValue("caption")
		file, err := c.FormFile("image")
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

		err = ct.s.Create(token, file, caption)
		if err != nil {
			log.Println("error upload image:", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
		}

		return c.JSON(helper.ResponseFormat(http.StatusCreated, "success create post", nil))
	}
}

func (ct *controller) Edit() echo.HandlerFunc {
	return func(c echo.Context) error {
		caption := c.FormValue("caption")
		file, err := c.FormFile("image")
		if err != nil {
			log.Print("error:", err.Error())
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, "image not found", nil))
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

		postID := c.Param("postID")

		err = ct.s.Edit(token, postID, file, caption)
		if err != nil {
			var code = http.StatusInternalServerError
			if strings.Contains(err.Error(), "validation") || strings.Contains(err.Error(), "cek kembali") {
				code = http.StatusBadRequest
			}
			return c.JSON(helper.ResponseFormat(code, err.Error(), nil))
		}
		return c.JSON(helper.ResponseFormat(http.StatusOK, "success edit post", nil))
	}
}

func (ct *controller) Posts() echo.HandlerFunc {
	return func(c echo.Context) error {
		var inputUsername = c.QueryParam("username")
		var inputPage = c.QueryParam("page")

		result, err := ct.s.Posts(inputUsername, inputPage)
		if err != nil {
			var code = http.StatusInternalServerError
			if strings.Contains(err.Error(), "validation") || strings.Contains(err.Error(), "cek kembali") {
				code = http.StatusBadRequest
			}
			return c.JSON(helper.ResponseFormat(code, err.Error(), nil))
		}
		return c.JSON(helper.ResponseFormat(http.StatusOK, "successfully get posts", result))
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

		postID := c.Param("postID")

		err := ct.s.Delete(token, postID)
		if err != nil {
			var code = http.StatusInternalServerError
			if strings.Contains(err.Error(), "validation") || strings.Contains(err.Error(), "cek kembali") {
				code = http.StatusBadRequest
			}
			return c.JSON(helper.ResponseFormat(code, err.Error(), nil))
		}
		return c.JSON(helper.ResponseFormat(http.StatusOK, "success delete post", nil))
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
			return c.JSON(helper.ResponseFormat(code, err.Error(), nil))
		}

		var withComment GetPostWithCommentsResponse
		withComment.Id = result.ID
		withComment.CreatedAt = result.CreatedAt
		withComment.Username = result.Username
		withComment.Image = result.Image
		withComment.Caption = result.Caption
		withComment.Comments = result.Comments

		return c.JSON(helper.ResponseFormat(http.StatusOK, "success get post", result))
	}
}
