package weatherbit

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
	Data []Data `json:"data"`
}

type Data struct {
	WindCdir     string   `json:"wind_cdir"`
	Rh           float64  `json:"rh"`
	Pod          string   `json:"pod"`
	Lon          float64  `json:"lon"`
	Pres         float64  `json:"pres"`
	Timezone     string   `json:"timezone"`
	ObTime       string   `json:"ob_time"`
	CountryCode  string   `json:"country_code"`
	Clouds       float64  `json:"clouds"`
	Vis          float64  `json:"vis"`
	WindSpd      float64  `json:"wind_spd"`
	Gust         float64  `json:"gust"`
	WindCdirFull string   `json:"wind_cdir_full"`
	AppTemp      float64  `json:"app_temp"`
	StateCode    string   `json:"state_code"`
	TS           float64  `json:"ts"`
	HAngle       float64  `json:"h_angle"`
	Dewpt        float64  `json:"dewpt"`
	Weather      Weather  `json:"weather"`
	Uv           float64  `json:"uv"`
	Aqi          float64  `json:"aqi"`
	Station      string   `json:"station"`
	Sources      []string `json:"sources"`
	WindDir      float64  `json:"wind_dir"`
	ElevAngle    float64  `json:"elev_angle"`
	Datetime     string   `json:"datetime"`
	Precip       float64  `json:"precip"`
	Ghi          float64  `json:"ghi"`
	Dni          float64  `json:"dni"`
	Dhi          float64  `json:"dhi"`
	SolarRad     float64  `json:"solar_rad"`
	CityName     string   `json:"city_name"`
	Sunrise      string   `json:"sunrise"`
	Sunset       string   `json:"sunset"`
	Temp         float64  `json:"temp"`
	Lat          float64  `json:"lat"`
	Slp          float64  `json:"slp"`
}

type Weather struct {
	Icon        string `json:"icon"`
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
