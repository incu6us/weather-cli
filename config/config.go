package config

type Config struct {
	LogEncoder     string         `yaml:"log_encoder"`
	LogLevel       string         `yaml:"log_level"`
	Google         Google         `yaml:"google"`
	OpenWeatherMap OpenWeatherMap `yaml:"open_weather_map"`
	WeatherAPI     WeatherAPI     `yaml:"weather_api"`
	WeatherBit     WeatherBit     `yaml:"weather_bit"`
}

type Google struct {
	APIKey string `yaml:"api_key"`
}

type OpenWeatherMap struct {
	APIKey string `yaml:"api_key"`
}

type WeatherAPI struct {
	APIKey string `yaml:"api_key"`
}

type WeatherBit struct {
	APIKey string `yaml:"api_key"`
}
