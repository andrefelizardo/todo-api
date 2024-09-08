package configs

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
    App   AppConfig
    DB    DBConfig
    JWT   JWTConfig
}

type AppConfig struct {
    Name string
    Env  string
}

type DBConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    Name     string
    SSLMode  string
}

type JWTConfig struct {
    Secret string
}

func LoadConfig() (*Config, error) {
var config Config

viper.SetConfigName(".env")
viper.SetConfigType("env")
viper.AddConfigPath(".")
viper.AutomaticEnv()
viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

env := viper.GetString("APP_ENV")
if env == "production" {
	viper.SetConfigName(".env.production")
}

if err := viper.ReadInConfig(); err != nil {
	log.Printf("Error reading config file", err)
	return nil, err
}

if err := viper.Unmarshal(&config); err != nil {
	log.Printf("Error unmarshalling config", err)
	return nil, err
}

return &config, nil

}