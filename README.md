## Weather API

Based in Open Weather API for current weather data and forecast data.

Make sure to include a **.env** file in the root directory, containing the following variables:

| Variable                           | Description                                                     |
|------------------------------------|-----------------------------------------------------------------|
| PORT                               | Localhost port to access the API                                |
| CONNECTION_STRING                  | Postgres connection string                                      |
| WEATHER_JOB_WORKERS                | Number of workers for weather job                               |
| FORECAST_JOB_WORKERS               | Number of workers for forecast job                              |
| WEATHER_JOB_SLEEP_DURATION_IN_MIN  | Duration in minutes to wait between each weather job execution  |
| FORECAST_JOB_SLEEP_DURATION_IN_MIN | Duration in minutes to wait between each forecast job execution |
| OPEN_WEATHER_URL                   | OpenWeather API URL                                             |
| OPEN_WEATHER_API_KEY               | OpenWeather API key                                             |


### List of available endpoints:
| Endpoint                            | Description                                                                                            |
|-------------------------------------|--------------------------------------------------------------------------------------------------------|
| GET /health                         | Returns API health check                                                                               |
| GET /weather/:id                    | Gets weather data for weather id                                                                       |
| GET /weather-last/:cityId           | Gets last weather data for city id                                                                     |
| GET /weather-history/:cityId/:limit | Gets weather history for city id, ordered by date.<br/>Limit specifies total of records to be returned |
| GET /cities                         | Gets list of cities, ordered by name                                                                   |
| GET /forecast/:cityId               | Gets forecast data for city id                                                                         |