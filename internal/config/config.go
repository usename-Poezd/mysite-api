package config

import (
	"github.com/spf13/viper"
	"os"
	"sync"
)

type Config struct {
	PostgresHost		string `mapstructure:"POSTGRES_HOST"`
	PostgresUser 		string `mapstructure:"POSTGRES_USER"`
	PostgresPassword 	string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDB 			string `mapstructure:"POSTGRES_DB"`

	JwtSecret 			string `mapstructure:"JWT_SECRET"`
}
var lock = &sync.Mutex{}

var config *Config

func GetConfig(configPath string) (*Config, error) {
	if config == nil {
		lock.Lock()
		defer lock.Unlock()
		if config == nil {
			conf := new(Config)
			setFromEnv(conf)

			viper.AddConfigPath(configPath)
			viper.SetConfigName(".env")
			viper.SetConfigType("env")

			viper.AutomaticEnv()

			viper.ReadInConfig()

			if err := viper.Unmarshal(&conf); err != nil {
				return nil, err
			}

			config = conf
		}
	}

	return config, nil
}

func setFromEnv(cfg *Config) {
	cfg.PostgresHost = os.Getenv("POSTGRES_HOST")
	cfg.PostgresDB = os.Getenv("POSTGRES_DB")
	cfg.PostgresUser = os.Getenv("POSTGRES_USER")
	cfg.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
}