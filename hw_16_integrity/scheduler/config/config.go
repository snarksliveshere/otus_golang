package config

type RabbitConfig struct {
	RbHost     string `envconfig:"RABBIT_HOST" required:"true"`
	RbPort     string `envconfig:"RABBIT_PORT" required:"true"`
	RbUser     string `envconfig:"RABBIT_USER" required:"true"`
	RbPassword string `envconfig:"RABBIT_PASSWORD" required:"true"`
}

type Addr struct {
	GRPCPort string `envconfig:"GRPC_PORT"`
	GRPCHost string `envconfig:"GRPC_HOST"`
	WEBPort  string `envconfig:"WEB_PORT"`
	ListenIP string `envconfig:"LISTEN_IP"`
}

type AppConfig struct {
	RabbitConfig
	Addr
	LogLevel string `envconfig:"LOG_LEVEL" required:"true"`
}
