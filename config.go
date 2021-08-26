package zlogres

import (
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool

	// Optional. Default: "requestid"
	RequestIDContextKey string

	// Optional. Default: "info"
	LogLevel string

	// Optiona. Default: "micro". Possible Value: ["nano", "micro", "milli"]
	ElapsedTimeUnit string
}

var ConfigDefault = Config{
	Next:                nil,
	RequestIDContextKey: "requestid",
	LogLevel:            "info",
	ElapsedTimeUnit:     "micro",
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

	if cfg.LogLevel == "" {
		cfg.LogLevel = ConfigDefault.LogLevel
	}

	if cfg.ElapsedTimeUnit == "" {
		cfg.ElapsedTimeUnit = ConfigDefault.ElapsedTimeUnit
	}

	return cfg
}
