package main

import (
	"context"
	"log"

	"github.com/antosdaniel/go-presentation-generate-code/internal/grpc"
	"github.com/antosdaniel/go-presentation-generate-code/internal/shutdown"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shutdown.SetupGraceful(ctx, cancel)
	log.Println("starting server...")
	err := grpc.StartServer()
	if err != nil {
		log.Fatalf("unable to start gRPC server: %v", err) //nolint:gocritic
	}

	<-ctx.Done() // wait for shutdown signal
}
