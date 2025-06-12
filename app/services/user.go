package services

import (
	"encoding/json"
	"fiber-starter/app/db"
	"fiber-starter/app/models"
	"fiber-starter/config"
	"fmt"
	"log"
	"time"
)

func GetUserById(id uint) (*models.User, error) {
	cacheKey := fmt.Sprintf(config.UserCacheKey, id)
	cachedData, err := db.RedisClient.Get(db.RedisCtx, cacheKey).Result()
	if err == nil {
		var user models.User
		err := json.Unmarshal([]byte(cachedData), &user)
		if err != nil {
			log.Println("Error unmarshalling cached data: ", err)
		}

		log.Println("User fetched from cache: ", user.ID)
		return &user, nil
	}

	var user models.User
	if err := db.DB.Where("id = ? AND deleted_at IS NULL", id).First(&user).Error; err != nil {
		return nil, err
	}

	// Cache the user data
	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Println("Error marshalling user data: ", err)
	}
	db.RedisClient.Set(db.RedisCtx, cacheKey, userJSON, 20*time.Minute).Err()

	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := db.DB.Where("email = ? AND deleted_at IS NULL", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUserById(id uint) error {
	if err := db.DB.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return err
	}

	if err := db.DeleteCacheByID(id, config.UserCacheKey); err != nil {
		return err
	}

	log.Println("User deleted: ", id)
	return nil
}
