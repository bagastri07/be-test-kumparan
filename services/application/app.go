package application

import (
	"github.com/bagastri07/be-test-kumparan/constants"
	"github.com/bagastri07/be-test-kumparan/services/api/health"
	"github.com/bagastri07/be-test-kumparan/services/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	config *config.Config
	E      *echo.Echo
}

func New(config *config.Config) *App {
	app := &App{
		config: config,
		E:      echo.New(),
	}

	app.initMiddleware()
	app.initRoutes()

	return app
}

func (app *App) initRoutes() {
	// init controler
	HealthController := health.NewController(app.E)

	app.E.GET("/", HealthController.Root).Name = constants.AuthLevelPublic
}

func (app *App) initMiddleware() {
	app.E.Use(middleware.Recover())
	app.E.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))
}

func (app *App) Start() {
	app.E.HideBanner = true

	// Start server
	if err := app.E.Start(":" + app.config.AppPort); err != nil {
		app.E.Logger.Info("shutting down the server")
	}
}
