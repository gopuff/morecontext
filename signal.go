package morecontext

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// sigCtx is a context that will be cancelled if certain signals are received.
// Its Err will include details if this is the reason it was cancelled.
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

	if sc.exitSignal == nil {
		return nil
	}

	return &MessageError{
		Message:  fmt.Sprintf("context cancelled: got signal %s", sc.exitSignal.String()),
		Original: err,
	}
}

// ForSignals returns a context.Context that will be cancelled if the given
// signals (or SIGTERM and SIGINT by default, if none are passed) are received
// by the process.
func ForSignals(sigs ...os.Signal) context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	// If no signals are returnd we will use a sensible default set.
	if len(sigs) == 0 {
		sigs = []os.Signal{syscall.SIGTERM, syscall.SIGINT}
	}

	sc := &sigCtx{Context: ctx}

	ch := make(chan os.Signal, 2)
	signal.Notify(ch, sigs...)

	go func() {

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
