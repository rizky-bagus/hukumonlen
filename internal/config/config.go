// Package config provides configuration for the project.
package config

import (
	"log"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

// Config holds configuration for the project.
type Config struct {
	Port  string `env:"PORT,default=8080"`
	Env   string `env:"ENV,default=development"`
	MySQL MySQLConfig
}

// MySQLConfig holds MySQL configurations.
type MySQLConfig struct {
	Host     string `env:"MYSQL_HOST,default=localhost"`
	Port     string `env:"MYSQL_PORT,default=3306"`
	User     string `env:"MYSQL_USER,default=mysql"`
	Protocol string `env:"MYSQL_PROTOCOL,default=tcp"`
	Password string `env:"MYSQL_PASSWORD,required"`
	Database string `env:"MYSQL_DATABASE"`
}

// NewConfig creates an instance of Config.
// It needs the path of the env file to be used.
func NewConfig(env string) (*Config, error) {
	var config Config

	if err := godotenv.Load(env); err != nil {
		log.Println(errors.Wrap(err, "[NewConfig] error reading .env file, defaulting to OS environment variables"))
	}

	if err := envdecode.Decode(&config); err != nil {
		return nil, errors.Wrap(err, "[NewConfig] error decoding env")
	}

	return &config, nil
}
