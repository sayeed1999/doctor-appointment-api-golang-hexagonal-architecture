package config

import (
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		ConnectionString string `yaml:"connectionString"`
	} `yaml:"database"`
	Redis struct {
		Address  string `yaml:"addr"`
		DB       int    `yaml:"db"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
	Mail struct {
		SmtpHost    string `yaml:"smtpHost"`
		SmtpPort    int    `yaml:"smtpPort"`
		SenderEmail string `yaml:"senderEmail"`
		Password    string `yaml:"password"`
	} `yaml:"mail"`
	Jwt struct {
		SecretKey string `yaml:"secretKey"`
	}
}

func NewConfig(configFile string) (*Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	cfg := &Config{}
	yd := yaml.NewDecoder(file)
	err = yd.Decode(cfg)

	if err != nil {
		return nil, err
	}
	return cfg, nil
}
