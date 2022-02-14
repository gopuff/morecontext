package morecontext

import (
	"context"
	"testing"
	"time"

	"github.com/alecthomas/assert"
)

func TestMore(t *testing.T) {
	asrt := assert.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	ctx = WithMessage(ctx, "my test context")
	<-ctx.Done()
	asrt.Contains(ctx.Err().Error(), "my test context")
}

func TestNoNest(t *testing.T) {
	asrt := assert.New(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	ctx = WithMessage(WithMessage(ctx, "parent will be cancelled"), "will not be cancelled directly")
	<-ctx.Done()
	asrt.Contains(ctx.Err().Error(), "parent will be cancelled")
	asrt.NotContains(ctx.Err().Error(), "will not be cancelled")
}
