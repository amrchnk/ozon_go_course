package config

import (
	"errors"
	"github.com/spf13/viper"
)

type Config struct {
	TelegramToken     string
	DBPath            string `mapstructure:"db_file"`

	//Messages Messages
}



func Init() (*Config, error) {
	var cfg Config

	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseEnv(cfg *Config) error {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	cfg.TelegramToken = viper.GetString("token")
	cfg.DBPath = viper.GetString("db_file")
	if cfg.TelegramToken == ""  {
		return errors.New("Переменные окружения не найдены")
	}
	return nil
}


