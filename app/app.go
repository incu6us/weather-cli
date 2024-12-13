package app

import (
	"context"
	"time"

	"github.com/incu6us/weather-cli/app/provider"
	"github.com/incu6us/weather-cli/pkg/logger"
	"github.com/incu6us/weather-cli/service"
)

const (
	timeout = 5 * time.Second
)

type Application struct {
	appContext context.Context
	service    *service.Service
	logger     logger.Logger
}

func NewApplication(ctx context.Context, configPath string, setDebug bool) (*Application, error) {
	cfg, err := provider.ProvideConfig(configPath)
	if err != nil {
		return nil, err
	}

	log, err := provider.ProvideLogger(cfg)
	if err != nil {
		return nil, err
	}

	svc := provider.ProvideService(cfg, setDebug, timeout, log)

	app := &Application{
		appContext: ctx,
		service:    svc,
		logger:     log,
	}

	return app, nil
}

func (a *Application) Run(ctx context.Context, country, city string) error {
	a.logger.Infof(ctx, "Runnig the weather-cli application")
	return a.service.PrintWeather(ctx, country, city, timeout)
}
