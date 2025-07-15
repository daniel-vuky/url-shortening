package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	}
	Database struct {
		Type     string `yaml:"type" mapstructure:"type"`
		User     string `yaml:"user" mapstructure:"user"`
		Password string `yaml:"password" mapstructure:"password"`
		Host     string `yaml:"host" mapstructure:"host"`
		Port     string `yaml:"port" mapstructure:"port"`
		DBName   string `yaml:"db_name" mapstructure:"db_name"`
		SSLMode  string `yaml:"ssl_mode" mapstructure:"ssl_mode"`
		MaxConns int    `yaml:"max_conns" mapstructure:"max_conns"`
		MinConns int    `yaml:"min_conns" mapstructure:"min_conns"`
	}
}

// LoadConfig loads the config from the env.yaml file
func LoadConfig() (*Config, error) {
	var config Config
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
