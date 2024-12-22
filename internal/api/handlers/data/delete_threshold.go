package data

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"
	service "goapi/internal/api/service/data"
)

// DeleteThresholdHandler deletes a threshold by ID.
func DeleteThresholdHandler(w http.ResponseWriter, r *http.Request, logger *log.Logger, ds service.DataService) {
	// Get the ID from the URL
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		// If the ID is not valid, return a 400 Bad Request
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid ID parameter."}`))
		return
	}

	// Set a context with timeout
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	// Call the service method to delete the threshold
	_, err = ds.DeleteThreshold(id, ctx)
	if err != nil {
		// If the deletion fails, log and return an internal server error
		logger.Println("Error deleting threshold:", err)
		http.Error(w, "Internal Server error.", http.StatusInternalServerError)
		return
	}

	// Return a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Threshold deleted successfully."}`))
}
