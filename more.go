package morecontext

import (
	"context"
	"fmt"
)

// Context is a context.Context wrapper that is intended to include some extra
// metadata about which context was cancelled. Useful for distinguishing e.g.
// http request cancellation vs deadline vs process sigterm handling.
type Context struct {
	context.Context
	Message string
}

// Err returns an error with extra context messaging
func (c Context) Err() error {
	return &CtxError{
		Original: c.Context.Err(),
		Message:  c.Message,
	}
}

var _ context.Context = Context{}

// With is a helper for creating an extended Context instance.
func With(ctx context.Context, format string, args ...interface{}) context.Context {
	return Context{
		Context: ctx,
		Message: fmt.Sprintf(format, args...),
	}
}

// CtxError implements error but separates the message from the original error
// in case you want it.
type CtxError struct {
	Message  string
	Original error
}

// Print out the message plus the metadata about which context was cancelled.
func (c *CtxError) Error() string {
	return fmt.Sprintf("%s context error: %s", c.Message, c.Original)
}

// Unwrap supports errors.Is/errors.As
func (c *CtxError) Unwrap() error {
	return c.Original
}
