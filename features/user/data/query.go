package data

import (
	"Social_Media_Project_BE/features/user"
	"Social_Media_Project_BE/helper"
	"errors"

	"gorm.io/gorm"
)

type model struct {
	connection *gorm.DB
}

func New(db *gorm.DB) user.Model {
	return &model{
		connection: db,
	}
}

func (m *model) Register(newData user.User) error {
	return m.connection.Create(&newData).Error
}

func (m *model) Login(email string) (user.User, error) {
	var result user.User
	err := m.connection.Where("email = ?", email).First(&result).Error
	return result, err
}

func (m *model) Profile(id string) (user.User, error) {
	var result user.User
	err := m.connection.Where("id = ?", id).Find(&result).Error
	return result, err
}

func (m *model) Update(data user.User, pass bool) error {
	var query *gorm.DB
	if pass {
		query = m.connection.Model(&data).Select("password").Updates(&data)
	} else {
		query = m.connection.Model(&data).Omit("password").Updates(&data)
	}

	if query.Error != nil {
		return errors.New(helper.ErrorGeneralDatabase)
	} else if query.RowsAffected == 0 {
		return errors.New(helper.ErrorNoRowsAffected)
	}
	return nil
}

func (m *model) Delete(id string) error {
	if query := m.connection.Where("id = ?", id).Delete(&user.User{}); query.Error != nil {
		return errors.New(helper.ErrorGeneralDatabase)
	} else if query.RowsAffected == 0 {
		return errors.New(helper.ErrorDatabaseNotFound)
	}
	return nil
}
