package main

import (
	"github.com/danielsantil/weather-api/internal/config"
	"github.com/danielsantil/weather-api/internal/handlers"
)

func main() {
	env := config.LoadEnv()
	db := config.AddDatabase(env.ConnectionString)
	injector := handlers.Injector{DB: db}
	go config.AddBackgroundJobs(db, env)
	config.AddRouter(env.Port, injector)
}
