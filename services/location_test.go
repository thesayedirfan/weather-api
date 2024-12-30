package services_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/thesayedirfan/weather/services"
)

func TestGetLocation_Success(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"city": "New York", "country": "USA"}`))
	}))
	defer mockServer.Close()

	locationService := services.NewLocationService(mockServer.URL + "/")

	ctx := context.Background()
	location, err := locationService.GetLocation(ctx, "127.0.0.1", 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if location.City != "new york" || location.Country != "usa" {
		t.Errorf("unexpected result: %+v", location)
	}
}

func TestGetLocation_Timeout(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second) // Simulate a delay
	}))
	defer mockServer.Close()

	locationService := services.NewLocationService(mockServer.URL + "/")

	ctx := context.Background()
	_, err := locationService.GetLocation(ctx, "127.0.0.1", 1)
	if err == nil || err.Error() != "request timed out" {
		t.Fatalf("expected timeout error, got: %v", err)
	}
}
