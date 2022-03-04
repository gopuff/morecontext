package morecontext

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReasonContext(t *testing.T) {
	asrt := assert.New(t)
	ctx, cancel := WithCancelReason(context.Background())

	cancel(fmt.Errorf("foo bar baz"))

	err := ctx.Err()
	asrt.Error(err)
	asrt.Contains(err.Error(), "foo bar baz")
}

func TestReasonParentCancel(t *testing.T) {
	asrt := assert.New(t)
	ctx, c1 := context.WithCancel(context.Background())
	ctx, cancel := WithCancelReason(ctx)
	defer cancel(nil)

	c1()

	err := ctx.Err()
	asrt.Error(err)
	asrt.NotContains(err.Error(), "foo bar baz")
}
