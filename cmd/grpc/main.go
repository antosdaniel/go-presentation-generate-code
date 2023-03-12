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
	server := grpc.Setup()
	defer server.Shutdown(ctx)

	log.Println("starting server...")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("unable to start gRPC server: %v", err) //nolint:gocritic
	}

	<-ctx.Done() // Wait for shutdown signal.
}
