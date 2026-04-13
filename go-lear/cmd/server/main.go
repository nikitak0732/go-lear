package main

import (
	"context"
	"log"
	"net"

	pb "test-module/grps/01/test-module/gen/calculator" // Импортируем из папки calculator

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalcServer
}

func (s *server) Add(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Printf("Adding %d + %d", req.A, req.B)
	return &pb.Response{Result: req.A + req.B}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterCalcServer(s, &server{})

	log.Println("Server started on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
