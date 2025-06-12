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

/*
export function generateTokens(user: User) {
    const accessTokenPayload = {
        userId: user._id?.toString(),
        role: user.role
    }

    const refreshTokenPayload = {
        userId: user._id?.toString(),
        role: user.role
    }

    if (!JWT_SECRET) {
        throw new Error("JWT_SECRET is not defined")
    }

    const accessToken = jwt.sign(accessTokenPayload, JWT_SECRET, { expiresIn: "15m" as jwt.SignOptions["expiresIn"] })
    const refreshToken = jwt.sign(refreshTokenPayload, JWT_SECRET, { expiresIn: "30d" as jwt.SignOptions["expiresIn"] })

    return { accessToken, refreshToken }
}
*/

func GenerateTokens(user *models.User) (string, string, error) {
	accessTokenClaims := jwt.MapClaims{
		"userId": user.ID,
		"role":   user.Role,
		"exp":    time.Now().Add(15 * time.Minute).Unix(),
	}

	refreshTokenClaims := jwt.MapClaims{
		"userId": user.ID,
		"role":   user.Role,
		"exp":    time.Now().Add(30 * 24 * time.Hour).Unix(),
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
