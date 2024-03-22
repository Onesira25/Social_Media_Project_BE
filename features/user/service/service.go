package service

import (
	"Social_Media_Project_BE/features/user"
	"Social_Media_Project_BE/helper"
	"Social_Media_Project_BE/middlewares"
	"errors"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
)

type service struct {
	model user.Model
	pm    helper.PasswordManager
	v     *validator.Validate
}

func NewService(m user.Model) user.Service {
	return &service{
		model: m,
		pm:    helper.NewPasswordManager(),
		v:     validator.New(),
	}
}

func (s *service) Register(register_data user.Register) error {
	// Check Validate
	var validate user.Register
	validate.Fullname = register_data.Fullname
	validate.Username = register_data.Username
	validate.Email = register_data.Email
	validate.Handphone = register_data.Handphone
	validate.Password = register_data.Password
	err := s.v.Struct(&validate)
	if err != nil {
		return errors.New(helper.ErrorInvalidValidate)
	}

	// Hashing Password
	newPassword, err := s.pm.HashPassword(register_data.Password)
	if err != nil {
		return errors.New(helper.ErrorGeneralServer)
	}

	user_data := user.User{
		Fullname:  register_data.Fullname,
		Username:  register_data.Username,
		Handphone: register_data.Handphone,
		Email:     register_data.Email,
		Password:  newPassword,
	}

	// Do Register
	err = s.model.Register(user_data)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			return errors.New(mysqlErr.Message)
		}
		return errors.New(helper.ErrorGeneralServer)
	}

	return nil
}

func (s *service) Login(login_data user.User) (user.LoginResponse, error) {
	// Do Login & Get Password
	dbData, err := s.model.Login(login_data.Email)
	if err != nil {
		return user.LoginResponse{}, errors.New(helper.ErrorDatabaseNotFound)
	}

	// Compare Password
	if err := s.pm.ComparePassword(login_data.Password, dbData.Password); err != nil {
		return user.LoginResponse{}, errors.New(helper.ErrorUserCredential)
	}

	// Create Token
	token, err := middlewares.GenerateJWT(strconv.Itoa(int(dbData.ID)), dbData.Username)
	if err != nil {
		return user.LoginResponse{}, errors.New(helper.ErrorGeneralServer)
	}

	// Finished
	var result user.LoginResponse
	result.CreatedAt = dbData.CreatedAt.UTC()
	result.UpdatedAt = dbData.UpdatedAt.UTC()
	result.Email = dbData.Email
	result.Fullname = dbData.Fullname
	result.Username = dbData.Username
	result.Handphone = dbData.Handphone
	result.Biodata = dbData.Biodata
	result.Token = token
	return result, nil
}

func (s *service) Profile(token *jwt.Token) (user.User, error) {
	// Get ID From Token
	decodeID := middlewares.DecodeToken(token)

	// Get Profile
	result, err := s.model.Profile(decodeID)
	if err != nil {
		return user.User{}, err
	}

	// Finished
	result.CreatedAt = result.CreatedAt.UTC()
	result.UpdatedAt = result.UpdatedAt.UTC()
	return result, nil
}

func (s *service) Update(token *jwt.Token, update_data user.User) error {
	// Get ID From Token
	decodeID := middlewares.DecodeToken(token)

	// Check Validate Password & Others
	var validate user.Update
	validate.Fullname = update_data.Fullname
	validate.Username = update_data.Username
	validate.Email = update_data.Email
	validate.Handphone = update_data.Handphone
	validate.Password = update_data.Password
	err := s.v.Struct(&validate)
	if err != nil {
		if strings.Contains(err.Error(), "Fullname") {
			update_data.Fullname = ""
		}
		if strings.Contains(err.Error(), "Username") {
			update_data.Username = ""
		}
		if strings.Contains(err.Error(), "Email") {
			update_data.Email = ""
		}
		if strings.Contains(err.Error(), "Handphone") {
			update_data.Handphone = ""
		}
		if strings.Contains(err.Error(), "Password") {
			update_data.Password = ""
		}
		if update_data.Biodata == "" && strings.Count(err.Error(), "\n") >= 4 {
			return errors.New(helper.ErrorInvalidValidate)
		}
	}

	// Convert id to uint
	id_int, err := strconv.ParseUint(decodeID, 10, 32)
	if err != nil {
		return errors.New(helper.ErrorUserInput)
	}
	update_data.ID = uint(id_int)

	// Hashing Password
	if update_data.Password != "" {
		newPassword, err := s.pm.HashPassword(update_data.Password)
		if err != nil {
			return errors.New(helper.ErrorGeneralServer)
		}
		update_data.Password = newPassword
	}

	// Update Data
	if err := s.model.Update(update_data); err != nil {
		return err
	}

	// Finished
	return nil
}

func (s *service) Delete(token *jwt.Token) error {
	// Get ID From Token
	decodeID := middlewares.DecodeToken(token)

	// Delete Date
	if err := s.model.Delete(decodeID); err != nil {
		return err
	}

	// Finished
	return nil
}
