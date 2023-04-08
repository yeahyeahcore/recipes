package models

import "penahub.gitlab.yandexcloud.net/pena-services/accruals-service/pkg/postgres"

type Config struct {
	HTTP      HTTPConfiguration
	Database  DatabaseConfiguration
	FilesDist string `env:"FILES_DIST,default=./data"`
}

type HTTPConfiguration struct {
	Host string `env:"HTTP_HOST,default=localhost"`
	Port string `env:"HTTP_PORT,default=8000"`
}

type DatabaseConfiguration struct {
	DatabaseName string `env:"POSTGRES_DB_NAME,default=admin"`
	Connection   postgres.Configuration
}
