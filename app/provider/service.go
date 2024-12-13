package provider

import (
	"time"

	"github.com/incu6us/weather-cli/client/openweathermap"
	"github.com/incu6us/weather-cli/client/weatherapi"
	"github.com/incu6us/weather-cli/client/weatherbit"
	"github.com/incu6us/weather-cli/config"
	"github.com/incu6us/weather-cli/pkg/geodecoder"
	"github.com/incu6us/weather-cli/pkg/logger"
	"github.com/incu6us/weather-cli/service"
)

func ProvideService(cfg *config.Config, setDebug bool, timeout time.Duration, log *logger.Log) *service.Service {
	svc := service.NewService(
		geodecoder.NewDecoder(cfg.Google.APIKey),
		[]service.WeatherClient{
			openweathermap.NewClient(cfg.OpenWeatherMap.APIKey, setDebug, timeout),
			weatherapi.NewClient(cfg.WeatherAPI.APIKey, setDebug, timeout),
			weatherbit.NewClient(cfg.WeatherBit.APIKey, setDebug, timeout),
		},
		log,
	)
	return svc
}
