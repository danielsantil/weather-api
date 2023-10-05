package models

// Env contains all environment variables from .env file
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
