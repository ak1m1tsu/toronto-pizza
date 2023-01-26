package config

import "time"

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
	}
	return config, nil
}
