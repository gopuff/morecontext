package morecontext

import (
	"context"
	"fmt"
)

// A CancelReasonContext is a cancellable context whose cancellation requires a
// reason, and whose Err() message will include the reason, if this context was
// the one cancelled.
type CancelReasonContext struct {
	context.Context
	Reason error
}

var _ context.Context = &CancelReasonContext{}

// If this context was the one cancelled, we return the reason it was
// cancelled. If the underlying context was cancelled, then we return its
// message.
func (crc CancelReasonContext) Err() error {
	if crc.Reason != nil {
		return fmt.Errorf("context cancelled: %w", crc.Reason)
	}
	return crc.Context.Err()
}

// WithCancelReason returns a context implementation that must be cancelled,
// and whose cancellation must include a Reason error that will be returned by
// any calls to `Err`.
func WithCancelReason(ctx context.Context) (*CancelReasonContext, func(error)) {
	ctx, cancel := context.WithCancel(ctx)
	crc := CancelReasonContext{
		Context: ctx,
	}

	c := func(err error) {
		crc.Reason = err
		cancel()
	}

	return &crc, c
}
