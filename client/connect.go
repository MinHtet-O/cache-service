package main

import (
	"cache-service/z_generated/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func connectRPC(c *pb.CacheClient) *grpc.ClientConn {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// create a new rpc cache client
	*c = pb.NewCacheClient(conn)
	return conn
}
