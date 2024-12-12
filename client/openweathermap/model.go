package openweathermap

type Result struct {
	clientName string
	Response   *Response
}

func (r *Result) ClientName() string {
	return r.clientName
}

func (r *Result) Data() any {
	return r.Response
}

type Response struct {
	Current Current `json:"current"`
}

type Current struct {
	Dt         float64   `json:"dt"`
	Sunrise    float64   `json:"sunrise"`
	Sunset     float64   `json:"sunset"`
	Temp       float64   `json:"temp"`
	FeelsLike  float64   `json:"feels_like"`
	Pressure   float64   `json:"pressure"`
	Humidity   float64   `json:"humidity"`
	DewPoint   float64   `json:"dew_point"`
	Uvi        float64   `json:"uvi"`
	Clouds     float64   `json:"clouds"`
	Visibility float64   `json:"visibility"`
	WindSpeed  float64   `json:"wind_speed"`
	WindDeg    float64   `json:"wind_deg"`
	WindGust   float64   `json:"wind_gust"`
	Weather    []Weather `json:"weather"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Error struct {
	Cod        int      `json:"cod"`
	Message    string   `json:"message"`
	Parameters []string `json:"parameters,omitempty"`
}
