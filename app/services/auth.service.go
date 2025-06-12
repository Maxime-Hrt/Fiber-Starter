package services

import (
	"errors"
	"fiber-starter/app/db"
	"fiber-starter/app/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func SignUp(name, email, password string) (*models.User, error) {
	// Check if user already exists
	_, err := GetUserByEmail(email)
	if err == nil {
		return nil, errors.New("user already exists")
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
		Role:     models.RoleUser,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func SignIn(email, password string) (*models.User, error) {
	user, err := GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := user.ComparePassword(password); err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func GenerateTokens(user *models.User) (string, string, error) {
	accessTokenClaims := models.CustomClaims{
		UserID: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}

	refreshTokenClaims := models.CustomClaims{
		UserID: user.ID,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims).SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims).SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
