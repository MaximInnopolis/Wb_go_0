package Test_Task_0

import (
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Database struct {
		Host     string `env:"DBHOST" envDefault:"postgresql-db"`
		Port     string `env:"DBPORT" envDefault:"5432"`
		User     string `env:"DBUSERNAME,required"`
		Dbname   string `env:"DBNAME,required"`
		Password string `env:"DBPASSWORD,required"`
	}
	Server struct {
		Host string `env:"SERVERHOST" envDefault:"0.0.0.0"`
		Port string `env:"SERVERPORT" envDefault:"8080"`
	}
	NatsStreamingCfg struct {
		ClusterId string `env:"CLUSTERID"`
		ClientId  string `env:"CLIENTID"`
	}
}

func LoadConfig() *Config {
	err := godotenv.Load(`./.env`)
	if err != nil {
		logrus.Fatalf("Failed to load .env file: %e", err)
	}
	cfg := new(Config)
	if err = env.Parse(cfg); err != nil {
		logrus.Fatal(err)
	}
	return cfg
}
