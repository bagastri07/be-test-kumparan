package health

import (
	"net/http"

	"github.com/bagastri07/be-test-kumparan/models"
	"github.com/bagastri07/be-test-kumparan/services/config"
	"github.com/labstack/echo/v4"
)

type HealthController struct {
	e *echo.Echo
}

func NewController(e *echo.Echo) *HealthController {
	return &HealthController{
		e: e,
	}
}

func (ctl *HealthController) Root(c echo.Context) error {
	conf := config.GetConfig()

	resp := models.MessageResponse{
		Message: conf.AppQuote,
	}

	return c.JSON(http.StatusOK, resp)
}
