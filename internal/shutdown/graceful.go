package shutdown

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func SetupGraceful(ctx context.Context, cancel context.CancelFunc) {
	quitCh := make(chan os.Signal, 1)
	// interrupt signal sent from terminal
	signal.Notify(quitCh, syscall.SIGINT)
	// sigterm signal sent from kubernetes or docker
	signal.Notify(quitCh, syscall.SIGTERM)

	go func() {
		sig := <-quitCh
		log.Printf("received interrupt %q, shutting down...", sig)
		cancel()
	}()
}
