package morecontext

import (
	"os"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSignals(t *testing.T) {
	asrt, rq := assert.New(t), require.New(t)

	ctx := ForSignals(syscall.SIGUSR1)

	proc, err := os.FindProcess(os.Getpid())
	asrt.NoError(err)
	rq.NotNil(proc)

	err = proc.Signal(syscall.SIGUSR1)
	asrt.NoError(err)

	<-ctx.Done()
	err = ctx.Err()
	rq.Error(err)
	asrt.Contains(err.Error(), "user defined")
}
