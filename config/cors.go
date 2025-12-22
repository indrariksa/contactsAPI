package config

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

var Internalhost string = os.Getenv("INTERNALHOST") + ":" + os.Getenv("PORT")

var Cors = cors.Config{
	AllowMethods:     "GET,HEAD,OPTIONS,POST,PUT,DELETE",
	AllowHeaders:     "Origin, X-Requested-With, Content-Type, Accept, Authorization, Access-Control-Request-Headers, Token, Access-Control-Allow-Origin",
	ExposeHeaders:    "Content-Length",
	AllowCredentials: true,
	AllowOriginsFunc: func(origin string) bool {
		// izinkan github pages
		if origin == "https://indrariksa.github.io" {
			return true
		}
		// izinkan localhost / 127.0.0.1 port berapa pun (DEV)
		return strings.HasPrefix(origin, "http://localhost:") ||
			strings.HasPrefix(origin, "http://127.0.0.1:")
	},
}
