package provider

import (
	"os"

	yaml "gopkg.in/yaml.v3"

	"github.com/incu6us/weather-cli/config"
)

func ProvideConfig(configPath string) (*config.Config, error) {
	var cfg *config.Config
	content, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(content, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
