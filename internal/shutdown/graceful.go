package shutdown

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type ComponentName string

type Closers map[ComponentName]func() error

func SetupGraceful(ctx context.Context, cancel context.CancelFunc, closers Closers) {
	quitCh := make(chan os.Signal, 1)
	// interrupt signal sent from terminal
	signal.Notify(quitCh, syscall.SIGINT)
	// sigterm signal sent from kubernetes or docker
	signal.Notify(quitCh, syscall.SIGTERM)

	go func() {
		sig := <-quitCh
		log.Printf("received interrupt %q, shutting down...", sig)
		cancel()
		closeAll(closers)
	}()
}

func closeAll(closers Closers) {
	for name, c := range closers {
		if c == nil {
			continue
		}
		err := c()
		if err != nil {
			log.Printf("unable to close %q", name)
		}
	}
}
