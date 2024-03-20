package handler

import (
	"Social_Media_Project_BE/features/user"
	"Social_Media_Project_BE/helper"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type controller struct {
	service user.Service
}

func NewHandler(s user.Service) user.Controller {
	return &controller{
		service: s,
	}
}

func (ct *controller) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input user.User
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.ErrorUserInputFormat))
			}
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		err = ct.service.Register(input)
		if err != nil {
			return c.JSON(helper.ResponseFormat(helper.ErrorCode(err), err.Error()))
		}

		return c.JSON(helper.ResponseFormat(http.StatusCreated, "Congratulations, the data has been registered"))
	}
}

func (ct *controller) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginRequest
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.ErrorUserInputFormat))
			}
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		var processData user.User
		processData.Email = input.Email
		processData.Password = input.Password

		token, err := ct.service.Login(processData)
		if err != nil {
			return c.JSON(helper.ResponseFormat(helper.ErrorCode(err), err.Error()))
		}

		return c.JSON(helper.ResponseFormat(http.StatusOK, "login successful", map[string]any{"token": token}))
	}
}

func (ct *controller) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		profile, err := ct.service.Profile(token)
		if err != nil {
			return c.JSON(helper.ResponseFormat(http.StatusInternalServerError, helper.ErrorGeneralServer))
		}

		var profileResponse ProfileResponse
		profileResponse.ID = profile.ID
		profileResponse.CreatedAt = profile.CreatedAt
		profileResponse.UpdatedAt = profile.UpdatedAt
		profileResponse.Name = profile.Name
		profileResponse.Email = profile.Email
		profileResponse.Hp = profile.Hp

		return c.JSON(helper.ResponseFormat(http.StatusOK, "success get user profile", map[string]any{"user": profileResponse}))
	}
}

func (ct *controller) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input user.User
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.ErrorUserInputFormat))
			}
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		err = ct.service.Update(token, input)
		if err != nil {
			return c.JSON(helper.ResponseFormat(helper.ErrorCode(err), err.Error()))
		}

		return c.JSON(helper.ResponseFormat(http.StatusOK, "success update user"))
	}
}

func (ct *controller) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(helper.ResponseFormat(http.StatusBadRequest, helper.ErrorUserInput))
		}

		err := ct.service.Delete(token)
		if err != nil {
			return c.JSON(helper.ResponseFormat(helper.ErrorCode(err), err.Error()))
		}
		return c.JSON(helper.ResponseFormat(http.StatusOK, "success delete user"))
	}
}
