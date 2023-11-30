package config

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	Port int `mapstructure:"port"`
	DB   DB  `mapstructure:"db"`
}

type DB struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"name"`
	SSLMode  string `mapstructure:"sslmode"`
}

func LoadConfig(path string) Config {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}
	var c Config
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatal("Failed to unmarshal config: ", err)
	}
	return c
}
