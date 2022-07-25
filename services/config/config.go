package config

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppName  string `envconfig:"APP_NAME"`
	AppQuote string `envconfig:"APP_QUOTE"`
	AppPort  string `envconfig:"APP_PORT"`
	AppHost  string `envconfig:"APP_HOST"`

	DBHost         string `envconfig:"DB_HOST"`
	DBPort         string `envconfig:"DB_PORT"`
	DBUser         string `envconfig:"DB_USER"`
	DBPassword     string `envconfig:"DB_PASSWORD"`
	DBName         string `envconfig:"DB_NAME"`
	DBConnection   string `envconfig:"DB_CONNECTION"`
	DBDialTimeout  int    `envconfig:"DB_DIAL_TIMEOUT"`
	DBReadTimeout  int    `envconfig:"DB_READ_TIMEOUT"`
	DBWriteTimeout int    `envconfig:"DB_WRITE_TIMEOUT"`
	DBMaxOpenConns int    `envconfig:"DB_MAX_OPEN_CONNS"`
	DBMaxIdleConns int    `envconfig:"DB_MAX_IDLE_CONNS"`
}

var (
	once     sync.Once
	instance Config
)

func GetConfig() Config {
	once.Do(func() {
		godotenv.Load()
		err := envconfig.Process("", &instance)
		if err != nil {
			log.Fatal(err.Error())
		}
	})

	return instance
}
