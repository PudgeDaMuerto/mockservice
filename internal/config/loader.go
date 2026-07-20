package config

import (
	"io"

	"github.com/goccy/go-yaml"
)

func Load(file io.Reader) (*ServiceConfig, error) {
	var serviceConfig ServiceConfig

	err := yaml.NewDecoder(file).Decode(&serviceConfig)
	if err != nil {
		return nil, err
	}

	return &serviceConfig, nil
}
