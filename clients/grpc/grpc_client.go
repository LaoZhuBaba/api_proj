package grpc_client

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func Client() {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "localhost:3000")
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	defer conn.Close()
}
