package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host              string
	Port              string
	DB                string
	MinioRootUser     string
	MinioRootPassword string
	MinioEndpoint     string
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

func LocalConfig() *Config {
	c, err := SetUp()
	if err != nil {
		log.Fatal(err.Error())
	}

	c.DB = "db"
	c.Port = "8080"
	c.Host = "0.0.0.0"
	c.MinioRootUser = "minio"
	c.MinioRootPassword = "minio123"
	c.UserBucketName = "users"
	c.MinioEndpoint = "0.0.0.0:9000"

	return c
}
