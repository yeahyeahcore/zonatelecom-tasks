package postgres

type PostgreSQLConfiguration struct {
	Host         string `env:"POSTGRES_HOST,default=localhost"`
	Port         int    `env:"POSTGRES_PORT,default=5000"`
	User         string `env:"POSTGRES_USER,required"`
	Password     string `env:"POSTGRES_PASSWORD,required"`
	DatabaseName string `env:"POSTGRES_NAME,required"`
	SSLMode      bool   `env:"POSTGRES_SSLMODE,default=false"`
}
