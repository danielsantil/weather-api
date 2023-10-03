package main

import (
	"github.com/danielsantil/weather-api/internal/config"
	"github.com/danielsantil/weather-api/internal/handlers"
)

func main() {
	// TODO add connection string to env
	connString := "host=localhost user=postgres password=admin dbname=weather port=5432 sslmode=disable"
	db := config.AddDatabase(connString)
	injector := handlers.Injector{DB: db}
	go config.AddBackgroundJobs(db)
	port := 8000
	config.AddRouter(port, injector)
}
