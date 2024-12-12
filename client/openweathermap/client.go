package openweathermap

import (
	"context"
	"fmt"
	"time"

	resty "github.com/go-resty/resty/v2"

	"github.com/incu6us/weather-cli/client"
)

const (
	ClientName = "openweathermap"
	url        = "https://api.openweathermap.org/data/3.0/onecall"
)

type Client struct {
	apiKey   string
	setDebug bool
	client   *resty.Client
}

func NewClient(apiKey string, setDebug bool, timeout time.Duration) *Client {
	client := resty.New()
	client.SetTimeout(timeout)

	return &Client{apiKey: apiKey, setDebug: setDebug, client: client}
}

func (c *Client) CurrentWeather(ctx context.Context, lat, lon float64) (client.Result, error) {
	resp, err := c.client.R().
		SetContext(ctx).
		SetDebug(c.setDebug).
		SetError(&Error{}).
		SetResult(&Response{}).
		SetQueryParams(
			map[string]string{
				"appid": c.apiKey,
				"lat":   fmt.Sprintf("%f", lat),
				"lon":   fmt.Sprintf("%f", lon),
			},
		).
		Get(url)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("%s", resp.Error().(*Error).Message)
	}

	data := resp.Result().(*Response)
	result := &Result{
		clientName: ClientName,
		Response:   data,
	}

	return result, nil
}
