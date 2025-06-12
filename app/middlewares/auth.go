package middlewares

import (
	"errors"
	"fiber-starter/app/models"
	"fiber-starter/app/services"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessTokenString := c.Cookies("accessToken")

		if accessTokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		// Parse access token
		accessToken, err := jwt.ParseWithClaims(accessTokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		// Check if access token is valid
		if err == nil && accessToken.Valid {
			if claims, ok := accessToken.Claims.(*models.CustomClaims); ok {
				user, err := services.GetUserById(claims.UserID)
				if err != nil {
					log.Println("Error getting user by id", err)
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
						"error": "Invalid token",
					})
				}
				c.Locals("user", user)
				return c.Next()
			}
		}

		// If the error is not a token expired error, return an unauthorized error
		if !errors.Is(err, jwt.ErrTokenExpired) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		// If the access token is expired, check the refresh token
		refreshTokenString := c.Cookies("refreshToken")
		if refreshTokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		// Get data from the expired access token
		expiredClaims := &models.CustomClaims{}
		_, _, err = new(jwt.Parser).ParseUnverified(accessTokenString, expiredClaims)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid access token",
			})
		}

		refreshToken, err := jwt.ParseWithClaims(refreshTokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		}, jwt.WithValidMethods([]string{"HS256"}))

		if err != nil || !refreshToken.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Refresh token is invalid or expired",
			})
		}

		refreshClaims, ok := refreshToken.Claims.(*models.CustomClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid refresh token claims",
			})
		}

		// Check if the user ID in the expired access token matches the user ID in the refresh token
		if expiredClaims.UserID != refreshClaims.UserID {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token mismatch",
			})
		}

		// Get the user from the database
		user, err := services.GetUserById(refreshClaims.UserID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid refresh token",
			})
		}

		// Generate new tokens
		aT, rT, err := services.GenerateTokens(user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to generate tokens",
			})
		}

		log.Println("Access token refreshed for user: ", user.ID)

		// Set the new tokens in cookies
		c.Cookie(&fiber.Cookie{
			Name:     "accessToken",
			Value:    aT,
			Expires:  time.Now().Add(15 * time.Minute),
			HTTPOnly: true,
			Secure:   true,
			SameSite: fiber.CookieSameSiteStrictMode,
		})

		// If the refresh token is expired in 7 days, refresh the token
		refreshTokenExp := (refreshToken.Claims.(*models.CustomClaims).ExpiresAt.Time.Unix() * 1000) - time.Now().Unix()
		sevenDays := 7 * 24 * 60 * 60 * 1000

		if refreshTokenExp > int64(sevenDays) {
			rT = refreshTokenString
		} else {
			log.Println("Refresh token refreshed for user: ", user.ID)
		}

		c.Cookie(&fiber.Cookie{
			Name:     "refreshToken",
			Value:    rT,
			Expires:  time.Now().Add(30 * 24 * time.Hour),
			HTTPOnly: true,
			Secure:   true,
			SameSite: fiber.CookieSameSiteStrictMode,
		})

		// Set the user in the context
		c.Locals("user", user)
		return c.Next()
	}
}
