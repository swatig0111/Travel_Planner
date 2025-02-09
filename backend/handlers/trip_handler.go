package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"fmt"
	"time"

	"travel-planner/backend/database"
	"travel-planner/backend/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddTrip - Handles adding a new trip
func AddTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var trip struct {
		Destination string `json:"destination"`
		StartDate   string `json:"start_date"`
		EndDate     string `json:"end_date"`
		
	}

	err := json.NewDecoder(r.Body).Decode(&trip)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if trip.Destination == "" || trip.StartDate == "" || trip.EndDate == "" {
		http.Error(w, "Missing required fields: destination, start_date, end_date", http.StatusBadRequest)
		return
	}

	// Convert date strings to time.Time
	startDate, err := time.Parse("2006-01-02", trip.StartDate)
	if err != nil {
		http.Error(w, "Invalid start_date format (Expected YYYY-MM-DD)", http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse("2006-01-02", trip.EndDate)
	if err != nil {
		http.Error(w, "Invalid end_date format (Expected YYYY-MM-DD)", http.StatusBadRequest)
		return
	}

	// Create a new trip object
	newTrip := models.Trip{
		ID:          primitive.NewObjectID(),
		Destination: trip.Destination,
		StartDate:   startDate,
		EndDate:     endDate,
		CreatedAt:   time.Now(),
	}

	// Database insertion with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = database.TravelCollection.InsertOne(ctx, newTrip)
	if err != nil {
		http.Error(w, "Could not insert trip into database", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Saving Trip: %+v\n", newTrip) 


	// Explicitly setting JSON field names in the response
	responseTrip := struct {
		ID          primitive.ObjectID `json:"id"`
		Destination string             `json:"destination"`
		StartDate   time.Time          `json:"start_date"`
		EndDate     time.Time          `json:"end_date"`
		CreatedAt   time.Time          `json:"created_at"`
	}{
		ID:          newTrip.ID,
		Destination: newTrip.Destination,
		StartDate:   newTrip.StartDate,
		EndDate:     newTrip.EndDate,
		CreatedAt:   newTrip.CreatedAt,
	}

	// Return created trip as response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseTrip)
}
