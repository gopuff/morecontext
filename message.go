package morecontext

import (
	"context"
	"fmt"
)

// MessageContext is a context.Context wrapper that is intended to include some extra
// metadata about which context was cancelled. Useful for distinguishing e.g.
// http request cancellation vs deadline vs process sigterm handling.
type MessageContext struct {
	context.Context
	Message string
}

// Err returns an error with extra context messaging
func (c MessageContext) Err() error {
	return &MessageError{
		Original: c.Context.Err(),
		Message:  c.Message,
	}
}

var _ context.Context = MessageContext{}

// WithMessage is a helper for creating a MessageContext instance, with more
// context about exactly which context cancellation occurred.
func WithMessage(ctx context.Context, format string, args ...interface{}) context.Context {
	return MessageContext{
		Context: ctx,
		Message: fmt.Sprintf(format, args...),
	}
}

// MessageError implements error but separates the message from the original error
// in case you want it.
type MessageError struct {
	Message  string
	Original error
}

// Error implements error and prints out the message plus the metadata about
// which context was cancelled.
func (c *MessageError) Error() string {
	return fmt.Sprintf("%s: %s", c.Message, c.Original)
}

// Unwrap supports errors.Is/errors.As
func (c *MessageError) Unwrap() error {
	return c.Original
}
