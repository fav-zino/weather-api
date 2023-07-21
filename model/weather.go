package model


type Weather struct {
    City       string    `json:"city"`
    CountryCode       string    `json:"country_code"`
    Temperature float64   `json:"temperature"`
    Description string    `json:"description"`
    Humidity    float64       `json:"humidity"`
    WindSpeed   float64   `json:"wind_speed"`
}
