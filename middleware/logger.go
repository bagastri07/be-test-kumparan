package middleware

import (
	"bytes"
	"io/ioutil"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
)

func Logger() echo.MiddlewareFunc {
	logger := zerolog.New(os.Stdout)
	zerolog.TimestampFieldName = "timestamp"

	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			req := c.Request()

			body := ""
			if req.Body != nil { // Read
				bodyRequestBytes, _ := ioutil.ReadAll(c.Request().Body)
				c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyRequestBytes)) // Reset

				body = string(bodyRequestBytes)
			}

			logger.Info().
				Time("Time", time.Now()).
				Str("Method", req.Method).
				Str("URI", v.URI).
				Int("status", v.Status).
				Str("body", body).
				Msg("request")

			return nil
		},
	})
}
