package data

import (
	"context"
	"encoding/json"
	service "goapi/internal/api/service/data"
	"log"
	"net/http"
	"strconv"
	"time"
)

// GetThresholdHandler retrieves a list of thresholds, supporting pagination.
func GetThresholdHandler(w http.ResponseWriter, r *http.Request, logger *log.Logger, ds service.DataService) {
	// Get the page number from the query parameters, default to 0 if not provided
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		if err.(*strconv.NumError) != nil {
			// If no page is specified, set to 0 (default)
			page = 0
		} else {
			// If the page parameter is invalid, return a 400 Bad Request
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Invalid page specified."}`))
			return
		}
	}

	// Get the number of items per page (rowsPerPage), defaulting to 10
	rowsPerPage := 10
	if query := r.URL.Query().Get("rowsPerPage"); query != "" {
		rowsPerPage, err = strconv.Atoi(query)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Invalid rowsPerPage specified."}`))
			return
		}
	}

	// Set a context with timeout to avoid blocking indefinitely
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	// Fetch the list of thresholds with pagination using GetAllThresholds method
	thresholds, err := ds.GetAllThresholds(page, rowsPerPage, ctx)
	if err != nil {
		// Log and return internal server error if data retrieval fails
		logger.Println("Error retrieving thresholds:", err)
		http.Error(w, "Internal Server error.", http.StatusInternalServerError)
		return
	}

	// If no thresholds are found, return a 404 Not Found
	if len(thresholds) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "No thresholds found."}`))
		return
	}

	// Return the list of thresholds as a JSON response with a 200 OK status
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(thresholds); err != nil {
		// Log and return an internal error if encoding fails
		logger.Println("Error encoding thresholds:", err)
		http.Error(w, "Internal Server error.", http.StatusInternalServerError)
		return
	}
}
