package main

import (
	"cache-service/z_generated/pb"
	"context"
	"errors"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedCacheServer
}

func (s *server) Set(ctx context.Context, in *pb.CacheSetInput) (*pb.CacheSetResp, error) {
	return &pb.CacheSetResp{}, errors.New("not implemented yet. Min will implement me")
}

func (s *server) Get(ctx context.Context, in *pb.CacheGetInput) (*pb.CacheGetResp, error) {
	return &pb.CacheGetResp{
		Value: []byte("Min will implement me"),
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCacheServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
