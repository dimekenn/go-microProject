package main

import (
	"context"
	"fmt"
	"github.com/unistack-org/micro/service"
	"github.com/unistack-org/micro/v3/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errCh := make(chan error, 1)

	logger.DefaultLogger = logger.NewLogger(logger.WithLevel(logger.TraceLevel))

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
		errCh <- fmt.Errorf("%s", <-sigChan)
	}()

	go service.StartHTTPService(ctx, errCh)

	fmt.Printf("Terminated: %v\n", <-errCh)
}
