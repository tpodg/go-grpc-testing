package main

import (
	"context"
	"fmt"
	"github.com/tpodg/go-grpc-testing/client/grpc/pb"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type grpcClient struct {
	pb.DemoServiceClient
}

type restClient struct {
	*http.Client
	url string
}

func main() {
	config := cfg{
		Grpc: grpcCfg{
			Target: "localhost:9090",
			Tls:    false,
		},
		Rest: restCfg{
			Target: "http://localhost:8081",
		},
	}
	config.Parse()
	fmt.Printf("config: %+v", config)

	s := http.Server{
		Addr:    ":8080",
		Handler: newHandler(config),
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("failed to listen and serve", err)
		}
	}()

	log.Println("server started on :8080")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("shutting down server...")
	if err := s.Shutdown(context.Background()); err != nil {
		log.Println("failed to shut down gracefully", err)
		_ = s.Close()
	}
}
