package config

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Host              string `mapstructure:"HOST"`
	Port              string `mapstructure:"PORT"`
	AuthServiceUrl    string `mapstructure:"AUTH_SERVICE_URL"`
	OrderServiceUrl   string `mapstructure:"ORDER_SERVICE_URL"`
	ProductServiceUrl string `mapstructure:"PRODUCT_SERVICE_URL"`
}

func LoadConfig() (*Config, error) {

	viper.AddConfigPath(".")
	if strings.ToLower(os.Getenv("ENVIRONMENT")) == "development" {
		viper.SetConfigName("dev")
	} else {
		viper.SetConfigName("prod")
	}
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var c *Config
	err = viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
