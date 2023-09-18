package models

import "github.com/google/uuid"

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
