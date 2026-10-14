package main

import (
	"log"
	"net/http"
	h "ride-sharing/services/trip-service/internal/infrastructure/http"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"ride-sharing/shared/env"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8083")
)

func main() {
	log.Println("Starting Trip service here")
	log.Printf("http address %s", httpAddr)

	mux := http.NewServeMux()
	inMemRepo := repository.NewInMemRepository()
	svc := service.NewService(inMemRepo)

	httpHandler := h.HttpHandler{
		Service: svc,
	}
	mux.HandleFunc("POST /preview", httpHandler.HandleTripPreview)
	server := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("HTTP server error: %v", err)
	}
}
