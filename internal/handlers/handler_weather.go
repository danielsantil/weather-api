package handlers

import (
	"github.com/danielsantil/weather-api/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerGetWeather(c *gin.Context) {
	c.JSON(http.StatusOK, models.WeatherResponse{
		Id:   3492908,
		Name: "Santo Domingo",
		Coord: models.Coordinates{
			Longitude: 0,
			Latitude:  0,
		},
		Weather: []models.WeatherCondition{
			{
				Id:          500,
				Main:        "Rain",
				Description: "light rain",
				Icon:        "10d",
			},
		},
		Main: models.MainInfo{
			Temp:      32.42,
			FeelsLike: 39.42,
			Pressure:  1008,
			Humidity:  70,
			TempMin:   31.69,
			TempMax:   33.29,
		},
		Visibility: 0,
		Clouds: models.Clouds{
			All: 75,
		},
		Rain: models.Rain{
			OneHour: 0.12,
		},
		DateUtcMillis: 1694630735,
		Timezone:      -14400,
		Sys: models.System{
			Country:          "DO",
			SunriseUtcMillis: 1694600847,
			SunsetUtcMillis:  1694645086,
		},
	})
}
