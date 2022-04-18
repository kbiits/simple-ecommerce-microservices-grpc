package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Host          string `mapstructure:"HOST"`
	Port          string `mapstructure:"PORT"`
	PostgreDsn    string `mapstructure:"DB_URL"`
	ProductSvcUrl string `mapstructure:"PRODUCT_SVC_URL"`
}

func LoadConfig() (*Config, error) {
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

	cfg := &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
