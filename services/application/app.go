package application

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bagastri07/be-test-kumparan/constants"
	"github.com/bagastri07/be-test-kumparan/database"
	"github.com/bagastri07/be-test-kumparan/infrastructure"
	midd "github.com/bagastri07/be-test-kumparan/middleware"
	"github.com/bagastri07/be-test-kumparan/services/api/article"
	"github.com/bagastri07/be-test-kumparan/services/api/author"
	"github.com/bagastri07/be-test-kumparan/services/api/health"
	"github.com/bagastri07/be-test-kumparan/services/config"
	"github.com/bagastri07/be-test-kumparan/utils"
	"github.com/bagastri07/be-test-kumparan/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	nrecho "github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
)

type App struct {
	config    *config.Config
	E         *echo.Echo
	DNmanager *database.Manager
}

func New(config *config.Config) *App {
	app := &App{
		config:    config,
		E:         echo.New(),
		DNmanager: &database.Manager{},
	}

	app.initMiddleware()
	app.initDatabase()
	app.initValidator()
	app.initRoutes()

	return app
}

func (app *App) initRoutes() {
	//init repositories
	authorRepository := author.NewRepository()
	articleRepository := article.NewRepository()

	// init services
	authorService := author.NewService(app.DNmanager.DB, authorRepository)
	articleService := article.NewService(app.DNmanager.DB, articleRepository, authorRepository)

	// init controler
	healthController := health.NewController(app.E)
	authorController := author.NewController(authorService)
	articleController := article.NewController(articleService)

	app.E.GET("/", healthController.Root).Name = constants.AuthLevelPublic

	author := app.E.Group("/authors")
	author.POST("", authorController.HandleCreateAuthor, midd.VerifyKey()).Name = constants.AuthLevelPassword

	article := app.E.Group("/articles")
	article.POST("", articleController.HandleCreateArticle, midd.VerifyKey()).Name = constants.AuthLevelPassword
	article.GET("", articleController.HandleGetArticlesPagination, midd.VerifyKey()).Name = constants.AuthLevelPassword
}

func (app *App) initMiddleware() {
	newRelic, err := infrastructure.NewRelicApm(app.config)
	if err != nil {
		panic(err)
	}

	app.E.Use(nrecho.Middleware(newRelic))
	app.E.Use(middleware.Recover())
	app.E.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))
	app.E.Use(middleware.Logger())
}

func (app *App) initDatabase() {
	db, err := database.NewConnection(*app.config)
	if err != nil {
		panic(err)
	}

	app.DNmanager.DB = db
}

func (app *App) initValidator() {
	validator.Init(app.E)
}

func (app *App) Start() {
	app.E.HideBanner = true

	// Start server
	go func() {
		if err := app.E.Start(":" + app.config.AppPort); err != nil {
			app.E.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	// Graceful Shutdown see: https://echo.labstack.com/cookbook/graceful-shutdown
	// Make sure no more in-flight request within 10seconds timeout
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	utils.Logger.Info().Strs("tags", []string{"application", "Shutdown"}).Msg("Shutting down the server gracefully")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.E.Shutdown(ctx); err != nil {
		app.E.Logger.Fatal(err)
	}
}
