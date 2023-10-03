package models

import (
	"github.com/danielsantil/weather-api/internal/models/database"
	"github.com/google/uuid"
)

type WeatherSummary struct {
	Id          uuid.UUID `json:"id"`
	CityId      int       `json:"cityId"`
	Name        string    `json:"name"`
	Country     string    `json:"country"`
	IconId      string    `json:"iconId"`
	Description string    `json:"description"`
	Temp        float64   `json:"temp"`
	FeelsLike   float64   `json:"feelsLike"`
}

func MapToWeatherSummary(weather database.Weather) WeatherSummary {
	return WeatherSummary{
		Id:          weather.ID,
		CityId:      weather.CityId,
		Name:        weather.Name,
		Country:     weather.Sys.Country,
		IconId:      weather.Conditions[0].Icon,
		Description: weather.Conditions[0].Description,
		Temp:        weather.Main.Temp,
		FeelsLike:   weather.Main.FeelsLike,
	}
}
