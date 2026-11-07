package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	h "ride-sharing/services/trip-service/internal/infrastructure/http"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"ride-sharing/shared/env"
	"syscall"
	"time"
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

	serverError := make(chan error, 1)

	go func() {
		log.Printf("Trip service listening on port: %v", httpAddr)
		serverError <- server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverError:
		log.Printf("HTTP server error: %v", err)
	case sig := <-shutdown:
		log.Printf("Server is shutting down due to %v signal", sig)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Could not server gracefully: %v", err)
			server.Close()
		}
	}
}
