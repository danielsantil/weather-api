package config

import (
	"github.com/danielsantil/weather-api/internal/background"
	"gorm.io/gorm"
	"time"
)

func AddBackgroundJobs(db *gorm.DB) {
	background.StartWeathersJob(db, time.Minute*5, 5)
}
