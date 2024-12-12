package service

import (
	"context"
	"testing"
	"time"

	"go.uber.org/mock/gomock"

	"github.com/incu6us/weather-cli/client"
	"github.com/incu6us/weather-cli/client/openweathermap"
	"github.com/incu6us/weather-cli/client/weatherbit"
	"github.com/incu6us/weather-cli/pkg/logger"
	"github.com/incu6us/weather-cli/service/mock"
)

func TestService_PrintWeather(t *testing.T) {
	type fields struct {
		geoDecoder     func(ctrl *gomock.Controller) GeoDecoder
		weatherClients func(ctrl *gomock.Controller) []WeatherClient
		log            logger.Logger
	}
	type args struct {
		ctx     context.Context
		country string
		city    string
		timeout time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "positive test",
			fields: fields{
				geoDecoder: func(ctrl *gomock.Controller) GeoDecoder {
					geoDecoder := mock.NewMockGeoDecoder(ctrl)
					geoDecoder.EXPECT().ToLatLon("Ukraine", "Kyiv").Return(50.450360, 30.524503, nil)
					return geoDecoder
				},
				weatherClients: func(ctrl *gomock.Controller) []WeatherClient {
					weatherClients := []WeatherClient{
						mock.NewMockWeatherClient(ctrl),
						&WeatherClientMock{},
					}
					weatherClients[0].(*mock.MockWeatherClient).EXPECT().
						CurrentWeather(gomock.Any(), 50.450360, 30.524503).
						Return(
							&TestWeatherClientResponse{
								clientName: openweathermap.ClientName,
								result: &openweathermap.Response{
									Current: openweathermap.Current{
										Temp: 20.0,
									},
								},
							},
							nil,
						)
					return weatherClients
				},
				log: logger.NewDiscardLogger(),
			},
			args: args{
				ctx:     context.Background(),
				country: "Ukraine",
				city:    "Kyiv",
				timeout: 5 * time.Second,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			s := NewService(tt.fields.geoDecoder(ctrl), tt.fields.weatherClients(ctrl), tt.fields.log)
			if err := s.PrintWeather(tt.args.ctx, tt.args.country, tt.args.city, tt.args.timeout); (err != nil) != tt.wantErr {
				t.Errorf("PrintWeather() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type TestWeatherClientResponse struct {
	clientName string
	result     any
}

func (t *TestWeatherClientResponse) ClientName() string {
	return t.clientName
}

func (t *TestWeatherClientResponse) Data() any {
	return t.result
}

type WeatherClientMock struct {
}

func (w *WeatherClientMock) CurrentWeather(ctx context.Context, _, _ float64) (client.Result, error) {
	time.Sleep(10 * time.Millisecond)
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	return &TestWeatherClientResponse{
		clientName: weatherbit.ClientName,
		result:     &weatherbit.Response{Data: []weatherbit.Data{{Temp: 20.0}}},
	}, nil
}
