package main

import (
	"github.com/danielsantil/weather-api/internal/utils"
)

func main() {
	connString := "host=localhost user=postgres password=admin dbname=weather port=5432 sslmode=disable"
	utils.AddDatabase(connString)
	port := 8000
	utils.AddRouter(port)
}
