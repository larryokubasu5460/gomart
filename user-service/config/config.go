package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBurl      string `mapstructure:"DB_URL"`
	JWTSecret  string `mapstructure:"JWT_SECRET"`
	ServerPort string `mapstructure:"SERVER_PORT"`
}

var Cfg Config

func LoadConfig() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	err = viper.Unmarshal(&Cfg)
	if err != nil {
		log.Fatalf("Unable to decode config: %v", err)
	}
}