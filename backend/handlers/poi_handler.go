package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"travel-planner/backend/services"
)

// GetPlaceDetailsHandler handles requests for place details
func GetPlaceDetailsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get destination from query
	destination := r.URL.Query().Get("destination")
	if destination == "" {
		http.Error(w, "Missing destination parameter", http.StatusBadRequest)
		return
	}

	// Fetch destination places from Google Maps API
	placeDetails, err := services.FetchNearbyPlaces(destination)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching place data: %v", err), http.StatusInternalServerError)
		return
	}

	// Fetch nearby places using coordinates
	nearbyPlaces, err := services.FetchNearbyPlaces(destination)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching nearby places: %v", err), http.StatusInternalServerError)
		return
	}

	// Combine both destination and nearby places into a single response
	response := map[string]interface{}{
		"destination": placeDetails["results"],
		"nearby":      nearbyPlaces["results"],
	}

	// Encode and send the combined response
	json.NewEncoder(w).Encode(response)
}
