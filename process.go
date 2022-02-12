package morecontext

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// sigCtx is a context that includes details
type sigCtx struct {
	context.Context

	exitSignal os.Signal
	m          sync.Mutex
}

// Err implements context.Context.Err but includes the os.Signal that caused
// the context cancellation.
func (sc *sigCtx) Err() error {
	sc.m.Lock()
	defer sc.m.Unlock()
	err := sc.Context.Err()

	return &MessageError{
		Message:  fmt.Sprintf("process exiting after getting signal %s", sc.exitSignal),
		Original: err,
	}
}

// Process returns a context.Context that will be Done if the given signals (or
// SIGTERM and SIGINT if none are passed) are received by the process.
func Process(sigs ...os.Signal) context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	// If no signals are returnd we will use a sensible default set.
	if len(sigs) == 0 {
		sigs = []os.Signal{syscall.SIGTERM, syscall.SIGINT}
	}

	sc := &sigCtx{Context: ctx}

	go func() {
		ch := make(chan os.Signal, 2)
		signal.Notify(ch, sigs...)

		i := 0
		for sig := range ch {
			i++
			if i > 1 {
				os.Exit(1)
			}
			sc.m.Lock()
			sc.exitSignal = sig
			sc.m.Unlock()

			cancel()
		}
	}()

	return sc
}
