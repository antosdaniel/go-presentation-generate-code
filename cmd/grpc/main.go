package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/antosdaniel/go-presentation-generate-code/internal/grpc"
	"github.com/antosdaniel/go-presentation-generate-code/internal/shutdown"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server := setup(ctx, cancel)

	log.Println("starting server...")
	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("unable to start gRPC server: %v", err) //nolint:gocritic
	}

	<-ctx.Done() // Wait for shutdown signal.
}

func setup(ctx context.Context, cancel context.CancelFunc) *http.Server {
	server := grpc.Setup()

	shutdown.SetupGraceful(ctx, cancel, shutdown.Closers{
		"grpc server": func() error { return server.Shutdown(ctx) },
	})

	return server
}
