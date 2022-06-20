package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server Server
	DB     DB
}

type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type DB struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

func SetUpConfig() (*Config, error) {
	if err := setUpViper(); err != nil {
		return nil, err
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func setUpViper() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func parseEnv(cfg *Config) error {
	if err := viper.BindEnv("DB_PASSWORD"); err != nil {
		return err
	}

	cfg.DB.Password = viper.GetString("DB_PASSWORD")

	return nil
}
