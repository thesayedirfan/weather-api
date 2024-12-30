# Weather API

A backend service built with Go to provide weather information based on an IP address. The API integrates with a geolocation service and a weather service to aggregate location and weather details.

## Setup Instructions

### Prerequisites

- Docker

### Steps

1. **Clone the Repository**
   ```bash
   git clone https://github.com/yourname/weather-api.git
   cd weather-api
2. **Run the Docker Container**
    ```bash
    docker-compose up --build
3. **Run Tests**

    > Make sure the docker container is running

    ```bash
    docker compose exec weather-api  go test ./...

## API Documentation

| **Endpoint**         | **Method** | **Description**                                                             | **Parameters**                                                                                                                                                                 | **Response**                                                                                                                                                  |
|-----------------------|------------|-----------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `/health`            | GET        | Health check endpoint to verify the server is running.                     | None                                                                                                                                                                        | **Status Code**: 200 - OK                                                                                                                                    |
| `/weather-by-ip/:ip` | GET        | Retrieve weather information based on an IP address.                       | **Path Parameter**: `ip` - The IP address to retrieve weather data for. If not provided, defaults to the client's IP address.                           | **Success**: Returns weather details and location information. <br> **Error**: 400 for invalid IP.  429 for too many request withing a specified timeframe |

## Project Structure and Design Explanation

```plaintext

├── cmd
│   └── main.go                  # Entry point for the application
├── internal
│   ├── cache
│   │   ├── request.go           # a thread safe inmemory cache to we store the count of the request to rate limit the ip address
│   │   ├── request_test.go
│   │   ├── weather.go            # a thread safe inmemory cache to store the result of weather with ttl
│   │   └── weather_test.go
│   ├── handler
│   │   └── http
│   │       ├── heath.go             # an endpoint to check heath
│   │       ├── weather.go           # an endpoint to get the weather of the ip address with caching and timeout on external services
│   │       └── weather_integration_test.go 
│   └── middleware
│       └── rate_limiter.go      # an middlware to rate limit the ip address  
├── postman_collection
│   └── weather.json
├── services
│   ├── location.go                # an location service call with timeout incase the service does not response
│   ├── location_test.go
│   ├── weather.go                 # an weather service call with timeout incase the service does not response
│   └── weather_test.go
├── types
│   └── main.go                    
└── utils
│    ├── main.go
│    ├── validation.go             
│    └── validation_test.go
├── docker-compose.yml
├── go.mod
├── go.sum
├── Dockerfile    # docker compose file with heath checks
├── Readme.md    # Documentation for the project
```