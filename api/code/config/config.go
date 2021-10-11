package config

import (
	"fmt"

	"github.com/guregu/dynamo"
	"github.com/spf13/viper"
)

// Config is global object that holds all application level variables.
var Config appConfig

type appConfig struct {
	// the shared DB ORM object
	DB *dynamo.DB
	// the error thrown be GORM when using DB ORM object
	DBErr error
	// the server port. Defaults to 8080
	ServerPort int `mapstructure:"server_port"`
	// the data source name (DSN) for connecting to the database. required.
	DSN string `mapstructure:"dsn"`
	// the API key needed to authorize to API. required.
	ApiKey string `mapstructure:"api_key"`
	// Certificate file for HTTPS
	CertFile string `mapstructure:"cert_file"`
	// Private key file for HTTPS
	KeyFile string `mapstructure:"key_file"`
}

// LoadConfig loads config from files
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("server")
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	v.SetDefault("server_port", 10000)

	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}
	return v.Unmarshal(&Config)
}
