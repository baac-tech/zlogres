package zlogres

import (
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool

	// Optional. Default: "requestid"
	RequestIDContextKey string
}

var ConfigDefault = Config{
	Next:                nil,
	RequestIDContextKey: "requestid",
}

func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	// set default
	if cfg.RequestIDContextKey == "" {
		cfg.RequestIDContextKey = ConfigDefault.RequestIDContextKey
	}

	return cfg
}
