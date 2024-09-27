package configs

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP `yaml:"http"`
		DB   `yaml:"db"`
		JWT  `yaml:"jwt"`
	}
	HTTP struct {
		Port string `env_required:"true" yaml:"port" env:"HTTP_PORT"`
	}
	DB struct {
		Username string `env_required:"true" yaml:"username" env:"DB_USERNAME"`
		Host     string `env_required:"true" yaml:"host" env:"DB_HOST"`
		Port     string `env_required:"true" yaml:"port" env:"DB_PORT"`
		DBName   string `env_required:"true" yaml:"dbname" env:"DB_NAME"`
		SSLMode  string `env_required:"true" yaml:"ssl_mode" env:"DB_SSLMODE"`
		Password string `env_required:"true" yaml:"password" env:"DB_PASSWORD"`
		Driver   string `env_required:"true" yaml:"driver" env:"DB_DRIVER"`
	}
	JWT struct {
		JwtSecret string `env_required:"true" yaml:"jwt_secret" env:"JWT_SECRET"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./configs/config.yaml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
