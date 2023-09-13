package utils

import (
	"github.com/danielsantil/weather-api/internal/models/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func AddDatabase(connString string) *gorm.DB {
	db, dbErr := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if dbErr != nil {
		log.Fatal("Can't connect to database: ", dbErr)
	}

	mErr := db.AutoMigrate(&database.Weather{}, &database.System{},
		&database.CityForecast{}, &database.Forecast{})
	if mErr != nil {
		log.Fatal("Can't migrate database: ", mErr)
	}

	return db
}
