package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/destafajri/system-pembayaran-spp-go-api/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
