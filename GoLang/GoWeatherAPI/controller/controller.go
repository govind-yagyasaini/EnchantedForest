package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/govind-yagyasaini/EnchantedForest/GoLang/GoWeatherAPI/models"
)

func CurrentByName(location string, apiKey string) (*models.CurrentWeatherData, error) {
	weather := &models.CurrentWeatherData{}
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", location, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get weather data: %v", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return weather, nil
}
