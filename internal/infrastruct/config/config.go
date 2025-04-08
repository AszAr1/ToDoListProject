package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type AppCfg struct {
	HttpServer `yaml:"http_server"`
	Database   `yaml:"database"`
	JwtConfig  `yaml:"jwt_config"`
}

type HttpServer struct {
	Address     string        `yaml:"address" env-default:"3000" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s" env-required:"true"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"10s" env-required:"true"`
}

type Database struct {
	Url string `yaml:"url" env-required:"true"`
}

type JwtConfig struct {
	AccessJwtSecret     string        `yaml:"accessJwtSecret" env-required:"true"`
	RefreshJwtSecret    string        `yaml:"refreshJwtSecret" env-required:"true"`
	AccessJwtExpiredIn  time.Duration `yaml:"accessJwtExpiredIn" env-required:"true"`
	RefreshJwtExpiredIn time.Duration `yaml:"refreshJwtExpiredIn" env-required:"true"`
}

func MustLoadAppConfig() *AppCfg {
	const configPath = "./config/cfg.yaml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var config AppCfg

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("Can not read config: %s", err.Error())
	}

	return &config
}
