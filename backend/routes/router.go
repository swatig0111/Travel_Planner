package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"travel-planner/backend/database"
	"travel-planner/backend/handlers" 

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TravelPlan model
type TravelPlan struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Dest  string             `bson:"destination" json:"destination"`
	Start time.Time          `bson:"start_date" json:"start_date"`
	End   time.Time          `bson:"end_date" json:"end_date"`
	CreatedAt time.Time      `bson:"created_at" json:"created_at"` 
}

// SetupRouter initializes the API routes
func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/travel", AddTravelPlan).Methods("POST")
	r.HandleFunc("/travel/{id}", GetTravelPlan).Methods("GET")
	r.HandleFunc("/places", handlers.GetPlaceDetailsHandler).Methods("GET")

	return r
}

// AddTravelPlan - Add a new travel plan
func AddTravelPlan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var plan struct {
		Dest  string `json:"destination"`
		Start string `json:"start_date"`
		End   string `json:"end_date"`
	}

	err := json.NewDecoder(r.Body).Decode(&plan)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Parse time from string to time.Time
	startDate, err := time.Parse("2006-01-02", plan.Start)
	if err != nil {
		http.Error(w, "Invalid start_date format (Expected YYYY-MM-DD)", http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse("2006-01-02", plan.End)
	if err != nil {
		http.Error(w, "Invalid end_date format (Expected YYYY-MM-DD)", http.StatusBadRequest)
		return
	}

	newPlan := TravelPlan{
		ID:    primitive.NewObjectID(),
		Dest:  plan.Dest,
		Start: startDate,
		End:   endDate,
	}

	_, err = database.TravelCollection.InsertOne(context.TODO(), newPlan)
	if err != nil {
		http.Error(w, "Failed to insert travel plan", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newPlan)
}

// GetTravelPlan - Fetch a travel plan by ID
func GetTravelPlan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid trip ID format", http.StatusBadRequest)
		return
	}

	var plan TravelPlan
	err = database.TravelCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&plan)
	if err != nil {
		http.Error(w, "Trip not found", http.StatusNotFound)
		return
	}

	fmt.Printf("Fetched Trip: %+v\n", plan)

	json.NewEncoder(w).Encode(plan)
}
