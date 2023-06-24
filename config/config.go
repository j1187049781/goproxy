package config

import (
	"github.com/spf13/viper"
)

// Config is a struct for config
type Config struct {
	ProxyAddr  string `mapstructure:"proxy_addr"`
	ServerConfig []ServerConfig `mapstructure:"server_config"`
}

type ServerConfig struct {
	ListenAddr  string `mapstructure:"listen_addr"`
	TargetAddr string `mapstructure:"target_addr"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}