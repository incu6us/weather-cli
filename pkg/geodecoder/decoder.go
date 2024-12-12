package geodecoder

import "github.com/kelvins/geocoder"

type Decoder struct {
	apiKey string
}

func NewDecoder(apiKey string) *Decoder {
	return &Decoder{apiKey: apiKey}
}

func (d *Decoder) ToLatLon(country, city string) (float64, float64, error) {
	geocoder.ApiKey = d.apiKey
	address := geocoder.Address{
		City:    city,
		Country: country,
	}

	location, err := geocoder.Geocoding(address)
	if err != nil {
		return 0, 0, err
	}

	return location.Latitude, location.Longitude, nil
}
