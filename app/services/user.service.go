package services

import (
	"fiber-starter/app/db"
	"fiber-starter/app/models"
)

func GetUserById(id uint) (*models.User, error) {
	var user models.User
	if err := db.DB.Where("id = ? AND deleted_at IS NULL", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := db.DB.Where("email = ? AND deleted_at IS NULL", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
