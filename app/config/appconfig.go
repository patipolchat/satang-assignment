package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type (
	AppConfig struct {
		AppName     string `mapstructure:"APP_NAME"`
		InfuraWSUrl string `mapstructure:"APP_INFURAWSURL"`
		Database    `mapstructure:",squash"`
	}

	Database struct {
		Host     string `mapstructure:"DB_HOST"`
		Port     int    `mapstructure:"DB_PORT"`
		Name     string `mapstructure:"DB_NAME"`
		Username string `mapstructure:"DB_USERNAME"`
		Password string `mapstructure:"DB_PASSWORD"`
	}
)

func (db *Database) Url() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/", db.Username, db.Password, db.Host, db.Port)
}

func ReadConfig() (AppConfig, error) {
	var config AppConfig

	viper.SetDefault("config.path", ".")
	viper.SetConfigName(".env.dev")
	viper.SetConfigType("env")
	viper.AddConfigPath(viper.GetString("config.path"))

	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, err
}
