package config

import "os"

const (
	UserCacheKey = "user:%d"
)

func GetRedisAddr() string {
	addr := os.Getenv("REDIS_URL")
	if addr == "" {
		return "localhost:6379"
	}
	return addr
}

func GetRedisPassword() string {
	password := os.Getenv("REDIS_PASSWORD")
	if password == "" {
		return ""
	}
	return password
}
