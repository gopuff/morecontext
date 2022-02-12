# morecontext
--
    import "github.com/gopuff/morecontext"


## Usage

#### func  With

```go
func With(ctx context.Context, format string, args ...interface{}) context.Context
```
With is a helper for creating an extended Context instance.

#### type Context

```go
type Context struct {
	context.Context
	Message string
}
```

Context is a context.Context wrapper that is intended to include some extra
metadata about which context was cancelled. Useful for distinguishing e.g. http
request cancellation vs deadline vs process sigterm handling.

#### func (Context) Err

```go
func (c Context) Err() error
```
Err returns an error with extra context messaging

#### type CtxError

```go
type CtxError struct {
	Message  string
	Original error
}
```

CtxError implements error but separates the message from the original error in
case you want it.

#### func (*CtxError) Error

```go
func (c *CtxError) Error() string
```
Print out the message plus the metadata about which context was cancelled.

#### func (*CtxError) Unwrap

```go
func (c *CtxError) Unwrap() error
```
Unwrap supports errors.Is/errors.As
