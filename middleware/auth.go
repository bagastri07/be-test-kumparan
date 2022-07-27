package middleware

import (
	"github.com/bagastri07/be-test-kumparan/services/config"
	"github.com/bagastri07/be-test-kumparan/utils"
	"github.com/labstack/echo/v4"
)

func VerifyKey() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			conf := config.GetConfig()

			key := c.Request().Header.Get("Api-Key")

			if key != conf.ApiKey {
				return utils.ErrNotAuthenticated
			}

			return next(c)
		}
	}
}
