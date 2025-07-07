package config

import (
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	}
	Database struct {
		Type     string `yaml:"type"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DBName   string `yaml:"db_name"`
		SSLMode  string `yaml:"ssl_mode"`
		MaxConns int32  `yaml:"max_conns"`
		MinConns int32  `yaml:"min_conns"`
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
