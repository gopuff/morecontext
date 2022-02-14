# morecontext
--
    import "github.com/gopuff/morecontext"


## Usage

#### func  ForSignals

```go
func ForSignals(sigs ...os.Signal) context.Context
```
ForSignals returns a context.Context that will be cancelled if the given signals
(or SIGTERM and SIGINT by default, if none are passed) are received by the
process.

#### func  WithMessage

```go
func WithMessage(ctx context.Context, format string, args ...interface{}) context.Context
```
WithMessage is a helper for creating a MessageContext instance, with more
context about exactly which context cancellation occurred.

#### type MessageContext

```go
type MessageContext struct {
	context.Context
	Message string
}
```

MessageContext is a context.Context wrapper that is intended to include some
extra metadata about which context was cancelled. Useful for distinguishing e.g.
http request cancellation vs deadline vs process sigterm handling.

#### func (MessageContext) Err

```go
func (c MessageContext) Err() error
```
Err returns an error with extra context messaging

#### type MessageError

```go
type MessageError struct {
	Message  string
	Original error
}
```

MessageError implements error but separates the message from the original error
in case you want it.

#### func (*MessageError) Error

```go
func (c *MessageError) Error() string
```
Error implements error and prints out the message plus the metadata about which
context was cancelled.

#### func (*MessageError) Unwrap

```go
func (c *MessageError) Unwrap() error
```
Unwrap supports errors.Is/errors.As
