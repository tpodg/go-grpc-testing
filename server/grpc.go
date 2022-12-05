package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"grpc-server/grpc/pb"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedDemoServiceServer
}

func (s *server) Send(_ context.Context, req *pb.Request) (*pb.Response, error) {
	log.Println("received request:", req)
	return &pb.Response{
		Value: fmt.Sprintf("returning '%s'", req.GetValue()),
	}, nil
}

func serveGrpc(s *grpc.Server) {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil && err != grpc.ErrServerStopped {
		log.Fatal("failed to listen", err)
	}

	pb.RegisterDemoServiceServer(s, &server{})
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	if err := s.Serve(listener); err != nil {
		log.Fatal("failed to serve", err)
	}
}
