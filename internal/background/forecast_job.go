package background

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/danielsantil/weather-api/internal/models"
	"github.com/danielsantil/weather-api/internal/models/database"
	"github.com/danielsantil/weather-api/internal/models/open_weather"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

// StartForecastJob starts job to retrieve forecast data
func StartForecastJob(db *gorm.DB, env models.Env) {
	sleepDuration := time.Minute * time.Duration(env.ForecastJobSleepDurationInMin)
	log.Printf("Starting job to retrieve Open Weather data for FORECASTS. "+
		"Using %d workers\n every %v", env.ForecastJobWorkers, sleepDuration)

	now := time.Now()
	ticker := time.NewTicker(sleepDuration)
	for ; ; <-ticker.C { // start job immediately
		log.Printf("Retrieving cities. Time: %v...\n", time.Now())
		var citiesList []database.City
		db.Order("name").Find(&citiesList)
		if len(citiesList) == 0 {
			continue
		}

		listOfCities := getCitiesForWorkers(citiesList, env.ForecastJobWorkers)
		wg := &sync.WaitGroup{}
		for _, cities := range listOfCities {
			wg.Add(1)
			go updateForecast(db, cities, wg, now, sleepDuration, env)
		}
		wg.Wait()
		log.Println("=====================")
	}
}

func updateForecast(db *gorm.DB, cities []database.City, wg *sync.WaitGroup,
	now time.Time, sleepDuration time.Duration, env models.Env) {
	defer wg.Done()

	for _, city := range cities {
		response, err := fetchForecastData(city, env.OpenWeatherUrl, env.OpenWeatherApiKey)
		if err != nil {
			log.Println(err)
			continue
		}

		var existingForecast database.CityForecast
		exists := db.Where("city_id = ?", city.CityId).First(&existingForecast)

		if exists.RowsAffected == 0 { // create city forecast
			dbForecast := mapToDbForecast(response, city.CityId)
			result := db.Create(&dbForecast)
			if result.Error != nil {
				log.Printf("Error inserting forecast data for city %s\n", city.Name)
			}
		} else {
			timeDiff := now.Sub(existingForecast.UpdatedAt) // skip if time diff is lower than sleep duration (3 hours)
			if timeDiff < sleepDuration {
				continue
			}

			// city exists and last update was more than 3 hours ago
			log.Printf("Forecast exists for city %s... upserting entries\n", city.Name)
			existingForecast.Forecasts = mapForecasts(response.List)
			db.Save(&existingForecast)
		}

	}
}

func fetchForecastData(city database.City, baseUrl string, apiKey string) (open_weather.ForecastResponse, error) {
	log.Printf("Fetching forecast data for city: %d - %s\n", city.CityId, city.Name)
	httpClient := http.Client{}
	res, hErr := httpClient.Get(fmt.Sprintf("%s/forecast?id=%d&appid=%s", baseUrl, city.CityId, apiKey))
	if hErr != nil {
		return open_weather.ForecastResponse{},
			errors.New(fmt.Sprintf("Http error for forecast. City %s: %v", city.Name, hErr))
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing response body", err)
		}
	}(res.Body)

	data, rErr := io.ReadAll(res.Body)
	if rErr != nil {
		return open_weather.ForecastResponse{},
			errors.New(fmt.Sprintf("Reader error. City %s: %v", city.Name, rErr))
	}

	var forecast open_weather.ForecastResponse
	mErr := json.Unmarshal(data, &forecast)
	if mErr != nil {
		return open_weather.ForecastResponse{},
			errors.New(fmt.Sprintf("Unmarshaler error. City %s: %v", city.Name, rErr))
	}

	return forecast, nil
}

func mapToDbForecast(fr open_weather.ForecastResponse, cityId int) database.CityForecast {
	return database.CityForecast{
		CityId:    cityId,
		Forecasts: mapForecasts(fr.List),
	}
}

func mapForecasts(forecastsArray []open_weather.Forecast) []database.Forecast {
	forecastsRes := make([]database.Forecast, len(forecastsArray))
	for i, f := range forecastsArray {
		forecastsRes[i] = database.Forecast{
			DateUtcMillis: f.DateUtcMillis,
			Main: database.MainInfo{
				Temp:      f.Main.Temp,
				FeelsLike: f.Main.FeelsLike,
				Pressure:  f.Main.Pressure,
				Humidity:  f.Main.Humidity,
				TempMin:   f.Main.TempMin,
				TempMax:   f.Main.TempMax,
			},
			Weather: mapConditions(f.Weather),
		}
	}
	return forecastsRes
}
