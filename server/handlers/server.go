package main

import (
	"cache-service/server/data"
	"cache-service/z_generated/pb"
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCacheServer
	cache cacheDB
}

type cacheDB interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) ([]byte, error)
	SetUser(ctx context.Context, key string, value []byte) error
	GetUser(ctx context.Context, key string) ([]byte, error)
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()

	s := grpc.NewServer()
	cacheDB := data.NewRedisClient("minhtet")

	pb.RegisterCacheServer(s, &server{
		cache: cacheDB,
	})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
