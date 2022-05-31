package services

import (
	"github.com/arnab333/golang-app-invite-service/models"
)

func (conn *dbConnection) InsertUser(user *models.User) (int64, error) {
	result := conn.DB.Create(user)

	if result.Error != nil {
		return 0, result.Error
	}

	return int64(user.ID), nil
}

func (conn *dbConnection) FindUser(filter *models.User) (user *models.User, err error) {
	result := conn.DB.Where(&filter).First(&user)
	if result.Error != nil {
		err = result.Error
		return
	}

	return
}
