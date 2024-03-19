package service

import (
	"Social_Media_Project_BE/features/user"
	"Social_Media_Project_BE/helper"
	"Social_Media_Project_BE/middlewares"
	"errors"
	"log"
	"strconv"

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
	helper.Recast(&update_data, &validate)
	err := s.v.Struct(&validate)
	if err != nil {
		log.Println("error validate", err.Error())
		return err
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
	helper.Recast(&login_data, &validate)
	err := s.v.Struct(&validate)
	if err != nil {
		log.Println("error validate", err.Error())
		return "", err
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
	var change_password = false
	if update_data.Password != "" {
		var validate user.UpdatePassword
		helper.Recast(&update_data, &validate)
		err := s.v.Struct(&validate)
		if err != nil {
			log.Println("error validate", err.Error())
			return err
		}

		change_password = true
	} else {
		var validate user.Update
		helper.Recast(&update_data, &validate)
		err := s.v.Struct(&validate)
		if err != nil {
			log.Println("error validate", err.Error())
			return err
		}
	}

	// Convert id to uint
	id_int, err := strconv.ParseUint(decodeID, 10, 32)
	if err != nil {
		return errors.New(helper.ErrorUserInput)
	}
	update_data.ID = uint(id_int)

	// Hashing Password
	if change_password {
		newPassword, err := s.pm.HashPassword(update_data.Password)
		if err != nil {
			return errors.New(helper.ErrorGeneralServer)
		}
		update_data.Password = newPassword
	}

	// Update Data
	if err := s.model.Update(update_data, change_password); err != nil {
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
