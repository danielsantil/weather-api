package models

type WeatherSummary struct {
	CityId      int     `json:"cityId"`
	Name        string  `json:"name"`
	Country     string  `json:"country"`
	IconId      string  `json:"iconId"`
	Description string  `json:"description"`
	Temp        float64 `json:"temp"`
	FeelsLike   float64 `json:"feelsLike"`
}
