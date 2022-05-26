package main

import (
	"cache-service/z_generated/pb"
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr    = flag.String("addr", "localhost:50051", "the address to connect rpc server")
	timeout = flag.Int("timeout", 1, "timeout duration in second")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// create a new rpc cache client
	c := pb.NewCacheClient(conn)
	// define context to set timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeout)*time.Second)
	defer cancel()

	// set cache value from rpc server
	_, err = c.Set(ctx, &pb.CacheSetInput{Key: "sample-key", Value: []byte("sample value")})
	if err != nil {
		log.Fatalf("unable to set cache: %v", err)
	}

	// get cache value from rpc server
	resp, err := c.Get(ctx, &pb.CacheGetInput{Key: "sample-key"})
	if err != nil {
		log.Fatalf("unable to get cache: %v", err)
	}
	fmt.Println(string(resp.GetValue()))
}
