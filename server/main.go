package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	grpcServer := grpc.NewServer()
	go serveGrpc(grpcServer)
	log.Println("grpc server started on :9090")

	restServer := &http.Server{
		Addr:    ":8081",
		Handler: newHandler(),
	}
	go serveRest(restServer)
	log.Println("rest server started on :8081")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down grpc server...")
	grpcServer.GracefulStop()

	log.Println("Shutting down rest server...")
	if err := restServer.Shutdown(context.Background()); err != nil {
		log.Println("failed to gracefully shut down rest server", err)
		_ = restServer.Close()
	}
}
