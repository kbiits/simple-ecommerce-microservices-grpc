package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Host       string `mapstructure:"HOST"`
	Port       string `mapstructure:"PORT"`
	PostgreDsn string `mapstructure:"DB_URL"`
}

func LoadConfig() (*Config, error) {
	config := &Config{}

	viper.AddConfigPath(".")
	if strings.ToLower(os.Getenv("ENVIRONMENT")) == "development" {
		viper.SetConfigName("dev")
	} else {
		viper.SetConfigName("prod")
	}
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
