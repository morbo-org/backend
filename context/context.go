package context

import (
	"context"
	"sync"
	"time"
)

type (
	Context    = context.Context
	CancelFunc = context.CancelFunc
)

var (
	ErrCanceled       = context.Canceled
	ErrDeadlineExceed = context.DeadlineExceeded
)

func Background() Context {
	return context.Background()
}

func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	return context.WithTimeout(parent, timeout)
}

type (
	waitGroupKeyType   int
	waitGroupValueType = *sync.WaitGroup
)

const waitGroupKey waitGroupKeyType = 0

func GetWaitGroup(ctx Context) waitGroupValueType {
	return ctx.Value(waitGroupKey).(waitGroupValueType)
}

func WithWaitGroup(parent Context) (Context, CancelFunc) {
	ctx := context.WithValue(parent, waitGroupKey, new(sync.WaitGroup))
	ctx, innerCancel := context.WithCancel(ctx)
	cancel := func() {
		innerCancel()
		GetWaitGroup(ctx).Wait()
	}
	return ctx, cancel
}
