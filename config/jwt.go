package config

import (
	"os"
	"strconv"
	"time"
)

func JWTSecret() []byte {
	sec := os.Getenv("JWT_SECRET")
	if sec == "" {
		panic("JWT_SECRET is not set")
	}
	return []byte(sec)
}

func JWTExpires() time.Duration {
	hoursStr := os.Getenv("JWT_EXPIRES_HOURS")
	if hoursStr == "" {
		return 24 * time.Hour
	}
	h, err := strconv.Atoi(hoursStr)
	if err != nil || h <= 0 {
		return 24 * time.Hour
	}
	return time.Duration(h) * time.Hour
}
