package data

import (
	"context"
	"encoding/json"
	"goapi/internal/api/repository/models"
	service "goapi/internal/api/service/data"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func UpdateThresholdHandler(w http.ResponseWriter, r *http.Request, logger *log.Logger, ds service.DataService) {
	// Get the ID from the URL path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 { // Expecting "/threshold/{id}" in URL
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid URL structure."}`))
		return
	}

	id, err := strconv.Atoi(parts[2]) // The ID is the 3rd part of the path
	if err != nil {
		// If the ID is invalid, return a 400 Bad Request
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid ID parameter."}`))
		return
	}

	// Set a context with timeout for the request
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	// Parse the request body to get the updated threshold data
	var threshold models.Threshold
	if err := json.NewDecoder(r.Body).Decode(&threshold); err != nil {
		// If there's an issue with the JSON decoding, return a 400 Bad Request
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid JSON body."}`))
		return
	}

	// Set the ID for the threshold, using the ID passed in the path
	threshold.ID = id

	// Call the service method to update the threshold
	_, err = ds.UpdateThreshold(&threshold, ctx)
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
			logger.Println("Error updating threshold:", err)
			http.Error(w, "Internal server error.", http.StatusInternalServerError)
			return
		}
	}

	// If the update is successful, send a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Threshold updated successfully."}`))
}
