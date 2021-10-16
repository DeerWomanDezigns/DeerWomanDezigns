package config

import (
	"fmt"

	"github.com/guregu/dynamo"
	"github.com/spf13/viper"
)

// Config is global object that holds all application level variables.
var Config appConfig

type appConfig struct {
	DB         *dynamo.DB
	ServerPort int    `mapstructure:"server_port"`
	ApiKey     string `mapstructure:"api_key"`
}

// LoadConfig loads config from files
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("server")
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	v.SetDefault("server_port", 80)

	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}
	return v.Unmarshal(&Config)
}
