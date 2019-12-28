package config

type Addr struct {
	GRPCPort string `envconfig:"GRPC_PORT" required:"true"`
	WEBPort  string `envconfig:"WEB_PORT" required:"true"`
	ListenIP string `envconfig:"LISTEN_IP" required:"true"`
}

type AppConfig struct {
	Addr
	LogLevel string `envconfig:"LOG_LEVEL" required:"true"`
}
