package types

type Location struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type Weather struct {
	Temperature string `json:"temperature"`
	Humidity    string `json:"humidity"`
	Description string `json:"description"`
}

type Response struct {
	IP       string   `json:"ip"`
	Location Location `json:"location"`
	Weather  Weather  `json:"weather"`
}
