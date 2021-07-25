package backend

import (
	"calendar/internal/pkg/log"

	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func Run() {
	components, err := newApplications()
	if err != nil {
		log.System(err.Error())

		return
	}

	run(components)
}

func run(apps *applications) {
	signalChan := make(chan os.Signal, 2)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	ctx := context.Background()
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)

	go func() {
		<-signalChan
		cancel()
	}()

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		runHTTPServer(ctx, apps.restAPI, apps.config.restAPIAddress)
		wg.Done()
	}()

	wg.Wait()
}
