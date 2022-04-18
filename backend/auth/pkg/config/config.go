package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port                string `mapstructure:"PORT"`
	Host                string `mapstructure:"HOST"`
	PostgreDsn          string `mapstructure:"DB_URL"`
	JwtSecretKey        string `mapstructure:"JWT_SECRET_KEY"`
	JwtExpirationMinute uint64 `mapstructure:"JWT_EXPIRATION_MINUTES"`
}

func LoadConfig() (*Config, error) {
	var cfg *Config

	viper.AddConfigPath(".")
	if os.Getenv("ENVIRONMENT") == "DEVELOPMENT" {
		viper.SetConfigName("dev")
	} else {
		viper.SetConfigName("prod")
	}
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
