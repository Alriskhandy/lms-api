package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppPort string
	Env     string
	DBUrl   string
	JwtKey  string
}

func Load() (*Config, error) {
	viper.SetDefault("APP_PORT", "8080")
	viper.AutomaticEnv()
	// optionally read config file
	c := &Config{
		AppPort: viper.GetString("APP_PORT"),
		Env:     viper.GetString("APP_ENV"),
		DBUrl:   viper.GetString("DATABASE_URL"),
		JwtKey:  viper.GetString("JWT_SECRET"),
	}
	return c, nil
}
