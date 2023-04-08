package postgres

type Configuration struct {
	Host         string `env:"POSTGRES_DB_HOST,default=localhost"`
	Port         int    `env:"POSTGRES_DB_PORT,default=5000"`
	User         string `env:"POSTGRES_DB_USER,required"`
	Password     string `env:"POSTGRES_DB_PASSWORD,required"`
	DatabaseName string `env:"POSTGRES_DB_NAME,required"`
	SSLMode      bool   `env:"POSTGRES_DB_SSLMODE,default=false"`
}
