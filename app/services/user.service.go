package services

import (
	"fiber-starter/app/db"
	"fiber-starter/app/models"
)

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := db.DB.Where("email = ? AND deleted_at IS NULL", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
