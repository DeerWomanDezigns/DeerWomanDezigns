package config

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/guregu/dynamo"
	"github.com/spf13/viper"
)

// Config is global object that holds all application level variables.
var Config appConfig

type appConfig struct {
	DB              *dynamo.DB
	BackendProtocol string `mapstructure:"backend_protocol"`
	BackendDomain   string `mapstructure:"backend_domain"`
	ServerPort      int    `mapstructure:"server_port"`
	ApiKey          string `mapstructure:"api_key"`
	EtsyClientId    string `mapstructure:"etsy_client_id"`
}

func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("server")
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	v.SetDefault("backend_protocol", "https")
	v.SetDefault("backend_domain", "backend.deerwoman-dezigns")
	v.SetDefault("server_port", 81)
	v.SetDefault("api_key", GetAwsSecretKey("dwd/apiKey", "API_Key"))
	v.SetDefault("etsy_client_id", GetAwsSecretKey("dwd/etsyKeystring", "Etsy_Keystring"))

	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}

	return v.Unmarshal(&Config)
}

func GetAwsSecretKey(secretName string, key string) string {
	var respSecrets string
	if sess, serr := session.NewSession(); serr != nil {
		fmt.Println("error creating secretsmanager session", serr.Error())
	} else {
		svc := secretsmanager.New(sess, aws.NewConfig().WithRegion(`us-east-2`))
		input := &secretsmanager.GetSecretValueInput{
			SecretId: aws.String(secretName),
		}

		result, err := svc.GetSecretValue(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case secretsmanager.ErrCodeDecryptionFailure:
					fmt.Println(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())
				case secretsmanager.ErrCodeInternalServiceError:
					fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())
				case secretsmanager.ErrCodeInvalidParameterException:
					fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())
				case secretsmanager.ErrCodeInvalidRequestException:
					fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())
				case secretsmanager.ErrCodeResourceNotFoundException:
					fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
				}
			} else {
				fmt.Println(err.Error())
			}
		}
		respSecrets = *result.SecretString
	}

	var secrets map[string]json.RawMessage
	json.Unmarshal([]byte(respSecrets), &secrets)
	var secret string
	if perr := json.Unmarshal(secrets[key], &secret); perr != nil {
		fmt.Println("Parsing error for response", perr.Error())
	}
	return secret
}
