package main

import (
	"fmt"
	"log"
	"net/http"
    "travel-planner/backend/config"
    "travel-planner/backend/database"
	"travel-planner/backend/routes"
	"travel-planner/backend/middleware" 
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}


func main() {
	config.LoadEnv()
	database.ConnectDB()
	r := routes.SetupRouter()

	corsHandler := middleware.CORS(r)

	

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
