package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"travel-planner/backend/config"
)

// Google Maps API Base URLs
const (
	googleGeocodeURL = "https://maps.googleapis.com/maps/api/geocode/json"
	googleNearbySearchURL = "https://maps.googleapis.com/maps/api/place/nearbysearch/json"
)

// GetCoordinates fetches latitude and longitude for a destination
func GetCoordinates(destination string) (float64, float64, error) {
	googleAPIKey := config.GetEnv("GOOGLE_MAPS_API_KEY")
	geocodeURL := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", url.QueryEscape(destination), googleAPIKey)

	resp, err := http.Get(geocodeURL)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	var geoData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&geoData); err != nil {
		return 0, 0, err
	}

	// Check if results exist and assert types safely
	results, ok := geoData["results"].([]interface{})
	if !ok || len(results) == 0 {
		return 0, 0, fmt.Errorf("no coordinates found for destination")
	}

	firstResult, ok := results[0].(map[string]interface{})
	if !ok {
		return 0, 0, fmt.Errorf("unexpected data format in results")
	}

	geometry, ok := firstResult["geometry"].(map[string]interface{})
	if !ok {
		return 0, 0, fmt.Errorf("unexpected data format in geometry")
	}

	location, ok := geometry["location"].(map[string]interface{})
	if !ok {
		return 0, 0, fmt.Errorf("unexpected data format in location")
	}

	lat, latOk := location["lat"].(float64)
	lng, lngOk := location["lng"].(float64)
	if !latOk || !lngOk {
		return 0, 0, fmt.Errorf("latitude or longitude not found")
	}

	return lat, lng, nil
}


// FetchNearbyPlaces fetches nearby places for a destination
func FetchNearbyPlaces(destination string) (map[string]interface{}, error) {
	lat, lng, err := GetCoordinates(destination)
	if err != nil {
		return nil, err
	}

	googleAPIKey := config.GetEnv("GOOGLE_MAPS_API_KEY")
	url := fmt.Sprintf("%s?location=%f,%f&radius=5000&key=%s", googleNearbySearchURL, lat, lng, googleAPIKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
