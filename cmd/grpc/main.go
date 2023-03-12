package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	dbsetup "github.com/antosdaniel/go-presentation-generate-code/internal/db"
	"github.com/antosdaniel/go-presentation-generate-code/internal/grpc"
	"github.com/antosdaniel/go-presentation-generate-code/internal/shutdown"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	server, err := setup(ctx, cancel)
	if err != nil {
		log.Fatalf("setup failed: %v", err)
	}

	log.Println("starting server...")
	err = server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("unable to start gRPC server: %v", err)
	}

	<-ctx.Done() // Wait for shutdown signal.
}

func setup(ctx context.Context, cancel context.CancelFunc) (*http.Server, error) {
	db, err := dbsetup.New(ctx)
	if err != nil {
		return nil, err
	}
	server := grpc.New(db)

	shutdown.SetupGraceful(ctx, cancel, shutdown.Closers{
		"grpc server": func() error { return server.Shutdown(ctx) },
		"db":          func() error { return db.Close() },
	})

	return server, nil
}
