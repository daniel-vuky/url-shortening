package config

import (
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	}
}

// LoadConfig loads the config from the env.yaml file
func LoadConfig() (*Config, error) {
	var config Config
	var once sync.Once
	var configErr error
	once.Do(func() {
		viper.SetConfigName("env")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		configErr = viper.ReadInConfig()
		if configErr != nil {
			return
		}

		configErr = viper.Unmarshal(&config)
		if configErr != nil {
			return
		}
	})

	return &config, configErr
}
