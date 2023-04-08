package app

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"penahub.gitlab.yandexcloud.net/pena-services/accruals-service/internal/initialize"
	"penahub.gitlab.yandexcloud.net/pena-services/accruals-service/internal/models"
	"penahub.gitlab.yandexcloud.net/pena-services/accruals-service/internal/server"
	"penahub.gitlab.yandexcloud.net/pena-services/accruals-service/pkg/postgres"
)

func Run(config *models.Config, logger *logrus.Logger) {
	ctx, cancel := context.WithCancel(context.Background())

	pool, err := postgres.Connect(ctx, &postgres.ConnectDeps{
		Configuration: &config.Database.Connection,
		Timeout:       10 * time.Second,
	})
	if err != nil {
		logger.Fatalln("failed connection to db: ", err)
		panic(err)
	}

	repositories := initialize.NewRepositories(&initialize.RepositoriesDeps{
		Logger: logger,
		Pool:   pool,
	})

	services := initialize.NewServices(&initialize.ServicesDeps{
		Logger:       logger,
		Repositories: repositories,
	})

	controllers := initialize.NewControllers(&initialize.ControllersDeps{
		Logger:   logger,
		Services: services,
	})

	serverHTTP := server.NewHTTP(logger).Register(controllers)

	go serverHTTP.Run(&config.HTTP)

	gracefulShutdown(ctx, &gracefulShutdownDeps{
		server: serverHTTP,
		pool:   pool,
		cancel: cancel,
	})
}
