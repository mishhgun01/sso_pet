package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env         string        `mapstructure:"env"`
	StoragePath string        `mapstructure:"storage_path"`
	TokenTtl    time.Duration `mapstructure:"token_ttl"`
	GRPC        GRPCConfig    `mapstructure:"grpc"`
}

type GRPCConfig struct {
	Port    int           `mapstructure:"port"`
	Timeout time.Duration `mapstructure:"timeout"`
}

func MustLoad() *Config {
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}

	cfg := new(Config)

	err = viper.Unmarshal(cfg)
	if err != nil {
		panic(err.Error())
	}

	return cfg
}
