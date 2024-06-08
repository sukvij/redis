package concurrenctutils

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"os/signal"
// 	"runtime/debug"
// 	"syscall"
// 	"time"
// )

// var (
// 	logTag = "concurrencyutils"
// )

// // Go is a grabpay version of go routine, which is to create new context and pass to go routine.
// // It will fix the problem of context cancel upon response.
// func Go(ctx context.Context, logTag string, myFunc func(ctx context.Context) error) <-chan error {
// 	goctx := context.WithValue(context.Background(), contextkey.RequestID, ctx.Value(contextkey.RequestID))
// 	return gconcurrent.Go(goctx, logTag, myFunc)
// }

// // WaitShutdownSignal waits for shutdown signal, and if so, then waits for all channels finish or timeout.
// func WaitShutdownSignal() {
// 	waitShutdownSignal(nil)
// }

// func waitShutdownSignal(inter chan os.Signal) {
// 	interrupt := inter
// 	if interrupt == nil {
// 		interrupt = make(chan os.Signal, 1)
// 	}

// 	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

// 	// Wait for syscall.SIGINT or syscall.SIGTERM
// 	<-interrupt
// }

// // WaitOrTimeout waits for all channels or timeout.
// func WaitOrTimeout(timeout time.Duration, chs ...<-chan struct{}) {
// 	select {
// 	case <-WaitChannels(chs...):
// 	case <-time.After(timeout):
// 	}
// }

// // WaitChannels returns a channel which will be closed when receiving from all input channels.
// func WaitChannels(chs ...<-chan struct{}) <-chan struct{} {
// 	ret := make(chan struct{})
// 	go func() {
// 		for _, ch := range chs {
// 			<-ch
// 		}
// 		close(ret)
// 	}()
// 	return ret
// }

// // WaitFunction makes a function to run in another go routine, and returns a channel that closes only after the function finish execution.
// func WaitFunction(myfunc func()) <-chan struct{} {
// 	waitFor := make(chan struct{})

// 	go func() {
// 		defer func() {
// 			if r := recover(); r != nil {
// 				err, ok := r.(error)
// 				if !ok {
// 					err = fmt.Errorf("gconcurrent: %v", r)
// 				}
// 				logging.Error(logTag, "goroutine panic err:%s trace:%s", err.Error(), string(debug.Stack()))
// 			}
// 		}()
// 		myfunc()
// 		close(waitFor)
// 	}()

// 	return waitFor
// }

// // NewExecutionGroup returns a new instance of execution group.
// func NewExecutionGroup(ctx context.Context, timeout time.Duration) ExecutionGroup {
// 	childCtx, cancel := context.WithTimeout(ctx, timeout)
// 	childCtx = context.WithValue(childCtx, contextkey.RequestID, ctx.Value(contextkey.RequestID))

// 	return ExecutionGroup{
// 		eg:      gconcurrent.NewExecutionGroup(),
// 		ctx:     childCtx,
// 		timeout: timeout,
// 		cancel:  cancel,
// 	}
// }

// // ExecutionGroup is a timed version of typical use cases for WaitGroup.
// // It ensures all the sub-goroutines will get a context deadline exceed notification when timeout,
// // which can stop the expired jobs immediately
// type ExecutionGroup struct {
// 	eg      *gconcurrent.ExecutionGroup
// 	ctx     context.Context
// 	timeout time.Duration
// 	cancel  context.CancelFunc
// }

// // Go will execute the function in another go routine.
// // tag is used for tracking purpose in case of timeout
// func (eg *ExecutionGroup) Go(tag string, runnable func(ctx context.Context)) {
// 	eg.eg.Go(tag, func() { runnable(eg.ctx) })
// }

// // WaitForDone will block until all goroutines in the execution group finish, or at most for the given
// // timeout interval. It timeout occurs, it will return an error indicating which tag was failed.
// func (eg *ExecutionGroup) WaitForDone() error {
// 	defer eg.cancel()
// 	err := eg.eg.WaitForDone(eg.timeout + 100*time.Millisecond)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
