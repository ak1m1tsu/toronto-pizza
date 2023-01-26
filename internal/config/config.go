package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Token struct {
	PrivateKey string        `yaml:",omitempty"`
	PublicKey  string        `yaml:",omitempty"`
	ExpiresIn  time.Duration `yaml:"expires_in"`
	MaxAge     int           `yaml:"max_age"`
}

type Config struct {
	Port         string `yaml:",omitempty"`
	MongoURL     string `yaml:",omitempty"`
	AccessToken  Token  `yaml:"access"`
	RefreshToken Token  `yaml:"refresh"`
}

var config *Config

func GetConfig(path string) (*Config, error) {
	if config == nil {
		yfile, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(yfile, &config)
		if err != nil {
			return nil, err
		}
		config.MongoURL = os.Getenv("MONGO_URL")
		config.Port = os.Getenv("PORT")
		config.AccessToken.PrivateKey = os.Getenv("ACCESS_PRIVATE_KEY")
		config.AccessToken.PublicKey = os.Getenv("ACCESS_PUBLIC_KEY")
		config.RefreshToken.PrivateKey = os.Getenv("REFRESH_PRIVATE_KEY")
		config.RefreshToken.PublicKey = os.Getenv("REFRESH_PUBLIC_KEY")
	}
	return config, nil
}
