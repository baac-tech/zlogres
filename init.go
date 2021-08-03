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
	TimeFieldFormat    string = time.RFC3339
	TimestampFieldName string = "response_time"

	URLTag                string = "url"
	MethodTag             string = "method"
	RequestTimeTag        string = "request_time"
	ResponseStatusCodeTag string = "status_code"
	TimeUsageTag          string = "elapsed_time"

	ContextRequestIDTag string = "transaction-id"
	ContextMessageTag   string = "message"
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
		elapsedTime := time.Since(begin).Microseconds() // <-- this should be configurable

		logger := log.Logger.Info() // <-- this should be configurable
		logger = logger.
			Str(URLTag, c.OriginalURL()).
			Str(MethodTag, c.Method()).
			Int(ResponseStatusCodeTag, statusCode).
			Str(RequestTimeTag, begin.Format(TimeFieldFormat)).
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
