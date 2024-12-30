package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thesayedirfan/weather/internal/cache"
	"github.com/thesayedirfan/weather/services"
	"github.com/thesayedirfan/weather/types"
	"github.com/thesayedirfan/weather/utils"
)

type HttpHandler struct {
	locationService   services.ILocationService
	weatherService    services.IWeatherService
	weatherCacheStore *cache.WeatherCacheStore
}

func NewHttpHandler(locationService services.ILocationService, weatherService services.IWeatherService, weatherCacheStore *cache.WeatherCacheStore) *HttpHandler {
	return &HttpHandler{
		locationService:   locationService,
		weatherService:    weatherService,
		weatherCacheStore: weatherCacheStore,
	}
}

func (h *HttpHandler) GetWeatherByIP(c *gin.Context) {
	defer utils.HandlePanic()

	ip := c.Param("ip")

	if ip == "" {
		ip = c.ClientIP()

	}

	if !utils.IsValidIP(ip) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "not valid IP address",
		})
		return
	}

	if cachedWeather, found := h.weatherCacheStore.Get(ip); found {
		c.JSON(http.StatusOK, cachedWeather)
		return
	}

	if ip == "" {
		ip = c.ClientIP()
	}

	location, err := h.locationService.GetLocation(c, ip, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch location",
		})
		return
	}

	weather, err := h.weatherService.GetWeatherInfo(c, location.City, "?format=j1", 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch weather",
		})
		return
	}

	response := types.Response{
		IP:       ip,
		Location: *location,
		Weather:  *weather,
	}

	h.weatherCacheStore.Set(ip, response, 10)
	c.JSON(http.StatusOK, response)
}
