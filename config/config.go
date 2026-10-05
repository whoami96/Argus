package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server    ServerConfig      `yaml:"server"`
	Instances []InstancesConfig `yaml:"instances"`
}

type ServerConfig struct {
	Port         int    `yaml:"port"`
	PollInterval string `yaml:"poll_interval"`
}

type InstancesConfig struct {
	Name  string `yaml:"name"`
	Env   string `yaml:"env"`
	Type  string `yaml:"type"`
	URL   string `yaml:"url"`
	Token string `yaml:"token"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
