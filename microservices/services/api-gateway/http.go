package main

import (
	"encoding/json"
	"net/http"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	// call the trip service
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Failed to parse json data", http.StatusBadRequest)
		return
	}

	// validation
	defer r.Body.Close()

	if reqBody.UserID == "" {
		http.Error(w, "userID is required", http.StatusBadRequest)
		return
	}
	if reqBody.Pickup.Latitude == 0 || reqBody.Pickup.Longitude == 0 {
		http.Error(w, "pickup longitude and latitude are required", http.StatusBadRequest)
		return
	}
	if reqBody.Destination.Latitude == 0 || reqBody.Destination.Longitude == 0 {
		http.Error(w, "destination longitude and latitude are required", http.StatusBadRequest)
		return
	}

	writeJSON(w, http.StatusCreated, "ok")
}
