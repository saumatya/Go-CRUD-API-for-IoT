package data

import (
	"context"
	"encoding/json"
	"goapi/internal/api/repository/models"
	service "goapi/internal/api/service/data"
	"log"
	"net/http"
	"time"
)

// PostThresholdHandler handles the creation of a new threshold.
func PostThresholdHandler(w http.ResponseWriter, r *http.Request, logger *log.Logger, ds service.DataService) {
	// Set a context with timeout for the request to ensure it doesn't hang indefinitely
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	// Parse the request body to get the new threshold data
	var threshold models.Threshold
	if err := json.NewDecoder(r.Body).Decode(&threshold); err != nil {
		// If there's an issue with the JSON decoding, return a 400 Bad Request
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid JSON body."}`))
		return
	}

	// Call the service method to create the new threshold
	err := ds.CreateThreshold(&threshold, ctx)
	if err != nil {
		// Handle specific error cases and return an appropriate HTTP status
		switch err.(type) {
		case service.DataError:
			// If a DataError is encountered, return a 400 Bad Request
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "` + err.Error() + `"}`))
			return
		default:
			// For any other unexpected errors, log and return a 500 Internal Server Error
			logger.Println("Error creating threshold:", err)
			http.Error(w, "Internal server error.", http.StatusInternalServerError)
			return
		}
	}

	// If creation is successful, send a success message
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Threshold created successfully."}`))
}
