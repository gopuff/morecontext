package morecontext

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// ForSignals returns a context.Context that will be cancelled if the given
// signals (or SIGTERM and SIGINT by default, if none are passed) are received
// by the process.
func ForSignals(sigs ...os.Signal) context.Context {
	ctx, cancel := WithCancelReason(context.Background())

	// If no signals are included we will use a sensible default set.
	if len(sigs) == 0 {
		sigs = []os.Signal{syscall.SIGTERM, syscall.SIGINT}
	}

	ch := make(chan os.Signal, 2)
	signal.Notify(ch, sigs...)

	go func() {
		i := 0
		for sig := range ch {
			i++
			if i > 1 {
				os.Exit(1)
			}
			cancel(fmt.Errorf("got signal %s", sig))
		}
	}()

	return ctx
}
