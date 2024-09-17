package configs

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

type Config struct {
	App AppConfig
	DB  DBConfig
	JWT JWTConfig
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

    viper.AutomaticEnv()

    config.DB.Host = viper.GetString("DB_HOST")
    config.DB.Port = viper.GetString("DB_PORT")
    config.DB.User = viper.GetString("DB_USER")
    config.DB.Password = viper.GetString("DB_PASSWORD")
    config.DB.Name = viper.GetString("DB_NAME")
    config.DB.SSLMode = viper.GetString("DB_SSLMODE")

	config.JWT.Secret = viper.GetString("JWT_SECRET")

	log.Info("Viper DB_HOST:", config.DB.Host)

    // Fallback para usar os.Getenv, caso o Viper não funcione
    if config.DB.Host == "" {
        config.DB.Host = os.Getenv("DB_HOST")
    }

    return &config, nil
}
