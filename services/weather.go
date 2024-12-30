package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/thesayedirfan/weather/types"
)

type IWeatherService interface {
	GetWeatherInfo(ctx context.Context, city, format string, second int) (*types.Weather, error)
}

type WeatherService struct {
	url string
}

func NewWeatherService(url string) IWeatherService {
	return &WeatherService{url: url}
}

func (w *WeatherService) GetWeatherInfo(ctx context.Context, city, format string, second int) (*types.Weather, error) {

	var result map[string]interface{}

	_, cancel := context.WithTimeout(ctx, time.Duration(second)*time.Microsecond)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, w.url+city+format, nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, fmt.Errorf("request timed out")
		}
		return nil, fmt.Errorf("failed to make request: %w", err)
	}

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	weatherCondition := result["current_condition"].([]interface{})[0].(map[string]interface{})

	weather := &types.Weather{
		Temperature: weatherCondition["temp_C"].(string),
		Humidity:    weatherCondition["humidity"].(string),
		Description: weatherCondition["weatherDesc"].([]interface{})[0].(map[string]interface{})["value"].(string),
	}

	return weather, err
}
