package config

type DbConfig struct {
	DBDriver   string `envconfig:"DB_DRIVER" default:"postgres"`
	DBHost     string `envconfig:"DB_HOST" required:"true"`
	DBPort     string `envconfig:"DB_PORT" required:"true"`
	DBUser     string `envconfig:"DB_USER" required:"true"`
	DBPassword string `envconfig:"DB_PASSWORD" required:"true"`
	DBName     string `envconfig:"DB_NAME" required:"true"`
}

type Addr struct {
	GRPCPort string `envconfig:"GRPC_PORT" required:"true"`
	WEBPort  string `envconfig:"WEB_PORT" required:"true"`
	ListenIP string `envconfig:"LISTEN_IP"`
}

type AppConfig struct {
	DbConfig
	Addr
	LogLevel string `envconfig:"LOG_LEVEL" required:"true"`
}
