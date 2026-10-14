package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"ride-sharing/shared/contracts"
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

	// todo call trip service
	port := 8083
	reqUrl := fmt.Sprintf("http://trip-service:%d/preview", port)

	jsonBody, _ := json.Marshal(reqBody)
	reader := bytes.NewReader(jsonBody)
	resp, err := http.Post(reqUrl, "application/json", reader)
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()

	var respBody any
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		http.Error(w, "Failed to parse json data", http.StatusBadRequest)
		return
	}

	response := contracts.APIResponse{Data: respBody}

	writeJSON(w, http.StatusCreated, response)
}
