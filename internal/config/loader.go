package config

import (
	"os"

	"github.com/goccy/go-yaml"
)

func Load(path string) (*ServiceConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var serviceConfig ServiceConfig

	err = yaml.NewDecoder(file).Decode(&serviceConfig)
	if err != nil {
		return nil, err
	}

	return &serviceConfig, nil
}
