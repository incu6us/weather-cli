package provider

import (
	"github.com/incu6us/weather-cli/config"
	"github.com/incu6us/weather-cli/pkg/logger"
)

func ProvideLogger(cfg *config.Config) (*logger.Log, error) {
	formatter, err := logger.FormatterFromStaring(cfg.LogEncoder)
	if err != nil {
		return nil, err
	}
	return logger.NewLog(logger.NewHandler(cfg.LogLevel, formatter)), nil
}
