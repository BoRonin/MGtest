package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port         string `env:"PORT" env-default:"3000"`
	RedisAddress string `env:"REDISADDR" env-default:"redis:6379"`
	LoggerLvl    string `env:"LOGGERLVL" env-default:"local"`
	MongoUrl     string `env:"MONGOURL" env-default:"secret"`
	MongoDB      string `env:"MONGO_INITDB_DATABASE" env-default:"profiles"`
	MongoUser    string `env:"MONGO_INITDB_ROOT_USERNAME" env-default:"admin"`
	MongoPass    string `env:"MONGO_INITDB_ROOT_PASSWORD" env-default:"password"`
}

func MustLoad() *Config {
	var cfg Config
	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		log.Fatalf("cannot read env file: %s", err)
	}
	return &cfg
}
