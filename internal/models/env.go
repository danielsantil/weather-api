package models

type Env struct {
	Port                          int
	ConnectionString              string
	WeatherJobWorkers             int
	ForecastJobWorkers            int
	WeatherJobSleepDurationInMin  int
	ForecastJobSleepDurationInMin int
	OpenWeatherUrl                string
	OpenWeatherApiKey             string
}
