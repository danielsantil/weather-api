package config

import (
	"github.com/danielsantil/weather-api/internal/background"
	"gorm.io/gorm"
	"time"
)

func AddBackgroundJobs(db *gorm.DB) {
	// TODO add duration and workers count to env
	background.StartWeathersJob(db, time.Minute*15, 5)
	background.StartForecastJob(db, time.Minute*180, 5)
}
