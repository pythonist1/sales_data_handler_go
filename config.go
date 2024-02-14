package main

import (
	"gopkg.in/yaml.v2"
	"os"
)


type Config struct {
	FilePath string
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
