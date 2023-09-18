package database

import (
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"time"
)

type DbBase struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Weather struct {
	DbBase
	CityId        int         `json:"cityId"`
	Name          string      `json:"name"`
	Coord         Coordinates `gorm:"embedded;embeddedPrefix:coord_" json:"coord"`
	Conditions    []Condition `gorm:"many2many:weather_conditions;constraint:OnDelete:CASCADE;" json:"conditions"`
	MainId        uuid.UUID   `gorm:"type:uuid;" json:"-"`
	Main          MainInfo    `json:"main"`
	Visibility    int         `json:"visibility"`
	Clouds        Clouds      `gorm:"embedded;embeddedPrefix:clouds_" json:"clouds"`
	Rain          Rain        `gorm:"embedded;embeddedPrefix:rain_" json:"rain"`
	Snow          Snow        `gorm:"embedded;embeddedPrefix:snow_" json:"snow"`
	DateUtcMillis int64       `json:"dtUnix"`
	Timezone      int         `json:"timezone"`
	SysId         uuid.UUID   `gorm:"type:uuid;" json:"-"`
	Sys           System      `json:"sys"`
}

type Coordinates struct {
	Longitude float64 `json:"lat"`
	Latitude  float64 `json:"lon"`
}

type Condition struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type MainInfo struct {
	DbBase    `json:"-"`
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feelsLike"`
	Pressure  float64 `json:"pressure"`
	Humidity  float64 `json:"humidity"`
	TempMin   float64 `json:"tempMin"`
	TempMax   float64 `json:"tempMax"`
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
	DbBase           `json:"-"`
	Country          string  `json:"country"`
	SunriseUtcMillis float64 `json:"sunriseDtUnix"`
	SunsetUtcMillis  float64 `json:"sunsetDtUnix"`
}

type CityForecast struct {
	DbBase
	CityId    int
	Forecasts []Forecast
}

type Forecast struct {
	DbBase
	DateUtcMillis  int64
	MainId         uuid.UUID `gorm:"type:uuid;"`
	Main           MainInfo
	Weather        []Condition `gorm:"many2many:forecast_conditions;constraint:OnDelete:CASCADE;"`
	CityForecastID uuid.UUID   `gorm:"type:uuid;constraint:OnDelete:CASCADE;"`
}

type City struct {
	DbBase
	CityId          int
	Name            string
	Country         string
	LastTimeFetched *time.Time
}

type Seeding struct {
	Name      string `gorm:"primaryKey"`
	CreatedAt time.Time
}
