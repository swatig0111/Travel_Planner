package handlers

import (
	"encoding/json"
	"net/http"
	 "travel-planner/backend/services"
)

func EstimateCost(w http.ResponseWriter, r *http.Request) {
	var request map[string]interface{}
	json.NewDecoder(r.Body).Decode(&request)
	cost := services.CalculateCost(request)
	json.NewEncoder(w).Encode(map[string]float64{"estimated_cost": cost})
}
