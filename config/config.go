package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Field :: Base field for config
type Field struct {
	Key		string
	Env		string
	Default interface{}
	Required bool
}

// DatabaseConfig :: Database config
type DatabaseConfig struct {
	Dialect		string
	Host		string
	Name		string
	Username	string
	Password	string
}

// Config :: Base config
type Config struct {
	Env		string
	Name	string
	Version	string
	Port	int
	Database DatabaseConfig
}

var configMaps = []Field {
	{
		Key:		"Env",
		Env:		"ENV",
		Default: 	"development",
	},
	{
		Key:		"Name",
		Env:		"APP_NAME",
		Default:	"place-api",
	},
	{
		Key:		"Version",
		Env:		"APP_VERSION",
		Default:	"1.0.0",
	},
	{
		Key:		"Port",
		Env:		"PORT",
		Default:	3000,
	},
	{
		Key:		"Database.Dialect",
		Env:		"DB_DIALECT",
		Default:	"postgres",
		Required:	true,
	},
	{
		Key:		"Database.Host",
		Env:		"DB_HOST",
		Default:	"127.0.0.1",
		Required:	true,
	},
	{
		Key:		"Database.Name",
		Env:		"DB_NAME",
		Default:	"postgres",
		Required:	true,
	},
	{
		Key:		"Database.Username",
		Env:		"DB_USERNAME",
		Default:	"postgres",
		Required:	true,
	},
	{
		Key:		"Database.Password",
		Env:		"DB_PASSWORD",
		Default:	nil,
		Required:	true,
	},
}

// Values :: Config that will be assigned to this &Values
var Values Config

// LoadEnv :: Load environment variable
func LoadEnv() error {
	for _, field := range configMaps {
		if field.Default != nil {
			viper.SetDefault(field.Key, field.Default)
		}
	}

	viper.BindEnv("Env", "ENV")
	env := viper.GetString("Env")
	if env == "" {
		env = "development"
		viper.SetDefault("Env", env)
	}

	for _, field := range configMaps {
		viper.BindEnv(field.Key, field.Env)
	}

	err := viper.Unmarshal(&Values)
	if err != nil {
		return err
	}
	fmt.Printf("Config: %v\n", Values)

	return nil
}
