package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/thesayedirfan/weather/internal/cache"
	"github.com/thesayedirfan/weather/internal/handler/http"
	"github.com/thesayedirfan/weather/internal/middleware"
	"github.com/thesayedirfan/weather/services"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	locationServiceURL := os.Getenv("LOCATION_SERVICE_URL")
	weatherServiceURL := os.Getenv("WEATHER_SERVICE_URL")

	locationService := services.NewLocationService(locationServiceURL)
	weatherService := services.NewWeatherService(weatherServiceURL)
	weatherCacheStore := cache.NewWeatherCache()
	handler := http.NewHttpHandler(locationService, weatherService, weatherCacheStore)
	router := gin.Default()
	cache := cache.NewRequestCache()
	router.Use(middleware.RateLimiter(cache, 5))
	router.GET("/health", handler.Health)
	router.GET("/weather-by-ip/:ip", handler.GetWeatherByIP)
	router.GET("/weather-by-ip", handler.GetWeatherByIP)
	router.Run()
}
