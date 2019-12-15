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
	GRPCPort string `envconfig:"GRPC_PORT"`
	WEBPort  string `envconfig:"WEB_PORT"`
	ListenIP string `envconfig:"LISTEN_IP"`
}

type AppConfig struct {
	DbConfig
	Addr
	LogLevel string `envconfig:"LOG_LEVEL" required:"true"`
}

//DB_DSN=postgres://md:secret@postgres:54321/md_calendar;LOG_LEVEL=info;LISTEN_IP=0:0:0:0;WEB_PORT=8888
//DB_USER=md; DB_NAME=md_calendar;DB_HOST=localhost;DB_PORT=54321;DB_PASSWORD=secret;LOG_LEVEL=info;LISTEN_IP=0:0:0:0;WEB_PORT=8888
//REG_SERVICE_DB_DSN: "postgres://test:test@postgres:5432/exampledb?sslmode=disable"
//REG_SERVICE_AMQP_DSN: "amqp://guest:guest@rabbit:5672/"
//REG_SERVICE_SERVER_ADDR: ":8088"

//Addr:     conf.DbHost + ":" + conf.DbPort,
//User:     conf.DbUser,
//Password: conf.DbPassword,
//Database: conf.DbName,
