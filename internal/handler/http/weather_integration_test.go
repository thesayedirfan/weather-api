package http_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/thesayedirfan/weather/internal/cache"
	http_service "github.com/thesayedirfan/weather/internal/handler/http"
	"github.com/thesayedirfan/weather/types"
)

type MockLocationService struct{}

func (m *MockLocationService) GetLocation(ctx context.Context, ip string, second int) (*types.Location, error) {
	return &types.Location{City: "MockCity", Country: "MockCountry"}, nil
}

type MockWeatherService struct{}

func (m *MockWeatherService) GetWeatherInfo(ctx context.Context, city, format string, second int) (*types.Weather, error) {
	return &types.Weather{Temperature: "20", Humidity: "50", Description: "Mock Weather"}, nil
}

func TestGetWeatherByIP_Integration(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	locationService := &MockLocationService{}
	weatherService := &MockWeatherService{}
	weatherCacheStore := cache.NewWeatherCache()

	handler := http_service.NewHttpHandler(locationService, weatherService, weatherCacheStore)
	router.GET("/weather-by-ip/:ip", handler.GetWeatherByIP)

	expectedResponse := types.Response{
		IP: "127.0.0.1",
		Location: types.Location{
			City:    "MockCity",
			Country: "MockCountry",
		},
		Weather: types.Weather{
			Temperature: "20",
			Humidity:    "50",
			Description: "Mock Weather",
		},
	}

	weatherCacheStore.Set("127.0.0.1", expectedResponse, 10)

	req := httptest.NewRequest(http.MethodGet, "/weather-by-ip/127.0.0.1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var actualResponse types.Response
	err := json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponse)

	req = httptest.NewRequest(http.MethodGet, "/weather-by-ip/192.168.1.1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	err = json.Unmarshal(w.Body.Bytes(), &actualResponse)
	assert.NoError(t, err)
	assert.Equal(t, "MockCity", actualResponse.Location.City)
	assert.Equal(t, "MockCountry", actualResponse.Location.Country)
	assert.Equal(t, "20", actualResponse.Weather.Temperature)
	assert.Equal(t, "50", actualResponse.Weather.Humidity)
	assert.Equal(t, "Mock Weather", actualResponse.Weather.Description)
}
