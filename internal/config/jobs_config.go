package config

import (
	"github.com/danielsantil/weather-api/internal/background"
	"github.com/danielsantil/weather-api/internal/models"
	"gorm.io/gorm"
)

func AddBackgroundJobs(db *gorm.DB, env models.Env) {
	go background.StartForecastJob(db, env)
	go background.StartWeathersJob(db, env)
}
