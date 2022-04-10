package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host              string
	Port              string
	DB                string
	MinioRootUser     string
	MinioRootPassword string
	UserBucketName    string
}

func SetUp() (*Config, error) {
	var c Config
	err := envconfig.Process("PET", &c)
	if err != nil {
		return &Config{}, err
	}
	return &c, nil
}
