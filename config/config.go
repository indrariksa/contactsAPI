package config

import "github.com/gofiber/fiber/v2"

var Indra = fiber.Config{
	Prefork:       false,
	CaseSensitive: true,
	StrictRouting: true,
	ServerHeader:  "INDRA",
	AppName:       "Message Router",
}
