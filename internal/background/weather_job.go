package background

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/danielsantil/weather-api/internal/models/database"
	"github.com/danielsantil/weather-api/internal/models/open_weather"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

func StartWeathersJob(db *gorm.DB, sleepDuration time.Duration, workersCount int) {
	log.Printf("Starting job to retrieve Open Weather data for WEATHERS. "+
		"Using %d workers\n every %v", workersCount, sleepDuration)

	ticker := time.NewTicker(sleepDuration)
	for ; ; <-ticker.C { // start job immediately
		log.Printf("Retrieving cities. Time: %v...\n", time.Now())
		var citiesList []database.City
		db.Order("name").Find(&citiesList)
		if len(citiesList) == 0 {
			continue
		}

		listOfCities := getCitiesForWorkers(citiesList, workersCount)
		wg := &sync.WaitGroup{}
		for _, cities := range listOfCities {
			wg.Add(1)
			go updateWeather(db, cities, wg)
		}
		wg.Wait()
		log.Println("=====================")
	}
}

func updateWeather(db *gorm.DB, cities []database.City, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, city := range cities {
		response, err := fetchWeatherData(city)
		if err != nil {
			log.Println(err)
			continue
		}

		dbWeather := mapToDbWeather(response)

		// fetch weather where city_id and date_utc equals dbWeather, if exists, skip insert
		var existingWeather database.Weather
		exists := db.Where("city_id = ? AND date_utc_millis = ?", dbWeather.CityId, dbWeather.DateUtcMillis).
			First(&existingWeather)

		if exists.RowsAffected > 0 {
			log.Printf("Skipping weather insert for city %s and dt %v\n", city.Name, dbWeather.DateUtcMillis)
			continue
		}

		result := db.Create(&dbWeather)
		if result.Error != nil {
			log.Printf("Error inserting weather data for city %s\n", city.Name)
		}

		now := time.Now()
		city.LastTimeFetched = &now
		db.Save(&city)
	}
}

func mapToDbWeather(w open_weather.WeatherResponse) database.Weather {
	return database.Weather{
		CityId: w.Id,
		Name:   w.Name,
		Coord: database.Coordinates{
			Longitude: w.Coord.Longitude,
			Latitude:  w.Coord.Latitude,
		},
		Conditions: mapConditions(w.Weather),
		Main: database.MainInfo{
			Temp:      w.Main.Temp,
			FeelsLike: w.Main.FeelsLike,
			Pressure:  w.Main.Pressure,
			Humidity:  w.Main.Humidity,
			TempMin:   w.Main.TempMin,
			TempMax:   w.Main.TempMax,
		},
		Visibility:    w.Visibility,
		Clouds:        database.Clouds{All: w.Clouds.All},
		Rain:          database.Rain{OneHour: w.Rain.OneHour},
		Snow:          database.Snow{OneHour: w.Snow.OneHour},
		DateUtcMillis: w.DateUtcMillis,
		Timezone:      w.Timezone,
		Sys: database.System{
			Country:          w.Sys.Country,
			SunriseUtcMillis: w.Sys.SunriseUtcMillis,
			SunsetUtcMillis:  w.Sys.SunsetUtcMillis,
		},
	}
}

func mapConditions(conditionsArray []open_weather.WeatherCondition) []database.Condition {
	conditionsRes := make([]database.Condition, len(conditionsArray))
	for i, cond := range conditionsArray {
		conditionsRes[i] = database.Condition{
			Id:          cond.Id,
			Main:        cond.Main,
			Description: cond.Description,
			Icon:        cond.Icon,
		}
	}
	return conditionsRes
}

func getCitiesForWorkers(cities []database.City, workerCount int) [][]database.City {
	citiesTotal := len(cities)
	newLength := citiesTotal / workerCount
	listOfCities := make([][]database.City, workerCount)

	for i := 0; i < workerCount; i++ {
		start := i * newLength
		end := start + newLength
		if i+1 == workerCount { // end of loop, add remaining cities to last array index
			listOfCities[i] = cities[start:]
		} else {
			listOfCities[i] = cities[start:end]
		}
	}

	return listOfCities
}

func fetchWeatherData(city database.City) (open_weather.WeatherResponse, error) {
	log.Printf("Fetching weather data for city: %d - %s\n", city.CityId, city.Name)
	httpClient := http.Client{}
	// TODO add baseUrl, api key to env
	baseUrl := "https://api.openweathermap.org/data/2.5"
	res, hErr := httpClient.Get(fmt.Sprintf("%s/weather?id=%d&appid=%s", baseUrl, city.CityId, "e9d23eedd545de00620c0c542ffb66e1"))
	if hErr != nil {
		return open_weather.WeatherResponse{},
			errors.New(fmt.Sprintf("Http error for weather. City %s: %v", city.Name, hErr))
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Error closing response body", err)
		}
	}(res.Body)

	data, rErr := io.ReadAll(res.Body)
	if rErr != nil {
		return open_weather.WeatherResponse{},
			errors.New(fmt.Sprintf("Reader error. City %s: %v", city.Name, rErr))
	}

	var weather open_weather.WeatherResponse
	mErr := json.Unmarshal(data, &weather)
	if mErr != nil {
		return open_weather.WeatherResponse{},
			errors.New(fmt.Sprintf("Unmarshaler error. City %s: %v", city.Name, rErr))
	}

	return weather, nil
}
