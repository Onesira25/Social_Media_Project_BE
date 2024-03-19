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

func (s *service) Register(update_data user.User) error {
	// Check Validate
	var validate user.Register
	validate.Name = update_data.Name
	validate.Email = update_data.Email
	validate.Hp = update_data.Hp
	validate.Password = update_data.Password
	err := s.v.Struct(&validate)
	if err != nil {
		return errors.New(helper.ErrorInvalidValidate)
	}

	// Hashing Password
	newPassword, err := s.pm.HashPassword(update_data.Password)
	if err != nil {
		return errors.New(helper.ErrorGeneralServer)
	}
	update_data.Password = newPassword

	// Do Register
	err = s.model.Register(update_data)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			return errors.New(mysqlErr.Message)
		}
		return errors.New(helper.ErrorGeneralServer)
	}

	return nil
}

func (s *service) Login(login_data user.User) (string, error) {
	// Check Validate
	var validate user.Login
	validate.Email = login_data.Email
	validate.Password = login_data.Password
	err := s.v.Struct(&validate)
	if err != nil {
		return "", errors.New(helper.ErrorInvalidValidate)
	}

	// Do Login & Get Password
	dbData, err := s.model.Login(login_data.Email)
	if err != nil {
		return "", errors.New(helper.ErrorDatabaseNotFound)
	}

	// Compare Password
	if err := s.pm.ComparePassword(login_data.Password, dbData.Password); err != nil {
		return "", errors.New(helper.ErrorUserCredential)
	}

	// Create Token
	token, err := middlewares.GenerateJWT(strconv.Itoa(int(dbData.ID)))
	if err != nil {
		return "", errors.New(helper.ErrorGeneralServer)
	}

	// Finished
	return token, nil
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
	return result, nil
}

func (s *service) Update(token *jwt.Token, update_data user.User) error {
	// Get ID From Token
	decodeID := middlewares.DecodeToken(token)

	// Check Validate Password & Others
	var validate user.Update
	validate.Name = update_data.Name
	validate.Email = update_data.Email
	validate.Hp = update_data.Hp
	validate.Password = update_data.Password
	err := s.v.Struct(&validate)
	if err != nil {
		if strings.Contains(err.Error(), "Name") {
			update_data.Name = ""
		}
		if strings.Contains(err.Error(), "Email") {
			update_data.Email = ""
		}
		if strings.Contains(err.Error(), "Hp") {
			update_data.Hp = ""
		}
		if strings.Contains(err.Error(), "Password") {
			update_data.Password = ""
		}
		if strings.Count(err.Error(), "\n") == 3 {
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
