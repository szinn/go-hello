package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/szinn/go-hello/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := flag.String("addr", "localhost:6952", "The address of the server to connect to")

	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGoHelloClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.IsHealthy(ctx, &pb.IsHealthyRequest{})
	if err != nil {
		log.Fatalf("Could not check if healthy: %v", err)
	}

	log.Printf("IsHealthy: %v", r.Healthy)
}
