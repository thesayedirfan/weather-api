package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/thesayedirfan/weather/types"
)

type LocationService struct {
	url string
}

type ILocationService interface {
	GetLocation(ctx context.Context, ip string, second int) (*types.Location, error)
}

func NewLocationService(url string) ILocationService {
	return &LocationService{url: url}
}

func (l *LocationService) GetLocation(ctx context.Context, ip string, second int) (*types.Location, error) {

	ctx, cancel := context.WithTimeout(ctx, time.Duration(second)*time.Second)
	defer cancel()

	var result map[string]interface{}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, l.url+ip, nil)
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

	location := &types.Location{
		City:    strings.ToLower(result["city"].(string)),
		Country: strings.ToLower(result["country"].(string)),
	}

	return location, nil
}
