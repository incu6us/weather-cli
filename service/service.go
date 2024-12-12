package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/incu6us/weather-cli/client"
	"github.com/incu6us/weather-cli/client/openweathermap"
	"github.com/incu6us/weather-cli/client/weatherapi"
	"github.com/incu6us/weather-cli/client/weatherbit"
	"github.com/incu6us/weather-cli/pkg/logger"
)

//go:generate mockgen -destination=./mock/geodecoder.go -package=mock github.com/incu6us/weather-cli/service GeoDecoder
//go:generate mockgen -destination=./mock/weather_client.go -package=mock github.com/incu6us/weather-cli/service WeatherClient

type GeoDecoder interface {
	ToLatLon(country, city string) (float64, float64, error)
}

type WeatherClient interface {
	CurrentWeather(ctx context.Context, lat, lon float64) (client.Result, error)
}

type Service struct {
	geoDecoder     GeoDecoder
	weatherClients map[WeatherClient]struct{}
	log            logger.Logger
}

func NewService(
	geoDecoder GeoDecoder,
	weatherClients []WeatherClient,
	log logger.Logger,
) *Service {
	clients := make(map[WeatherClient]struct{}, len(weatherClients))
	for _, weatherClient := range weatherClients {
		clients[weatherClient] = struct{}{}
	}

	return &Service{
		geoDecoder:     geoDecoder,
		weatherClients: clients,
		log:            log,
	}
}

func (s *Service) PrintWeather(ctx context.Context, country, city string, timeout time.Duration) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	resultCh := make(chan client.Result)
	defer close(resultCh)

	lat, lon, err := s.geoDecoder.ToLatLon(country, city)
	if err != nil {
		return err
	}

	for weatherClient := range s.weatherClients {
		go func(client WeatherClient) {
			res, err := client.CurrentWeather(ctx, lat, lon)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					return
				}
				s.log.Warnf(ctx, "failed to fetch weather: %v", err)
				return
			}

			select {
			case <-ctx.Done():
				return
			default:
			}

			resultCh <- res
		}(weatherClient)
	}

	select {
	case result := <-resultCh:
		cancel() // Cancel remaining API calls
		s.log.Infof(ctx, "Client: %v", result.ClientName())
		s.log.Infof(ctx, "Temperature: %v", parseBasics(result.ClientName(), result.Data()))
		s.log.Infof(ctx, "All Details: %v", result.Data())
	case <-time.After(timeout):
		cancel()
		s.log.Errorf(ctx, "Timeout: No API responded in time")
	}

	return nil
}

func parseBasics(clientName string, data interface{}) string {
	switch clientName {
	case openweathermap.ClientName:
		return fmt.Sprintf("%f", data.(*openweathermap.Response).Current.Temp)
	case weatherapi.ClientName:
		return fmt.Sprintf("%f", data.(*weatherapi.Response).Current.TempC)
	case weatherbit.ClientName:
		d := data.(*weatherbit.Response).Data
		if len(d) == 0 {
			return "No data"
		}
		return fmt.Sprintf("%f", d[0].Temp)
	}
	return "No data"
}
