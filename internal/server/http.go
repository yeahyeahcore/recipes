package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"

	"github.com/yeahyeahcore/recipes/internal/initialize"
	"github.com/yeahyeahcore/recipes/internal/models"
)

type HTTP struct {
	logger *logrus.Logger
	server *http.Server
	echo   *echo.Echo
}

func NewHTTP(logger *logrus.Logger) *HTTP {
	echo := echo.New()

	echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${method}]: ${uri} ${status} ${time_rfc3339}\n",
	}))
	echo.Use(middleware.Recover())

	return &HTTP{
		echo:   echo,
		logger: logger,
		server: &http.Server{
			Handler:        echo,
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}
}

func (receiver *HTTP) Register(controllers *initialize.Controllers, dist string) *HTTP {
	receiver.echo.Static("/", dist)

	recipeGroup := receiver.echo.Group("/recipe")
	{
		recipeGroup.GET("/all", controllers.RecipeController.GetAll)
	}

	return receiver
}

func (receiver *HTTP) Run(config *models.HTTPConfiguration) {
	defer func() {
		if err := recover(); err != nil {
			receiver.logger.Errorf("run http recover: %s", err)
		}
	}()

	connectionString := fmt.Sprintf("%s:%s", config.Host, config.Port)
	startServerMessage := fmt.Sprintf("Starting HTTP Server on: %s", connectionString)

	receiver.logger.Infoln(startServerMessage)

	if err := receiver.Listen(connectionString); err != nil && err != http.ErrServerClosed {
		receiver.logger.Infoln("HTTP Listen error:", err)
		panic(err)
	}
}

func (receiver *HTTP) Listen(address string) error {
	receiver.server.Addr = address

	return receiver.server.ListenAndServe()
}

func (receiver *HTTP) Stop(ctx context.Context) {
	receiver.echo.Shutdown(ctx)
	receiver.server.Shutdown(ctx)
}
