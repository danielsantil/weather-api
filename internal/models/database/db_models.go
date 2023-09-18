package database

import (
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"time"
)

type DbBase struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Weather struct {
	DbBase
	CityId        int
	Name          string
	Coord         Coordinates `gorm:"embedded;embeddedPrefix:coord_"`
	Conditions    []Condition `gorm:"many2many:weather_conditions;constraint:OnDelete:CASCADE;"`
	MainId        uuid.UUID   `gorm:"type:uuid;"`
	Main          MainInfo
	Visibility    int
	Clouds        Clouds `gorm:"embedded;embeddedPrefix:clouds_"`
	Rain          Rain   `gorm:"embedded;embeddedPrefix:rain_"`
	Snow          Snow   `gorm:"embedded;embeddedPrefix:snow_"`
	DateUtcMillis int64
	Timezone      int
	SysId         uuid.UUID `gorm:"type:uuid;"`
	Sys           System
}

type Coordinates struct {
	Longitude float64
	Latitude  float64
}

type Condition struct {
	Id          int
	Main        string
	Description string
	Icon        string
}

type MainInfo struct {
	DbBase
	Temp      float64
	FeelsLike float64
	Pressure  float64
	Humidity  float64
	TempMin   float64
	TempMax   float64
}

type Clouds struct {
	All int
}

type Rain struct {
	OneHour float64
}

type Snow struct {
	OneHour float64
}

type System struct {
	DbBase
	Country          string
	SunriseUtcMillis float64
	SunsetUtcMillis  float64
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
