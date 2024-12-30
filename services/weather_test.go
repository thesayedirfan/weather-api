package services_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/thesayedirfan/weather/services"
)

func TestGetWeatherInfo_Success(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"current_condition": [
				{
					"temp_C": "25",
					"humidity": "60",
					"weatherDesc": [{"value": "Partly cloudy"}]
				}
			]
		}`))
	}))
	defer mockServer.Close()
	weatherService := services.NewWeatherService(mockServer.URL + "/")

	ctx := context.Background()
	weather, err := weatherService.GetWeatherInfo(ctx, "new york", "?format=json", 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if weather.Temperature != "25" || weather.Humidity != "60" || weather.Description != "Partly cloudy" {
		t.Errorf("unexpected result: %+v", weather)
	}
}

func TestGetWeatherInfo_Timeout(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second) // Simulate a delay
	}))
	defer mockServer.Close()

	weatherService := services.NewWeatherService(mockServer.URL + "/")

	ctx := context.Background()
	_, err := weatherService.GetWeatherInfo(ctx, "new york", "?format=json", 1)
	if err == nil {
		t.Fatalf("expected timeout error, got: %v", err)
	}
}
