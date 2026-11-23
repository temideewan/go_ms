package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"ride-sharing/services/trip-service/internal/infrastructure/grpc"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"ride-sharing/shared/env"
	"syscall"

	grpcserver "google.golang.org/grpc"
)

var (
	GrpcAddr = env.GetString("GRPC_ADDR", ":9093")
)

func main() {
	inMemRepo := repository.NewInMemRepository()
	svc := service.NewService(inMemRepo)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigch := make(chan os.Signal, 1)
		signal.Notify(sigch, os.Interrupt, syscall.SIGTERM)
		<-sigch
		cancel()
	}()

	ls, err := net.Listen("tcp", GrpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// starting the grpc server
	grpcServer := grpcserver.NewServer()
	grpc.NewGRPCHandler(grpcServer, svc)
	log.Printf("starting GRPC server trip service on port %s", ls.Addr().String())

	go func() {
		if err := grpcServer.Serve(ls); err != nil {
			log.Printf("Failed to serve: %v", err)
			cancel()
		}
	}()

	// wait for the shutdown signal
	<-ctx.Done()
	log.Printf("Shutting down the server: %v", err)
	grpcServer.GracefulStop()
}
