package zlogres

import (
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/gofiber/fiber/v2"
)

const ( // <-- this should be configurable
	// LogType    string = "info"
	// TimeLevel  string = "microsecond"

	// Note: This line below is not need to defined, bcoz time.RFC3339 is the default of zerolog
	TimeFieldFormat    string = time.RFC3339Nano
	TimestampFieldName string = "response_time"

	URLTag                string = "url"
	MethodTag             string = "method"
	RequestTimeTag        string = "request_time"
	ResponseStatusCodeTag string = "status_code"
	TimeUsageTag          string = "elapsed_time"

	ContextMessageTag string = "message"
)

func init() {
	zerolog.TimeFieldFormat = TimeFieldFormat
	zerolog.TimestampFieldName = TimestampFieldName
}

func New(config ...Config) fiber.Handler {
	// set default config
	cfg := configDefault(config...)

	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		// start clock
		begin := time.Now()

		// let do the request as normal
		c.Next()

		// 'baby come back to me'
		interceptedResponse := c.Response()
		statusCode := interceptedResponse.StatusCode()
		elapsedTime := getTimeDuration(time.Since(begin), cfg.ElapsedTimeUnit)

		logger := getLogLevel(cfg.LogLevel)
		logger = logger.
			Str(URLTag, c.OriginalURL()).
			Str(MethodTag, c.Method()).
			Int(ResponseStatusCodeTag, statusCode).
			Str(RequestTimeTag, begin.Format(TimeFieldFormat)). // <-- this should be configurable
			Int64(TimeUsageTag, elapsedTime)

		if reqID := c.Locals(cfg.RequestIDContextKey); reqID != nil {
			logger = logger.Str(strings.ReplaceAll(cfg.RequestIDContextKey, "-", "_"), reqID.(string))
		}

		msg := c.Locals(ContextMessageTag)
		if msg == nil {
			msg = ""
		}
		logger.Msgf("%v", msg)

		// Idk to return the same response; if you have the better way, please tell me.
		return c.Send(interceptedResponse.Body())
	}
}

func getTimeDuration(timeDuration time.Duration, unit string) int64 {
	switch unit {
	case "nano":
		return timeDuration.Nanoseconds()
	case "micro":
		return timeDuration.Microseconds()
	case "milli":
		return timeDuration.Milliseconds()
	default:
		return 0
	}
}

func getLogLevel(level string) *zerolog.Event {
	switch level {
	case "debug":
		return log.Logger.Debug()
	case "info":
		return log.Logger.Info()
	case "warn":
		return log.Logger.Warn()
	case "error":
		return log.Logger.Error()
	case "fatal":
		return log.Logger.Fatal()
	case "panic":
		return log.Logger.Panic()
	default:
		return log.Logger.Log()
	}
}
