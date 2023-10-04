package config

import (
	"github.com/danielsantil/weather-api/internal/models"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func LoadEnv() models.Env {
	_ = godotenv.Load()
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("PORT env variable not set. ", err)
	}
	connectionString := os.Getenv("CONNECTION_STRING")
	if connectionString == "" {
		log.Fatal("CONNECTION_STRING env variable not set.")
	}
	weatherJobWorkers, err := strconv.Atoi(os.Getenv("WEATHER_JOB_WORKERS"))
	if err != nil {
		log.Fatal("WEATHER_JOB_WORKERS env variable not set. ", err)
	}
	weatherJobSleepDurationInMin, err := strconv.Atoi(os.Getenv("WEATHER_JOB_SLEEP_DURATION_IN_MIN"))
	if err != nil {
		log.Fatal("WEATHER_JOB_SLEEP_DURATION_IN_MIN env variable not set. ", err)
	}
	forecastJobWorkers, err := strconv.Atoi(os.Getenv("FORECAST_JOB_WORKERS"))
	if err != nil {
		log.Fatal("FORECAST_JOB_WORKERS env variable not set. ", err)
	}
	forecastJobSleepDurationInMin, err := strconv.Atoi(os.Getenv("FORECAST_JOB_SLEEP_DURATION_IN_MIN"))
	if err != nil {
		log.Fatal("FORECAST_JOB_SLEEP_DURATION_IN_MIN env variable not set. ", err)
	}
	openWeatherUrl := os.Getenv("OPEN_WEATHER_URL")
	if openWeatherUrl == "" {
		log.Fatal("OPEN_WEATHER_URL env variable not set.")
	}
	openWeatherApiKey := os.Getenv("OPEN_WEATHER_API_KEY")
	if openWeatherApiKey == "" {
		log.Fatal("OPEN_WEATHER_API_KEY env variable not set.")
	}

	return models.Env{
		Port:                          port,
		ConnectionString:              connectionString,
		WeatherJobWorkers:             weatherJobWorkers,
		ForecastJobWorkers:            forecastJobWorkers,
		WeatherJobSleepDurationInMin:  weatherJobSleepDurationInMin,
		ForecastJobSleepDurationInMin: forecastJobSleepDurationInMin,
		OpenWeatherUrl:                openWeatherUrl,
		OpenWeatherApiKey:             openWeatherApiKey,
	}
}
