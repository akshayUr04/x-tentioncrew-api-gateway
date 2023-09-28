package config

import (
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type Config struct {
	Port        string `mapstructure:"PORT"`
	UserSvcUrl  string `mapstructure:"USER_SERVICE_URL"`
	SERVICE2URL string `mapstructure:"SERVICE2_URL"`
}

var envs = []string{"PORT", "USER_SERVICE_URL", "SERVICE2_URL"}

func LoadConfig() (cfg Config, err error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	if err := validator.New().Struct(&config); err != nil {
		return config, err
	}

	return config, nil
}
