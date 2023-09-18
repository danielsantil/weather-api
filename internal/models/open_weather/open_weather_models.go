package open_weather

type WeatherResponse struct {
	Id            int                `json:"id"`
	Name          string             `json:"name"`
	Coord         Coordinates        `json:"coord"`
	Weather       []WeatherCondition `json:"weather"`
	Main          MainInfo           `json:"main"`
	Visibility    int                `json:"visibility"`
	Clouds        Clouds             `json:"clouds"`
	Rain          Rain               `json:"rain"`
	Snow          Snow               `json:"snow"`
	DateUtcMillis int64              `json:"dt"`
	Timezone      int                `json:"timezone"`
	Sys           System             `json:"sys"`
}

type Coordinates struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

type WeatherCondition struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type MainInfo struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	Pressure  float64 `json:"pressure"`
	Humidity  float64 `json:"humidity"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
}

type Clouds struct {
	All int `json:"all"`
}

type Rain struct {
	OneHour float64 `json:"1h"`
}

type Snow struct {
	OneHour float64 `json:"1h"`
}

type System struct {
	Country          string  `json:"country"`
	SunriseUtcMillis float64 `json:"sunrise"`
	SunsetUtcMillis  float64 `json:"sunset"`
}

type ForecastResponse struct {
	Count int        `json:"cnt"`
	List  []Forecast `json:"list"`
}

type Forecast struct {
	DateUtcMillis int64              `json:"dt"`
	Main          MainInfo           `json:"main"`
	Weather       []WeatherCondition `json:"weather"`
}
