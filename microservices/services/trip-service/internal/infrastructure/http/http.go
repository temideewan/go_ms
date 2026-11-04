package http

import (
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/types"
	"time"
)

type HttpHandler struct {
	Service domain.TripService
}

type previewTripRequest struct {
	UserID      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

func (h *HttpHandler) HandleTripPreview(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 9)
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Failed to parse json data", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	log.Printf("pickup: %+v. destination: %+v", reqBody.Pickup, reqBody.Destination)
	t, err := h.Service.GetRoute(ctx, &reqBody.Pickup, &reqBody.Destination, true)
	if err != nil {
		log.Println(err)
	}

	writeJSON(w, http.StatusCreated, t)
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
